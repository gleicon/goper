package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/fatih/color"
)

var currentColor, red, green, yellow *color.Color

func main() {
	red = color.New(color.FgRed)
	green = color.New(color.FgGreen)
	yellow = color.New(color.FgYellow)

	if len(os.Args) < 2 {
		fmt.Println("Usage: goper <url>")
		os.Exit(-1)
	}

	argUrl := os.Args[1]

	_, err := url.Parse(argUrl)

	if err != nil {
		red.Println(err)
		os.Exit(-2)
	}

	UrlWalk(argUrl)
}
