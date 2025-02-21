package cache

import (
	"crave/hub/cmd/model"
	hubModel "crave/hub/cmd/model"
	"time"
)

type CacheLocal struct {
	data map[string]hubModel.WorkCache
}

func NewCache() *CacheLocal {
	return &CacheLocal{
		data: make(map[string]hubModel.WorkCache),
	}
}

// Delete implements ICache.
func (c *CacheLocal) Delete(key string) error {
	panic("unimplemented")
}

// Get implements ICache.
func (c *CacheLocal) Get(key string) (model.WorkCache, bool) {
	value, exist := c.data[key]
	return value, exist
}

// Set implements ICache.
func (c *CacheLocal) Set(key string, value *hubModel.Work, ttl time.Duration) {
	c.data[key] = *value.ToCache()
}
