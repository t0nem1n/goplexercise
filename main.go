package main

import (
	"fmt"
	"time"

	"github.com/t0nem1n/goplexercise/ch9/c9_7"
)

func main() {
	start := time.Now()
	c9_7.Run()
	fmt.Printf("running in %v\n", time.Since(start).String())
}
