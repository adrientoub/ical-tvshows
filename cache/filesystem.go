package cache

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

const cachingFolder = "ical-cache"

func initFilesystem() {
	err := os.MkdirAll(cachingFolder, 0755)
	if err != nil {
		log.Fatal("Error creating folder:", err)
	}
}

func createCacheKey(key string) string {
	return cachingFolder + "/" + key
}

func storeInFiles(key string, value string, ttl int) {
	cacheKey := createCacheKey(key)
	file, err := os.Create(cacheKey)
	if err != nil {
		log.Printf("Error (%v) creating `%s' in cache.", err, key)
	}
	_, err = file.WriteString(value)
	if err != nil {
		log.Printf("Error (%v) writing to cache.", err)
	}

	// Create a ttl file
	file, err = os.Create(cacheKey + "--ttl")
	if err != nil {
		log.Printf("Error (%v) creating ttl file for `%s' in cache.", err, key)
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", ttl))
	if err != nil {
		log.Printf("Error (%v) writing to ttl file.", err)
	} else {
		log.Printf("Created cache for `%s'", key)
	}
}

func getTtlFromFile(cacheKey string) *int {
	ttlCacheKey := cacheKey + "--ttl"
	ttlFile, err := ioutil.ReadFile(ttlCacheKey)
	if err != nil {
		log.Printf("Impossible to read file `%s' (%v).", ttlCacheKey, err)
		return nil
	}
	ttl, err := strconv.Atoi(string(ttlFile))
	if err != nil {
		log.Printf("ttl is not a number %v", err)
		return nil
	}
	return &ttl
}

func getFileModTime(filename string) *time.Time {
	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Impossible to read file `%s' (%v).", filename, err)
		return nil
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		log.Printf("Impossible to get file info from `%s' (%v).", filename, err)
		return nil
	}
	modTime := fileInfo.ModTime()
	return &modTime
}

func getFromFiles(key string) *string {
	cacheKey := createCacheKey(key)
	ttl := getTtlFromFile(cacheKey)
	if ttl == nil {
		return nil
	}
	if *ttl != 0 {
		modTime := getFileModTime(cacheKey)
		if modTime == nil {
			return nil
		}
		duration := time.Now().Sub(*modTime)
		if duration.Seconds() > float64(*ttl) {
			return nil
		}
		log.Printf("Got key: %s from filesystem. Expires in %f seconds.", key, duration.Seconds())
	}
	cacheFile, serr := ioutil.ReadFile(cacheKey)
	if serr != nil {
		return nil
	}
	stringed := string(cacheFile)
	return &stringed
}
