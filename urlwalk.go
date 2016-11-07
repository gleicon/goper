package main

import (
	"net/http"
	"time"
)

func UrlWalk(argUrl string) {
	var httpClient = &http.Client{
		Timeout: time.Second * 5,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	for {

		response, err := httpClient.Get(argUrl)
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
		currentColor.Printf(">> %s (%d)\n", argUrl, response.StatusCode)

		l, err := response.Location()
		if l == nil {
			return
		}
		argUrl = l.String()
	}
}
