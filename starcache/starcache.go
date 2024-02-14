package starcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	_cache *starCache
)

const (
	DefCacheExp        time.Duration = 60 * time.Minute
	DefCleanupInterval time.Duration = 10 * time.Minute
)

func Init(defaultExp, cleanupInt string) {

	defaultExpiration, err := time.ParseDuration(defaultExp)
	if err != nil {
		defaultExpiration = DefCacheExp
	}
	cleanupInterval, err := time.ParseDuration(cleanupInt)
	if err != nil {
		cleanupInterval = DefCleanupInterval
	}

	_cache = newCache(defaultExpiration, cleanupInterval)
}

func GetCache() ICache {
	return _cache
}

type starCache struct {
	cache *cache.Cache
}

func newCache(defaultExpiration, cleanupInterval time.Duration) *starCache {

	c := &starCache{}
	c.cache = cache.New(defaultExpiration, cleanupInterval)

	return c
}

func (c *starCache) Set(key string, value interface{}) {

	if c.cache != nil {
		c.cache.SetDefault(key, value)
	}

}

func (c *starCache) SetWithoutExp(key string, value interface{}) {

	if c.cache != nil {
		c.cache.Set(key, value, cache.NoExpiration)
	}

}

func (c *starCache) SetWithExp(key string, value interface{}, exp time.Duration) {

	if c.cache != nil {
		c.cache.Set(key, value, exp)
	}
}

func (c *starCache) Get(key string) (interface{}, bool) {
	if c.cache != nil {
		return c.cache.Get(key)
	}

	return nil, false
}

func (c *starCache) GetString(key string) (string, bool) {

	if c.cache != nil {
		value, found := c.cache.Get(key)

		if !found {
			return "", false
		}
		str, ok := value.(string)

		return str, ok
	}

	return "", false
}

func (c *starCache) Delete(key string) {
	if c.cache != nil {
		c.cache.Delete(key)
	}
}

func (c *starCache) Purge() {
	if c.cache != nil {
		c.cache.Flush()
	}
}
