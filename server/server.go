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
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	resp, err := GetOriginalIcs()
	if err != nil {
		return
	}
	GetShowList("adrientoub")
	fmt.Fprintf(w, resp)
}
