package main

import (
	"fmt"
	"time"

	"github.com/t0nem1n/goplexercise/ch8/exer8_7"
)

func main() {
	start := time.Now()
	timeout := 10 * time.Second
	exer8_7.Shout(timeout)
	dur := time.Since(start)
	fmt.Println(dur.String())
}
