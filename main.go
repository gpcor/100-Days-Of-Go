package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{"https://google.com/", "https://facebook.com/", "https://twitter.com/", "https://github.com/"}

	for _, link := range sites {
		go siteRequest(link)
	}
}

func siteRequest(s string) {
	_, err := http.Get(s)
	if err != nil {
		fmt.Println("Site may be down: ", s)
		return
	}

	fmt.Println(s, "is up!")
}
