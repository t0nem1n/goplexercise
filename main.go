package main

import (
	"log"
	"os"

	linkLib "github.com/t0nem1n/goplexercise/link"
)

type Link struct {
	URL   string
	Depth int
}

type WorkList struct {
	Links []Link
}

func NewWorkList(urls []string, depth int) WorkList {
	wl := WorkList{}
	for _, url := range urls {
		wl.Links = append(wl.Links, Link{url, depth})
	}
	return wl
}

func crawl(link string) []string {
	res, err := linkLib.Extract(link)
	if err != nil {
		log.Printf("link: %s, error: %v", link, err)
	}
	return res
}

func main() {
	maxDepth := 2
	queue := make(chan WorkList)
	go func() {
		queue <- NewWorkList(os.Args[1:], 0)
	}()

	seen := make(map[string]bool)

	unseenLink := make(chan Link)

	for i := 0; i < 10; i++ {
		go func(id int) {
			for link := range unseenLink {
				log.Println(id, link)
				subLinks := crawl(link.URL)
				go func(links []string, depth int) {
					queue <- NewWorkList(links, depth+1)
				}(subLinks, link.Depth)
			}
		}(i)
	}

	n := 1

	for ; n > 0; n-- {
		wl := <-queue
		for _, link := range wl.Links {
			if !seen[link.URL] && link.Depth <= maxDepth {
				n++
				seen[link.URL] = true
				unseenLink <- link
			}
		}
	}
}
