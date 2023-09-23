package exer8_7

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func handler(conn net.Conn, timeout time.Duration) {
	defer conn.Close()
	textCh := make(chan string)
	delay := 1 * time.Second
	go func() {
		input := bufio.NewScanner(conn)
		for input.Scan() {
			textCh <- input.Text()
		}
	}()

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for {
		select {
		case text := <-textCh:
			timer.Reset(timeout)
			echo(conn, text, delay)
		case <-timer.C:
			log.Println("client timeout")
			return
		}
	}
}

func echo(w io.Writer, s string, delay time.Duration) {
	fmt.Fprintf(w, "%s\n", strings.ToUpper(s))
	time.Sleep(delay)
	fmt.Fprintf(w, "%s\n", s)
	time.Sleep(delay)
	fmt.Fprintf(w, "%s\n", strings.ToLower(s))
}

// Add a timeout to the echo server from Section 8.3
// so that it disconnects any client that shouts nothing within 10 seconds.
func Shout(timeout time.Duration) {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handler(conn, timeout)
	}
}
