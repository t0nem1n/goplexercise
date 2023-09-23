package exer8_6

import (
	"log"

	linkLib "github.com/t0nem1n/goplexercise/ch8/exer8_6/link"
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

func Crawl(urls []string, maxDepth int, numWorker int) {
	queue := make(chan WorkList)
	go func() {
		queue <- NewWorkList(urls, 0)
	}()

	seen := make(map[string]bool)

	unseenLink := make(chan Link)

	for i := 0; i < numWorker; i++ {
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
