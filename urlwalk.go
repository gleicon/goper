package main

import (
	"net/http"
	"time"
)

// URLWalk walks through all HTTP redirects of a given URL
func URLWalk(argURL string) {
	var httpClient = &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for {

		response, err := httpClient.Get(argURL)
		if err != nil {
			red.Println(err)
			return
		}

		switch {

		case 300 <= response.StatusCode && response.StatusCode <= 305:
			currentColor = yellow
			break
		case 400 <= response.StatusCode && response.StatusCode <= 505:
			currentColor = red
			break
		default:
			currentColor = green
			break
		}
		currentColor.Printf(">> %s (%d)\n", argURL, response.StatusCode)

		l, err := response.Location()
		if err == http.ErrNoLocation {
			return
		}
		argURL = l.String()
	}
}
