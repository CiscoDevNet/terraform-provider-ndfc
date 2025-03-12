// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Package nd is a Cisco Nexus Dashboard REST client library for Go.
package nd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"sync"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const DefaultMaxRetries int = 3
const DefaultBackoffMinDelay int = 2
const DefaultBackoffMaxDelay int = 60
const DefaultBackoffDelayFactor float64 = 3

// Client is an HTTP Nexus Dashboard client.
// Use nd.NewClient to initiate a client.
// This will ensure proper cookie handling and processing of modifiers.
type Client struct {
	// HttpClient is the *http.Client used for API requests.
	HttpClient *http.Client
	// Url is the Nexus Dashboard IP or hostname, e.g. https://10.0.0.1:443 (port is optional).
	Url string
	// BasePath is the Nexus Dashboard URL prefix to use, e.g. '/appcenter/cisco/ndfc/api/v1'.
	BasePath string
	// Token is the current authentication token
	Token string
	// Usr is the Nexus Dashboard username.
	Usr string
	// Pwd is the Nexus Dashboard password.
	Pwd string
	// Domain is the Nexus Dashboard domain.
	Domain string
	// Insecure determines if insecure https connections are allowed.
	Insecure bool
	// Maximum number of retries
	MaxRetries int
	// Minimum delay between two retries
	BackoffMinDelay int
	// Maximum delay between two retries
	BackoffMaxDelay int
	// Backoff delay factor
	BackoffDelayFactor float64
	// Authentication mutex
	AuthenticationMutex *sync.Mutex
	AuthTimeStamp       time.Time
	AuthTokenTimeout    time.Duration
}

// NewClient creates a new Nexus Dashboard HTTP client.
// Pass modifiers in to modify the behavior of the client, e.g.
//
//	client, _ := NewClient("https://10.1.1.1", "/appcenter/cisco/ndfc/api/v1", "user", "password", "", true, RequestTimeout(120))
func NewClient(url, basePath, usr, pwd, domain string, insecure bool, mods ...func(*Client)) (Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	cookieJar, _ := cookiejar.New(nil)
	httpClient := http.Client{
		Timeout:   60 * time.Second,
		Transport: tr,
		Jar:       cookieJar,
	}

	if domain == "" {
		domain = "DefaultAuth"
	}

	client := Client{
		HttpClient:          &httpClient,
		Url:                 url,
		BasePath:            basePath,
		Usr:                 usr,
		Pwd:                 pwd,
		Domain:              domain,
		Insecure:            insecure,
		MaxRetries:          DefaultMaxRetries,
		BackoffMinDelay:     DefaultBackoffMinDelay,
		BackoffMaxDelay:     DefaultBackoffMaxDelay,
		BackoffDelayFactor:  DefaultBackoffDelayFactor,
		AuthenticationMutex: &sync.Mutex{},
		AuthTokenTimeout:    2 * time.Minute,
	}

	for _, mod := range mods {
		mod(&client)
	}
	if client.checkAndFillTokenTimeout() {
		log.Printf("[INFO] Token timeout set to %v", client.AuthTokenTimeout)
	} else {
		log.Printf("[ERROR] Token timeout could not be read, using default - %v", client.AuthTokenTimeout)
	}
	return client, nil
}

// RequestTimeout modifies the HTTP request timeout from the default of 60 seconds.
func RequestTimeout(x time.Duration) func(*Client) {
	return func(client *Client) {
		client.HttpClient.Timeout = x * time.Second
	}
}

// MaxRetries modifies the maximum number of retries from the default of 3.
func MaxRetries(x int) func(*Client) {
	return func(client *Client) {
		client.MaxRetries = x
	}
}

// BackoffMinDelay modifies the minimum delay between two retries from the default of 2.
func BackoffMinDelay(x int) func(*Client) {
	return func(client *Client) {
		client.BackoffMinDelay = x
	}
}

// BackoffMaxDelay modifies the maximum delay between two retries from the default of 60.
func BackoffMaxDelay(x int) func(*Client) {
	return func(client *Client) {
		client.BackoffMaxDelay = x
	}
}

// BackoffDelayFactor modifies the backoff delay factor from the default of 3.
func BackoffDelayFactor(x float64) func(*Client) {
	return func(client *Client) {
		client.BackoffDelayFactor = x
	}
}

// NewReq creates a new Req request for this client.
func (client Client) NewReq(method, uri string, body io.Reader, mods ...func(*Req)) Req {
	httpReq, _ := http.NewRequest(method, client.Url+uri, body)
	httpReq.Header.Add("Content-Type", "application/json")
	req := Req{
		HttpReq:    httpReq,
		LogPayload: true,
	}
	for _, mod := range mods {
		mod(&req)
	}
	return req
}

