package main

import (
	"fmt"
	"time"

	"github.com/t0nem1n/goplexercise/ch9/exer9_1"
)

func main() {
	start := time.Now()
	exer9_1.Bank()
	fmt.Printf("running in %v\n", time.Since(start).String())
}
