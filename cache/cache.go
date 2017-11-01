package cache

import (
	"log"

	"../config"
)

type Cache int

const (
	None  Cache = 1 << iota
	Redis Cache = 2 << iota
	Files Cache = 3 << iota
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
	} else if cache == "files" {
		initFilesystem()
		log.Println("Using local Filesystem as cache.")
		return Files
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
	} else if cacheSystem() == Files {
		go storeInFiles(key, value, ttl)
	}
}

func GetFromCache(key string) *string {
	if cacheSystem() == Redis {
		return getFromRedis(key)
	} else if cacheSystem() == Files {
		return getFromFiles(key)
	}
	return nil
}
