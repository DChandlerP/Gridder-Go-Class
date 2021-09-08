package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.golang.org",
		"https://www.stackoverflow.com",
		"https://www.amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down!"
		return
	}
	c <- link + " is up!"
}