// Do makes a request.
// Requests for Do are built ouside of the client, e.g.
//
//	req := client.NewReq("GET", "/appcenter/cisco/ndfc/api/v1/lan-fabric/rest/control/fabrics", nil)
//	res, _ := client.Do(req)
func (client *Client) Do(req Req) (Res, error) {

	var res Res
	defer log.Printf("[DEBUG] Exit from Do method")
	bodyBytes, err := client.doHttpReq(req)
	if err != nil {
		return res, err
	}

	if !json.Valid(bodyBytes) {
		res = Res(gjson.Parse(`{"response": "` + string(bodyBytes) + `"}`))
	} else {
		res = Res(gjson.ParseBytes(bodyBytes))
	}
	if req.LogPayload {
		log.Printf("[DEBUG] HTTP Response: %s", res)
	}
	return res, nil
}

func (client *Client) DoRaw(req Req) ([]byte, error) {
	defer log.Printf("[DEBUG] Exit from DoRaw method")
	bodyBytes, err := client.doHttpReq(req)
	if err != nil {
		return bodyBytes, err
	}
	if req.LogPayload {
		log.Printf("[DEBUG] HTTP Response: %s", string(bodyBytes))
	}
	return bodyBytes, nil
}

func (client *Client) doHttpReq(req Req) ([]byte, error) {
	// add token
	req.HttpReq.Header.Add("Authorization", "Bearer "+client.Token)
	// retain the request body across multiple attempts
	var body []byte
	if req.HttpReq.Body != nil {
		body, _ = io.ReadAll(req.HttpReq.Body)
	}
	var bodyBytes []byte
	defer log.Printf("[DEBUG] Exit from doHttpReq method")
	for attempts := 0; ; attempts++ {
		req.HttpReq.Body = io.NopCloser(bytes.NewBuffer(body))
		if req.LogPayload {
			log.Printf("[DEBUG] HTTP Request: %s, |%s|, |%s|", req.HttpReq.Method, req.HttpReq.URL, req.HttpReq.Body)
		} else {
			log.Printf("[DEBUG] HTTP Request: %s, %s", req.HttpReq.Method, req.HttpReq.URL)
		}

		httpRes, err := client.HttpClient.Do(req.HttpReq)
		if err != nil {
			if ok := client.Backoff(attempts); !ok {
				log.Printf("[ERROR] HTTP Connection error occured: %+v", err)
				return nil, err
			} else {
				log.Printf("[ERROR] HTTP Connection failed: %q, retries: %v", err, attempts)
				continue
			}
		}

		defer httpRes.Body.Close()
		bodyBytes, err = io.ReadAll(httpRes.Body)
		if err != nil {
			if ok := client.Backoff(attempts); !ok {
				log.Printf("[ERROR] Cannot decode response body: %+v", err)
				return nil, err
			} else {
				log.Printf("[ERROR] Cannot decode response body: %s, retries: %v", err, attempts)
				continue
			}
		}

		if httpRes.StatusCode >= 200 && httpRes.StatusCode <= 299 {
			break
		} else {
			if ok := client.Backoff(attempts); !ok {
				log.Printf("[ERROR] HTTP Request failed: StatusCode %v", httpRes.StatusCode)
				return bodyBytes, fmt.Errorf("HTTP Request failed: StatusCode %v", httpRes.StatusCode)
			} else if httpRes.StatusCode == 408 || (httpRes.StatusCode >= 501 && httpRes.StatusCode <= 599) {
				log.Printf("[ERROR] HTTP Request failed: StatusCode %v, Retries: %v", httpRes.StatusCode, attempts)
				continue
			} else if httpRes.StatusCode == 401 && strings.Contains(string(bodyBytes), "token has expired") {
				log.Printf("[ERROR] HTTP Request failed: StatusCode %v, Retries: %v", httpRes.StatusCode, attempts)
				client.Token = ""
				err := client.Authenticate()
				if err != nil {
					log.Printf("[ERROR] Authentication failed: StatusCode %v, Retries: %v", httpRes.StatusCode, attempts)
					return bodyBytes, fmt.Errorf("HTTP Request failed: StatusCode %v", httpRes.StatusCode)
				}
			} else {
				log.Printf("[ERROR] HTTP Request failed: StatusCode %v", httpRes.StatusCode)
				return bodyBytes, fmt.Errorf("HTTP Request failed: StatusCode %v", httpRes.StatusCode)
			}
		}
	}

	return bodyBytes, nil
}

// Get makes a GET request and returns a GJSON result.
// Results will be the raw data structure as returned by Nexus Dashboard
func (client *Client) Get(path string, mods ...func(*Req)) (Res, error) {
	req := client.NewReq("GET", client.BasePath+path, nil, mods...)
	err := client.Authenticate()
	if err != nil {
		return Res{}, err
	}
	return client.Do(req)
}

func (client *Client) GetRawJson(path string, mods ...func(*Req)) ([]byte, error) {
	log.Printf("Staring get request %v", path)
	req := client.NewReq("GET", client.BasePath+path, nil, mods...)
	err := client.Authenticate()
	if err != nil {
		return nil, err
	}
	return client.DoRaw(req)
}

