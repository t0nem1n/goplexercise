package c9_7

import "sync"

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
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
