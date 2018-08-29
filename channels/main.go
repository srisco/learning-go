package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// for range links {
	// 	fmt.Println(<-c)
	// 	We can only wait without print the content of the channel
	// 	<-c
	// }

	for l := range c {
		// Using a function literal (Anonymous function)
		go func(link string) {
			// Sleep the go routine for 5 seconds
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
