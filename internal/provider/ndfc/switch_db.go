// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package ndfc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

// SwitchEntry represents essential switch information
type SwitchEntry struct {
	SerialNumber string
	IPAddress    string
	SwitchName   string
	SwitchRole   string
}

// SwitchLoaderFunc is the function signature for loading switch data from NDFC
type SwitchLoaderFunc func(context.Context, string) ([]byte, error)

// SwitchDB is a thread-safe database for switch information
// Structure: fabric -> IP -> SwitchEntry
type SwitchDB struct {
	mu     sync.RWMutex
	data   map[string]map[string]SwitchEntry // fabric -> IP -> SwitchEntry
	loader SwitchLoaderFunc                  // function to load switch data from NDFC
}

// SwitchInventoryItem represents a single switch from the NDFC inventory API response
type SwitchInventoryItem struct {
	IPAddress    string `json:"ipAddress"`
	SerialNumber string `json:"serialNumber"`
	FabricName   string `json:"fabricName"`
	LogicalName  string `json:"logicalName"`
	HostName     string `json:"hostName"`
	SwitchRole   string `json:"switchRole"`
}

// NewSwitchDB creates a new SwitchDB instance with an optional loader function
func NewSwitchDB(loader SwitchLoaderFunc) *SwitchDB {
	log.Printf("[DEBUG] SwitchDB: Creating new SwitchDB instance")
	return &SwitchDB{
		data:   make(map[string]map[string]SwitchEntry),
		loader: loader,
	}
}

// SetLoader sets the loader function for fetching switch data
func (db *SwitchDB) SetLoader(loader SwitchLoaderFunc) {
	log.Printf("[DEBUG] SwitchDB: Setting loader function")
	db.loader = loader
}

