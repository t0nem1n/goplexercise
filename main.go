package main

import (
	"fmt"
	"time"

	"github.com/t0nem1n/goplexercise/ch8/exer8_7"
)

func main() {
	start := time.Now()
	exer8_7.Launch(10)
	dur := time.Since(start)
	fmt.Println(dur.String())
}
