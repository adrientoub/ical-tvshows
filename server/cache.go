package server

import (
	"github.com/adrientoub/ical-tvshows/config"
	"gopkg.in/redis.v4"
	"log"
	"time"
)

var redis_client *redis.Client

func redisClient() *redis.Client {
	if redis_client != nil {
		return redis_client
	}

	config := config.GetConfig()
	if config["redis"] == nil {
		log.Fatal("No redis configuration")
	}
	redis_config := config["redis"].(map[string]interface{})
	db := redis_config["db"].(float64)
	address := redis_config["address"].(string)

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: redis_config["password"].(string),
		DB:       int(db),
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Impossible to connect to Redis client.", err)
	} else {
		log.Printf("Connection to Redis at `%s' successful.", address)
	}
	redis_client = client
	return redis_client
}

func StoreInCache(key string, value string, ttl int) {
	// TODO: add more caching systems
	err := redisClient().Set(key, value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		log.Printf("Error (%v) storing `%s'", err, key)
	} else {
		log.Printf("Stored key: %s in Redis", key)
	}
}

func GetFromCache(key string) *string {
	// TODO: add more caching systems
	val, err := redisClient().Get(key).Result()
	if err != nil {
		return nil
	} else {
		log.Printf("Got key: %s from Redis", key)
		return &val
	}
}
