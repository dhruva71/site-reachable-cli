package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	urls := []string{"https://www.google.com", "https://www.facebook.com", "chatgpt.com"}
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	for i := 0; i < len(urls); i++ {
		targetUrl := urls[i]

		// check if we have a prefix, if not, append https
		if !strings.HasPrefix(targetUrl, "https://") && !strings.HasPrefix(targetUrl, "http://") {
			targetUrl = "https://" + targetUrl
		}
		req, err := http.NewRequest("GET", targetUrl, nil)
		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Accept", `text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8`)
		req.Header.Add("User-Agent", `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_5) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.64 Safari/537.11`)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error attempting to reach " + targetUrl)
		} else {
			// message based on status code
			if resp.Status == "200 OK" {
				fmt.Println("Success in reaching " + targetUrl)
			} else {
				fmt.Println("Failed to reach " + targetUrl + "; status code: " + resp.Status)
			}
		}
	}
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	// TODO: improve this
	return nil
}
