package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adrientoub/ical-tvshows/config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const apiBase = "https://api.betaseries.com/"
const searchEndpoint = "members/search?"
const infosEndpoint = "members/infos?only=shows"

func getIdFromSearchBetaseries(body []byte) (int, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal(body, &parsed)
	if err != nil {
		return 0, err
	}
	users := parsed["users"].([]interface{})
	if len(users) <= 0 {
		return 0, errors.New("No user found")
	}
	user := users[0].(map[string]interface{})
	id, ok := user["id"].(float64)
	if !ok {
		return 0, errors.New("Unable to cast to float64")
	}
	return int(id), nil
}

func getBetaseriesUserId(username string, apiKey string) (int, error) {
	url := fmt.Sprintf("%s%slogin=%s&key=%s", apiBase, searchEndpoint, username, apiKey)
	fmt.Println("GET " + url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return 0, errors.New("Error code: " + resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Impossible to read: ", err)
		return 0, err
	}
	return getIdFromSearchBetaseries(body)
}

func getTitlesFromInfosBetaseries(body []byte) ([]string, error) {
	var parsed map[string]interface{}
	err := json.Unmarshal(body, &parsed)
	if err != nil {
		return []string{}, err
	}
	member := parsed["member"].(map[string]interface{})
	shows := member["shows"].([]interface{})
	fmt.Println(len(shows))
	titles := make([]string, len(shows))
	for i, show := range shows {
		title := show.(map[string]interface{})["title"].(string)
		titles[i] = title
	}
	fmt.Println(titles)

	return titles, nil
}

func getShowListFromBetaseries(id int, apiKey string) ([]string, error) {
	url := fmt.Sprintf("%s%s&id=%d&key=%s", apiBase, infosEndpoint, id, apiKey)
	fmt.Println("GET " + url)
	resp, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	if resp.StatusCode != 200 {
		return []string{}, errors.New("Error code: " + resp.Status)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("Impossible to read: ", err)
		return []string{}, err
	}

	return getTitlesFromInfosBetaseries(body)
}

func getShowListFromInternet(username string, apiKey string) ([]string, error) {
	id, err := getBetaseriesUserId(username, apiKey)
	if err != nil {
		return []string{}, err
	}
	fmt.Printf("Username: `%s' Id: %d\n", username, id)

	return getShowListFromBetaseries(id, apiKey)
}

func GetShowList(username string) []string {
	apiKey := config.GetConfig()["api_key"].(string)

	cacheKey := "showList-" + username
	cached := GetFromCache(cacheKey)
	if cached != nil {
		return strings.Split(*cached, "\n")
	}
	shows, err := getShowListFromInternet(username, apiKey)
	if err != nil {
		log.Print("Impossible to get shows: ", err)
		return []string{}
	}
	StoreInCache(cacheKey, strings.Join(shows, "\n"), 3600)

	return shows
}