// Delete makes a DELETE request and returns a GJSON result.
// Hint: Use the Body struct to easily create DELETE body data.
func (client *Client) Delete(path string, data string, mods ...func(*Req)) (Res, error) {
	req := client.NewReq("DELETE", client.BasePath+path, strings.NewReader(data), mods...)
	err := client.Authenticate()
	if err != nil {
		return Res{}, err
	}
	return client.Do(req)
}

// Post makes a POST request and returns a GJSON result.
// Hint: Use the Body struct to easily create POST body data.
func (client *Client) Post(path, data string, mods ...func(*Req)) (Res, error) {
	req := client.NewReq("POST", client.BasePath+path, strings.NewReader(data), mods...)
	err := client.Authenticate()
	if err != nil {
		return Res{}, err
	}
	return client.Do(req)
}

// Put makes a PUT request and returns a GJSON result.
// Hint: Use the Body struct to easily create PUT body data.
func (client *Client) Put(path, data string, mods ...func(*Req)) (Res, error) {
	req := client.NewReq("PUT", client.BasePath+path, strings.NewReader(data), mods...)
	err := client.Authenticate()
	if err != nil {
		return Res{}, err
	}
	return client.Do(req)
}

// Login authenticates to the Nexus Dashboard instance.
func (client *Client) Login() error {
	body := ""
	body, _ = sjson.Set(body, "userName", client.Usr)
	body, _ = sjson.Set(body, "userPasswd", client.Pwd)
	body, _ = sjson.Set(body, "domain", client.Domain)
	req := client.NewReq("POST", "/login", strings.NewReader(body), NoLogPayload)
	log.Printf("[TRACE] Client Login: starting http request")
	httpRes, err := client.HttpClient.Do(req.HttpReq)
	if err != nil {
		log.Printf("[ERROR] Client Login:HTTP request failed - %v", err)
		return err
	}
	defer httpRes.Body.Close()
	if httpRes.StatusCode != 200 {
		log.Printf("[ERROR] Authentication failed: StatusCode %v", httpRes.StatusCode)
		return fmt.Errorf("Authentication failed")
	}
	bodyBytes, _ := io.ReadAll(httpRes.Body)
	res := Res(gjson.ParseBytes(bodyBytes))
	token := res.Get("token").String()
	if token == "" {
		log.Printf("[ERROR] Token retrieval failed: no token in payload")
		return fmt.Errorf("Authentication failed")
	}
	client.Token = token
	client.AuthTimeStamp = time.Now()
	log.Printf("[DEBUG] Authentication successful")
	return nil
}

// Login if no token available or token timeout has reached
func (client *Client) Authenticate() error {
	var err error
	log.Printf("[TRACE] Attempting authentication...")
	client.AuthenticationMutex.Lock()
	loginNeeded := false
	if client.Token == "" {
		log.Printf("[DEBUG] No token available, attempting login...")
		loginNeeded = true
	} else if time.Since(client.AuthTimeStamp) > client.AuthTokenTimeout {
		log.Printf("[DEBUG] Token has expired, attempting login...")
		loginNeeded = true
	}
	if loginNeeded {
		err = client.Login()
	}
	log.Printf("[TRACE] Authentication complete")
	client.AuthenticationMutex.Unlock()
	return err
}

// Backoff waits following an exponential backoff algorithm
func (client *Client) Backoff(attempts int) bool {
	log.Printf("[DEBUG] Begining backoff method: attempts %v on %v", attempts, client.MaxRetries)
	if attempts >= client.MaxRetries {
		log.Printf("[DEBUG] Exit from backoff method with return value false")
		return false
	}

	minDelay := time.Duration(client.BackoffMinDelay) * time.Second
	maxDelay := time.Duration(client.BackoffMaxDelay) * time.Second

	min := float64(minDelay)
	backoff := min * math.Pow(client.BackoffDelayFactor, float64(attempts))
	if backoff > float64(maxDelay) {
		backoff = float64(maxDelay)
	}
	backoff = (rand.Float64()/2+0.5)*(backoff-min) + min
	backoffDuration := time.Duration(backoff)
	log.Printf("[TRACE] Starting sleeping for %v", backoffDuration.Round(time.Second))
	time.Sleep(backoffDuration)
	log.Printf("[DEBUG] Exit from backoff method with return value true")
	return true
}

func (client *Client) checkAndFillTokenTimeout() bool {

	req := client.NewReq("GET", "/api/config/dn/apigwcfg/default", nil, NoLogPayload)
	err := client.Authenticate()
	if err != nil {
		log.Printf("[ERROR] Get API Config: Authentication failed - %v", err)
		return false
	}
	result, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] Get API Config: %v", err)
		return false
	}
	/* Reposne Format
			{
			"config": {
				"idle_session_timeout_sec": 3600,
				"jwt_session_timeout_sec": 1200,
				"log_level": "info"
			},
	    }
	*/
	tokenTimeout := result.Get("config.jwt_session_timeout_sec").Int()
	if tokenTimeout > 0 {
		client.AuthTokenTimeout = (time.Duration(tokenTimeout/2) * time.Second)
		return true
	}
	log.Printf("[ERROR] Token timeout could not be read %v, using default value", result)
	return false
}
