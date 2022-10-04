package main

import (
	"log"
	"net/http"
)

func main() {
	simpleHttpGet("https://google.com")
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}

	return
}
