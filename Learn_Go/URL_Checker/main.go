package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	var results = map[string]string{}
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		res := "OK"
		err := checkURL(url)
		if err != nil {
			res = "FAIL"
		}
		results[url] = res
	}

	fmt.Println("--------------------------------")

	for url, res := range results {
		fmt.Println(url, res)
	}
}

func checkURL(url string) error {
	fmt.Println("Checking...", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
