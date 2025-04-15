package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "chatgpt.com"}
	for i := 0; i < len(urls); i++ {
		targetUrl := urls[i]

		// check if we have a prefix, if not, append https
		if !strings.HasPrefix(targetUrl, "https://") && !strings.HasPrefix(targetUrl, "http://") {
			targetUrl = "https://" + targetUrl
		}
		resp, err := http.Head(targetUrl)
		if err != nil {
			fmt.Println("Error attempting to reach " + targetUrl)
		} else {
			// message based on status code
			if resp.Status == "200 OK" {
				fmt.Println("Success in reaching " + targetUrl)
			}
		}
	}
}
