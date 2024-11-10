package cachedb

type CacheStatus int64

const (
	None CacheStatus = iota
	Select
	Insert
	Update
	Delete
)

type CacheContent interface {
	Table() string
	UniqueKeyColumnValueStr() string
	SetCacheStatus(s CacheStatus)
	GetCacheStatus() CacheStatus
	Update(content CacheContent) error
}
