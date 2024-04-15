package cache

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/coocood/freecache"
)

type freeCache struct {
	mu    sync.Mutex
	cache *freecache.Cache
}

var _ Cache = (*freeCache)(nil)

func NewFreeCache(maxMemorySize string) (Cache, error) {
	size, err := parseUnit(maxMemorySize)
	if err != nil {
		return nil, err
	}
	return &freeCache{
		cache: freecache.NewCache(size),
	}, nil
}

func (f *freeCache) SetMaxMemory(_ string) bool {
	return false
}

func (f *freeCache) Set(key string, val any, expire time.Duration) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}
	err = f.cache.Set([]byte(key), b, int(expire.Seconds()))
	return err
}

func (f *freeCache) Get(key string, result any) error {
	value, err := f.cache.Get([]byte(key))
	if err != nil {
		return err
	}
	if err := json.Unmarshal(value, result); err != nil {
		return err
	}
	return nil
}

func (f *freeCache) Del(key string) bool {
	ok := f.cache.Del([]byte(key))
	return ok
}

func (f *freeCache) Exists(key string) bool {
	_, err := f.cache.Get([]byte(key))
	return err == nil
}

func (f *freeCache) Clear() bool {
	f.cache.Clear()
	return true
}

func (f *freeCache) Keys() int64 {
	count := f.cache.EntryCount()
	return count
}