// LoadFromJSON parses JSON inventory data and populates the SwitchDB for a given fabric
func (db *SwitchDB) LoadFromJSON(fabricName string, jsonData []byte) error {
	var switches []SwitchInventoryItem

	if err := json.Unmarshal(jsonData, &switches); err != nil {
		return fmt.Errorf("failed to unmarshal switch inventory: %w", err)
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	// Initialize fabric map if not exists
	if db.data[fabricName] == nil {
		db.data[fabricName] = make(map[string]SwitchEntry)
	}

	for _, sw := range switches {
		if sw.IPAddress == "" {
			continue // Skip entries without IP
		}

		entry := SwitchEntry{
			SerialNumber: sw.SerialNumber,
			IPAddress:    sw.IPAddress,
			SwitchName:   sw.LogicalName,
			SwitchRole:   sw.SwitchRole,
		}
		// Use hostname if logical name is empty
		if entry.SwitchName == "" {
			entry.SwitchName = sw.HostName
		}

		db.data[fabricName][sw.IPAddress] = entry
		log.Printf("[DEBUG] SwitchDB: Added switch IP=%s Serial=%s Fabric=%s",
			sw.IPAddress, sw.SerialNumber, fabricName)
	}

	log.Printf("[DEBUG] SwitchDB: Loaded %d switches for fabric %s", len(db.data[fabricName]), fabricName)
	return nil
}

// GetSerialByIP retrieves serial number given fabric name and IP address
// Auto-loads fabric data if not present and loader is configured
func (db *SwitchDB) GetSerialByIP(ctx context.Context, fabricName, ipAddress string) (string, bool) {
	log.Printf("[DEBUG] SwitchDB: GetSerialByIP called for fabric=%s ip=%s", fabricName, ipAddress)
	// Try cache first
	db.mu.RLock()
	if fabricMap, ok := db.data[fabricName]; ok {
		if entry, ok := fabricMap[ipAddress]; ok {
			db.mu.RUnlock()
			log.Printf("[DEBUG] SwitchDB: Cache hit for fabric=%s ip=%s serial=%s", fabricName, ipAddress, entry.SerialNumber)
			return entry.SerialNumber, true
		}
	}
	db.mu.RUnlock()
	log.Printf("[DEBUG] SwitchDB: Cache miss for fabric=%s ip=%s, attempting to load", fabricName, ipAddress)

	// Not in cache, try to load if loader is available
	if err := db.ensureLoaded(ctx, fabricName); err != nil {
		log.Printf("[WARN] SwitchDB: Failed to load fabric %s: %v", fabricName, err)
		return "", false
	}

	// Try again after loading
	db.mu.RLock()
	defer db.mu.RUnlock()
	if fabricMap, ok := db.data[fabricName]; ok {
		if entry, ok := fabricMap[ipAddress]; ok {
			log.Printf("[DEBUG] SwitchDB: Found after load for fabric=%s ip=%s serial=%s", fabricName, ipAddress, entry.SerialNumber)
			return entry.SerialNumber, true
		}
	}
	log.Printf("[DEBUG] SwitchDB: Not found for fabric=%s ip=%s", fabricName, ipAddress)
	return "", false
}

// GetSwitchByIP retrieves full switch entry given fabric name and IP address
// Auto-loads fabric data if not present and loader is configured
func (db *SwitchDB) GetSwitchByIP(ctx context.Context, fabricName, ipAddress string) (SwitchEntry, bool) {
	log.Printf("[DEBUG] SwitchDB: GetSwitchByIP called for fabric=%s ip=%s", fabricName, ipAddress)
	// Try cache first
	db.mu.RLock()
	if fabricMap, ok := db.data[fabricName]; ok {
		if entry, ok := fabricMap[ipAddress]; ok {
			db.mu.RUnlock()
			log.Printf("[DEBUG] SwitchDB: Cache hit for fabric=%s ip=%s serial=%s", fabricName, ipAddress, entry.SerialNumber)
			return entry, true
		}
	}
	db.mu.RUnlock()
	log.Printf("[DEBUG] SwitchDB: Cache miss for fabric=%s ip=%s, attempting to load", fabricName, ipAddress)

	// Not in cache, try to load if loader is available
	if err := db.ensureLoaded(ctx, fabricName); err != nil {
		log.Printf("[WARN] SwitchDB: Failed to load fabric %s: %v", fabricName, err)
		return SwitchEntry{}, false
	}

	// Try again after loading
	db.mu.RLock()
	defer db.mu.RUnlock()
	if fabricMap, ok := db.data[fabricName]; ok {
		if entry, ok := fabricMap[ipAddress]; ok {
			log.Printf("[DEBUG] SwitchDB: Found after load for fabric=%s ip=%s serial=%s", fabricName, ipAddress, entry.SerialNumber)
			return entry, true
		}
	}
	log.Printf("[DEBUG] SwitchDB: Not found for fabric=%s ip=%s", fabricName, ipAddress)
	return SwitchEntry{}, false
}

// ensureLoaded loads fabric data if not already loaded and loader is configured
func (db *SwitchDB) ensureLoaded(ctx context.Context, fabricName string) error {
	log.Printf("[DEBUG] SwitchDB: ensureLoaded called for fabric=%s", fabricName)
	if db.loader == nil {
		log.Printf("[DEBUG] SwitchDB: No loader configured")
		return fmt.Errorf("no loader configured")
	}
	// Check if already loaded (with write lock to prevent double-loading)
	db.mu.Lock()
	if _, ok := db.data[fabricName]; ok {
		db.mu.Unlock()
		log.Printf("[DEBUG] SwitchDB: Fabric %s already loaded", fabricName)
		return nil
	}
	db.mu.Unlock()

	// Load the data
	log.Printf("[DEBUG] SwitchDB: Loading fabric %s from NDFC", fabricName)
	jsonData, err := db.loader(ctx, fabricName)
	if err != nil {
		log.Printf("[DEBUG] SwitchDB: Loader error for fabric %s: %v", fabricName, err)
		return err
	}
	log.Printf("[DEBUG] SwitchDB: Received %d bytes for fabric %s", len(jsonData), fabricName)

	return db.LoadFromJSON(fabricName, jsonData)
}

// IsFabricLoaded checks if a fabric's switch data has been loaded
func (db *SwitchDB) IsFabricLoaded(fabricName string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	_, ok := db.data[fabricName]
	log.Printf("[DEBUG] SwitchDB: IsFabricLoaded fabric=%s loaded=%v", fabricName, ok)
	return ok
}

// ClearFabric removes all switch data for a fabric (useful for refresh)
func (db *SwitchDB) ClearFabric(fabricName string) {
	log.Printf("[DEBUG] SwitchDB: ClearFabric called for fabric=%s", fabricName)
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, fabricName)
	log.Printf("[DEBUG] SwitchDB: Cleared fabric=%s", fabricName)
}

// GetIPBySerial retrieves IP address given fabric name and serial number (reverse lookup)
func (db *SwitchDB) GetIPBySerial(fabricName, serialNumber string) (string, bool) {
	log.Printf("[DEBUG] SwitchDB: GetIPBySerial called for fabric=%s serial=%s", fabricName, serialNumber)
	db.mu.RLock()
	defer db.mu.RUnlock()

	if fabricMap, ok := db.data[fabricName]; ok {
		for ip, entry := range fabricMap {
			if entry.SerialNumber == serialNumber {
				log.Printf("[DEBUG] SwitchDB: Found ip=%s for fabric=%s serial=%s", ip, fabricName, serialNumber)
				return ip, true
			}
		}
	}
	log.Printf("[DEBUG] SwitchDB: Not found for fabric=%s serial=%s", fabricName, serialNumber)
	return "", false
}

// EnsureFabricLoaded loads fabric switch data if not already loaded
// Uses the configured loader function
func (db *SwitchDB) EnsureFabricLoaded(ctx context.Context, fabricName string) error {
	log.Printf("[DEBUG] SwitchDB: EnsureFabricLoaded called for fabric=%s", fabricName)
	return db.ensureLoaded(ctx, fabricName)
}
