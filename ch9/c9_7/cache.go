package c9_7

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}

func NewMemo(f Func) *Memo {
	return &Memo{
		f,
		make(map[string]*entry),
		sync.Mutex{},
	}
}

type Func func(key string) (any, error)

type entry struct {
	res   result
	ready chan bool
}

type result struct {
	value any
	err   error
}

func (memo *Memo) Get(key string) (any, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan bool)}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
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
