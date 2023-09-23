package main

import (
	"flag"
	"fmt"

	"github.com/t0nem1n/goplexercise/ch8/exer8_8"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fmt.Println(roots)
	exer8_8.DirSize(roots)
}
