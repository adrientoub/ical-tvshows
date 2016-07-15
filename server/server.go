package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Listen(uri string) {
	// Use printIcs to handle requests to root
	http.HandleFunc("/", printIcs)
	err := http.ListenAndServe(uri, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func printIcs(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username string
	for k, v := range r.Form {
		if k == "username" {
			username = strings.Join(v, "")
		}
	}

	resp, err := GetOriginalIcs()
	if err != nil {
		return
	}
	shows := GetShowList(username)
	cal, err := GetFilteredIcs(resp, shows)
	if err != nil {
		return
	}
	fmt.Fprintf(w, cal)
}
