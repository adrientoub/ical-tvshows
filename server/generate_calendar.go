package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const url = "https://calendar.google.com/calendar/ical/td5o0neuo68pb2ush6bun9inu8%40group.calendar.google.com/public/basic.ics"
const cacheFile = "cache.dat"

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

func createCache(content string) {
	fout, err := os.Create(cacheFile)
	if err != nil {
		fmt.Println(cacheFile, err)
		return
	}
	defer fout.Close()
	fout.WriteString(content)
}

func getIcsFromCache() (string, error) {
	content, err := ioutil.ReadFile(cacheFile)
	return string(content), err
}

func GetOriginalIcs() (string, error) {
	ics, err := getIcsFromCache()
	if err != nil {
		ics, err = getIcsFromInternet()
		if err == nil {
			go createCache(ics)
		}
	}
	return ics, err
}
