package main

import (
	"fmt"
	"os"
	"time"

	"github.com/t0nem1n/goplexercise/ch8/exer8_6"
)

func main() {
	maxDepth := 2
	worker := 20
	start := time.Now()
	exer8_6.Crawl(os.Args[1:], maxDepth, worker)
	dur := time.Since(start)
	fmt.Println(dur.String())
}
