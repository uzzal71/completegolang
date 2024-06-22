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
		"http://go.dev/",
		"http://amazon.com",
	}

	ch := make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for l := range ch {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, ch)
		}(l)
	}

	anotherLinks := []string{
		"http://anandabazar.com",
		"http://prothomalo.com",
		"http://twitter.com",
		"http://github.com",
		"http://thedailystar.com",
	}

	c := make(chan string)

	for _, link := range anotherLinks {
		go anotherCheckLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		ch <- link
		return
	} else {
		fmt.Println(link, "is up")
		ch <- link
	}
}

func anotherCheckLink(link string, ch chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		ch <- "Might be down I think"
		return
	} else {
		fmt.Println(link, "is up")
		ch <- "Yeo its up"
	}
}
