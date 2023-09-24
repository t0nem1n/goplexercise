package main

import (
	"fmt"
	"time"

	"github.com/t0nem1n/goplexercise/ch8/exer8_10"
)

func main() {
	start := time.Now()
	exer8_10.ChatServer()
	fmt.Printf("running in %v\n", time.Since(start).String())
}
