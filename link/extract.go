package link

import (
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var client = http.Client{
	Timeout: 2 * time.Second,
}

func forEachNode(node *html.Node, visitFunc func(node *html.Node)) {
	if node == nil {
		return
	}
	visitFunc(node)
	forEachNode(node.FirstChild, visitFunc)
	forEachNode(node.NextSibling, visitFunc)
}

func Extract(link string) ([]string, error) {
	resp, err := client.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	page, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}
	subLinks := make([]string, 0)
	visitNode := func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					if l, err := resp.Request.URL.Parse(a.Val); err == nil {
						subLinks = append(subLinks, l.String())
					}
				}
			}
		}
	}
	forEachNode(page, visitNode)
	return subLinks, nil
}
