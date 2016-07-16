package cache

import (
	"log"

	"github.com/adrientoub/ical-tvshows/config"
)

type Cache int

const (
	None  Cache = 1 << iota
	Redis Cache = 2 << iota
)

var cacheSystemConfig Cache

func cacheSystemFromConfig() Cache {
	config := config.GetConfig()
	if config["cache"] == nil {
		log.Println("Using no cache.")
		return None
	}
	cache := config["cache"].(string)
	if cache == "redis" {
		log.Println("Using Redis as cache.")
		return Redis
	} else {
		log.Println("Using no cache.")
		return None
	}
}

func cacheSystem() Cache {
	if cacheSystemConfig == 0 {
		cacheSystemConfig = cacheSystemFromConfig()
	}
	return cacheSystemConfig
}

func StoreInCache(key string, value string, ttl int) {
	if cacheSystem() == Redis {
		storeInRedis(key, value, ttl)
	}
}

func GetFromCache(key string) *string {
	if cacheSystem() == Redis {
		return getFromRedis(key)
	} else {
		return nil
	}
}
