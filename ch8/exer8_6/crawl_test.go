package exer8_6

import "testing"

func TestCrawl(t *testing.T) {
	urls := []string{"http://www.gopl.io/"}
	depth := 2
	worker := 20
	Crawl(urls, depth, worker)
}
