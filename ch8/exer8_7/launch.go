package exer8_7

import (
	"fmt"
	"os"
	"time"
)

func Launch(countDown int) {
	abort := make(chan struct{})
	go func() {
		_, _ = os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop()
	for i := countDown; i > 0; i-- {
		select {
		case <-tick.C:
			fmt.Printf("\033[2K\r%d\r", i)
		case <-abort:
			fmt.Printf("abort launch\n")
			return
		}
	}
	fmt.Println("\rlaunching...")
}
