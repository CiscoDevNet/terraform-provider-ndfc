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
	HostName     string
}

// SwitchLoaderFunc is the function signature for loading switch data from NDFC
type SwitchLoaderFunc func(context.Context) ([]byte, error)

// SwitchDB is a thread-safe database for switch information
// Supports lookup by IP address or serial number
type SwitchDB struct {
	mu           sync.RWMutex
	dataByIP     map[string]*SwitchEntry // IP -> SwitchEntry
	dataBySerial map[string]*SwitchEntry // Serial -> SwitchEntry
	loader       SwitchLoaderFunc        // function to load switch data from NDFC
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
		dataByIP:     make(map[string]*SwitchEntry),
		dataBySerial: make(map[string]*SwitchEntry),
		loader:       loader,
	}
}

// SetLoader sets the loader function for fetching switch data
func (db *SwitchDB) SetLoader(loader SwitchLoaderFunc) {
	log.Printf("[DEBUG] SwitchDB: Setting loader function")
	db.loader = loader
}

// LoadFromJSON parses JSON inventory data and populates the SwitchDB for a given fabric
func (db *SwitchDB) LoadFromJSON(jsonData []byte) error {
	var switches []SwitchInventoryItem

	if err := json.Unmarshal(jsonData, &switches); err != nil {
		return fmt.Errorf("failed to unmarshal switch inventory: %w", err)
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	// Initialize fabric map if not exists
	if db.dataByIP == nil {
		db.dataByIP = make(map[string]*SwitchEntry)
	}

	if db.dataBySerial == nil {
		db.dataBySerial = make(map[string]*SwitchEntry)
	}

	skippedNoIP := 0
	skippedExists := 0

	for _, sw := range switches {
		if sw.IPAddress == "" {
			skippedNoIP++
			continue // Skip entries without IP
		}
		// skip entries that are already there; we don't expect changes in the same session
		if _, ok := db.dataByIP[sw.IPAddress]; ok {
			skippedExists++
			continue
		}

		entry := SwitchEntry{
			SerialNumber: sw.SerialNumber,
			IPAddress:    sw.IPAddress,
			SwitchName:   sw.LogicalName,
			SwitchRole:   sw.SwitchRole,
			HostName:     sw.HostName,
		}
		// Use hostname if logical name is empty
		if entry.SwitchName == "" {
			entry.SwitchName = sw.HostName
		}

		db.dataByIP[sw.IPAddress] = &entry
		db.dataBySerial[sw.SerialNumber] = &entry
		log.Printf("[DEBUG] SwitchDB: Added switch IP=%s Serial=%s",
			sw.IPAddress, sw.SerialNumber)
	}

	log.Printf("[DEBUG] SwitchDB: Loaded %d switches, skipped %d no IP, %d exists", len(db.dataByIP), skippedNoIP, skippedExists)
	return nil
}

// GetSwitchByIP retrieves full switch entry given IP address
// Auto-loads data if not present and loader is configured
// fabricName is kept for backward compatibility but ignored
func (db *SwitchDB) GetSwitchByIP(ctx context.Context, fabricName, ipAddress string) (*SwitchEntry, bool) {
	log.Printf("[DEBUG] SwitchDB: GetSwitchByIP called for ip=%s", ipAddress)
	// Try cache first
	db.mu.RLock()
	if entry, ok := db.dataByIP[ipAddress]; ok {
		db.mu.RUnlock()
		log.Printf("[DEBUG] SwitchDB: Cache hit for ip=%s serial=%s", ipAddress, entry.SerialNumber)
		return entry, true
	}
	db.mu.RUnlock()
	log.Printf("[DEBUG] SwitchDB: Cache miss for ip=%s, attempting to load", ipAddress)

	// Not in cache, try to load if loader is available
	if err := db.ensureLoaded(ctx); err != nil {
		log.Printf("[WARN] SwitchDB: Failed to load: %v", err)
		return nil, false
	}

	// Try again after loading
	db.mu.RLock()
	defer db.mu.RUnlock()
	if entry, ok := db.dataByIP[ipAddress]; ok {
		log.Printf("[DEBUG] SwitchDB: Found after load for ip=%s serial=%s", ipAddress, entry.SerialNumber)
		return entry, true
	}
	log.Printf("[DEBUG] SwitchDB: Not found for ip=%s", ipAddress)
	return nil, false
}

// ensureLoaded loads fabric data if not already loaded and loader is configured
func (db *SwitchDB) ensureLoaded(ctx context.Context) error {
	log.Printf("[DEBUG] SwitchDB: ensureLoaded called")
	if db.loader == nil {
		log.Printf("[DEBUG] SwitchDB: No loader configured")
		return fmt.Errorf("no loader configured")
	}
	// Load the data
	log.Printf("[DEBUG] SwitchDB: Loading data from NDFC")
	jsonData, err := db.loader(ctx)
	if err != nil {
		log.Printf("[DEBUG] SwitchDB: Loader error", err)
		return err
	}
	log.Printf("[DEBUG] SwitchDB: Received %d bytes", len(jsonData))

	return db.LoadFromJSON(jsonData)
}

// IsFabricLoaded checks if switch data has been loaded
// fabricName is kept for backward compatibility but ignored
func (db *SwitchDB) IsFabricLoaded(fabricName string) bool {
	db.mu.RLock()
	defer db.mu.RUnlock()

	loaded := len(db.dataByIP) > 0
	log.Printf("[DEBUG] SwitchDB: IsLoaded=%v count=%d", loaded, len(db.dataByIP))
	return loaded
}

// ClearFabric removes all switch data (useful for refresh)
// fabricName is kept for backward compatibility but ignored
func (db *SwitchDB) ClearFabric(fabricName string) {
	log.Printf("[DEBUG] SwitchDB: Clear called")
	db.mu.Lock()
	defer db.mu.Unlock()

	db.dataByIP = make(map[string]*SwitchEntry)
	db.dataBySerial = make(map[string]*SwitchEntry)
	log.Printf("[DEBUG] SwitchDB: Cleared all data")
}

// GetIPBySerial retrieves IP address given serial number (reverse lookup)
// fabricName is kept for backward compatibility but ignored
func (db *SwitchDB) GetIPBySerial(ctx context.Context, fabricName, serialNumber string) (string, bool) {
	log.Printf("[DEBUG] SwitchDB: GetIPBySerial called for serial=%s", serialNumber)

	db.mu.RLock()
	if entry, ok := db.dataBySerial[serialNumber]; ok {
		db.mu.RUnlock()
		log.Printf("[DEBUG] SwitchDB: Cache hit for serial=%s ip=%s", serialNumber, entry.IPAddress)
		return entry.IPAddress, true
	}
	db.mu.RUnlock()

	log.Printf("[DEBUG] SwitchDB: Cache miss for serial=%s, attempting to load", serialNumber)

	// Not in cache, try to load if loader is available
	if err := db.ensureLoaded(ctx); err != nil {
		log.Printf("[WARN] SwitchDB: Failed to load: %v", err)
		return "", false
	}

	// Try again after loading
	db.mu.RLock()
	defer db.mu.RUnlock()
	if entry, ok := db.dataBySerial[serialNumber]; ok {
		log.Printf("[DEBUG] SwitchDB: Found after load for serial=%s ip=%s", serialNumber, entry.IPAddress)
		return entry.IPAddress, true
	}
	log.Printf("[DEBUG] SwitchDB: Not found for serial=%s", serialNumber)
	return "", false
}

// EnsureFabricLoaded loads switch data if not already loaded
// Uses the configured loader function
// fabricName is kept for backward compatibility but ignored
func (db *SwitchDB) EnsureFabricLoaded(ctx context.Context, fabricName string) error {
	log.Printf("[DEBUG] SwitchDB: EnsureLoaded called")
	return db.ensureLoaded(ctx)
}

// GetSwitchBySerial retrieves full switch entry given serial number
// Auto-loads data if not present and loader is configured
func (db *SwitchDB) GetSwitchBySerial(ctx context.Context, serialNumber string) (*SwitchEntry, bool) {
	log.Printf("[DEBUG] SwitchDB: GetSwitchBySerial called for serial=%s", serialNumber)
	// Try cache first
	db.mu.RLock()
	if entry, ok := db.dataBySerial[serialNumber]; ok {
		db.mu.RUnlock()
		log.Printf("[DEBUG] SwitchDB: Cache hit for serial=%s ip=%s", serialNumber, entry.IPAddress)
		return entry, true
	}
	db.mu.RUnlock()
	log.Printf("[DEBUG] SwitchDB: Cache miss for serial=%s, attempting to load", serialNumber)

	// Not in cache, try to load if loader is available
	if err := db.ensureLoaded(ctx); err != nil {
		log.Printf("[WARN] SwitchDB: Failed to load: %v", err)
		return nil, false
	}

	// Try again after loading
	db.mu.RLock()
	defer db.mu.RUnlock()
	if entry, ok := db.dataBySerial[serialNumber]; ok {
		log.Printf("[DEBUG] SwitchDB: Found after load for serial=%s ip=%s", serialNumber, entry.IPAddress)
		return entry, true
	}
	log.Printf("[DEBUG] SwitchDB: Not found for serial=%s", serialNumber)
	return nil, false
}
