package server

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/adrientoub/ical-tvshows/cache"
)

const url = "https://calendar.google.com/calendar/ical/td5o0neuo68pb2ush6bun9inu8%40group.calendar.google.com/public/basic.ics"

func getIcsFromInternet() (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print("Impossible to read from source: ", err)
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Impossible to read: ", err)
		return "", err
	}
	return string(body), nil
}

func GetOriginalIcs() (string, error) {
	cacheKey := "seriesBaseICS"
	ics := cache.GetFromCache(cacheKey)
	if ics == nil {
		internetIcs, err := getIcsFromInternet()
		if err == nil {
			cache.StoreInCache(cacheKey, internetIcs, 15*60)
		}
		return internetIcs, err
	}
	return *ics, nil
}
