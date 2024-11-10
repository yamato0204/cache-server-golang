package cachedb

import "sync"

type DBOperationCacheManager struct {
	dbOperationResult map[string]map[string]CacheContent // key: table, uk_info
	mutex             sync.RWMutex
}

func newDBOperationCacheManager() *DBOperationCacheManager {
	return &DBOperationCacheManager{
		dbOperationResult: make(map[string]map[string]CacheContent, 0),
	}
}
