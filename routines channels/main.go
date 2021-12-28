package main

import (
	"fmt"
	"net/http"
)

func testLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!", err)
		c <- " hello"
		return
	}
	fmt.Println(link, "up and running successfully!")
	c <- " how are you"
}

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://www.yahoo.com",
		"http://www.manabadi.com",
	}

	// declaring a chanel

	c := make(chan string)

	for _, link := range links {
		go testLink(link, c)
	}
	for range links {
		fmt.Println(<-c)
	}

}
