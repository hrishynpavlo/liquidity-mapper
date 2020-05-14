package common

import (
	"github.com/patrickmn/go-cache"
)

//Cache in pointer of memory cache
var memcache *cache.Cache

//CacheInit creates cache instance
func CacheInit() *cache.Cache {
	memcache := cache.New(cache.NoExpiration, cache.NoExpiration)
	return memcache
}

//GetCacheInstance returns
func GetCacheInstance() *cache.Cache {
	return memcache
}
