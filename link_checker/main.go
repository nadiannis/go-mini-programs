// Command:
// check [url]

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"
)

const commandInfo = "Available command: check [URL]"

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Printf("Please specify a command.\n\n%s\n", commandInfo)
		return
	}

	command := flag.Args()[0]

	if command == "check" {
		args := flag.Args()[1:]

		if len(args) == 0 {
			fmt.Printf("Invalid argument. There should be a URL.\n\n%s\n", commandInfo)
			return
		}

		url := args[0]
		fmt.Print(checkURL(url))
	} else {
		fmt.Printf("Invalid command: '%s'\n\n%s\n", command, commandInfo)
	}
}

func checkURL(url string) string {
	if !isValidURL(url) {
		return fmt.Sprintf("Invalid URL: '%s'\n", url)
	}

	response, err := http.Get(url)
	if err != nil {
		return err.Error() + "\n"
	}
	defer response.Body.Close()

	info := fmt.Sprintf("%v %v\n\n", response.Request.Method, url)
	info += fmt.Sprintf("%v %v\n", response.Proto, response.Status)

	for key, _ := range response.Header {
		keys := getHeaderKeys()
		if containsString(keys, key) && response.Header.Get(key) != "" {
			info += fmt.Sprintf("%v: %v\n", key, response.Header.Get(key))
		}
	}

	return info
}

func isValidURL(rawURL string) bool {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	return true
}

func getHeaderKeys() []string {
	return []string{
		"Access-Control-Allow-Origin",
		"Age", 
		"Cache-Control",
		"Content-Disposition",
		"Content-Type",
		"Date", 
		"Etag", 
		"Expires",
		"Last-Modified",
		"Pragma",
		"Server",
		"Strict-Transport-Security",
		"X-Content-Type-Options",
		"X-Frame-Options",
	}
}

func containsString(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}