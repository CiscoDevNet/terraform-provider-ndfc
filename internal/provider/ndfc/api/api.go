package api

import (
	"log"
	"sync"
	"terraform-provider-ndfc/tfutils/go-nd"

	"github.com/tidwall/gjson"
)

type NDFCAPI interface {
	GetLock() *sync.Mutex
	//ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error)
	GetUrl() string
	PostUrl() string
	PutUrl() string
	DeleteUrl() string
	GetDeleteQP() []string
}

type NDFCAPICommon struct {
	NDFCAPI
	client *nd.Client
}

/*
	func (c NDFCAPICommon) GetLock() *sync.Mutex {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) ProcessResponse(ctx context.Context, res gjson.Result) ([]string, error) {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) GetUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) PostUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) PutUrl() string {
		panic("Not implemented")
	}

	func (c NDFCAPICommon) DeleteUrl() string {
		panic("Not implemented")
	}
*/
func (c NDFCAPICommon) Get() ([]byte, error) {
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	url := c.NDFCAPI.GetUrl()
	log.Printf("Get URL: %s %v\n", url, c.client)
	if c.client == nil {
		log.Printf("************Client is nil********************")
	}
	res, err := c.client.GetRawJson(url)
	if err != nil {
		lock.Unlock()
		return nil, err
	}

	lock.Unlock()
	log.Printf("Finished GET: %s %v\n", c.NDFCAPI.GetUrl(), lock)
	return res, nil
}

func (c NDFCAPICommon) Post(payload []byte) (gjson.Result, error) {
	log.Printf("Post URL: %s\n", c.NDFCAPI.PostUrl())
	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	defer lock.Unlock()
	log.Printf("Post URL acquired lock: %s\n", c.NDFCAPI.PostUrl())
	res, err := c.client.Post(c.NDFCAPI.PostUrl(), string(payload))
	if err != nil {
		return res, err
	}

	return res, nil
}

func (c NDFCAPICommon) Put(payload []byte) (gjson.Result, error) {

	lock := c.NDFCAPI.GetLock()
	lock.Lock()
	defer lock.Unlock()
	res, err := c.client.Put(c.NDFCAPI.PutUrl(), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDFCAPICommon) Delete() (gjson.Result, error) {
	c.NDFCAPI.GetLock().Lock()
	defer c.NDFCAPI.GetLock().Unlock()
	res, err := c.client.DeleteRaw(c.NDFCAPI.DeleteUrl(), c.NDFCAPI.GetDeleteQP())
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c NDFCAPICommon) DeleteWithPayload(payload []byte) (gjson.Result, error) {
	c.NDFCAPI.GetLock().Lock()
	defer c.NDFCAPI.GetLock().Unlock()
	res, err := c.client.Delete(c.NDFCAPI.DeleteUrl(), string(payload))
	if err != nil {
		return res, err
	}
	return res, nil
}
