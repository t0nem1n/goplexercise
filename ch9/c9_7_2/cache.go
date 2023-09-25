package c9_7_2

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Memo struct {
	f    Func
	reqs chan request
}

type Func func(key string) (any, error)

type entry struct {
	res   result
	ready chan bool
}

type request struct {
	key      string
	response chan result
}

type result struct {
	value any
	err   error
}

func NewMemo(f Func) *Memo {
	return &Memo{
		f,
		make(chan request),
	}
}

func (memo *Memo) Get(key string) (any, error) {
	req := request{key, make(chan result)}
	memo.reqs <- req
	res := <-req.response
	return res.value, res.err
}

func (memo *Memo) server() {
	cache := make(map[string]*entry)
	for req := range memo.reqs {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan bool)}
			cache[req.key] = e
			go memo.call(req.key, e)
		}
		go memo.deliver(e, req.response)
	}
}

func (memo *Memo) call(key string, e *entry) {
	e.res.value, e.res.err = memo.f(key)
	close(e.ready)
}

func (memo *Memo) deliver(e *entry, resp chan<- result) {
	<-e.ready
	resp <- e.res
}

func httpGet(url string) (any, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func Run() {
	memo := NewMemo(httpGet)
	go memo.server()
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"http://gopl.io",
	}
	for _, url := range urls {
		start := time.Now()
		res, err := memo.Get(url)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%s %s %d bytes\n", url, time.Since(start).String(), len(res.([]byte)))
	}
}
