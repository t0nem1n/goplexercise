package exer8_10

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	// handle entering, sending message, leaving event

	client := client{
		name: conn.RemoteAddr().String(),
		msg:  make(chan string),
	}

	defer func() {
		leavingMsg := fmt.Sprintf("%s leave chat", client.name)
		leaving <- client
		sending <- newMessage(client.name, leavingMsg)
	}()
	enterMsg := fmt.Sprintf("%s join chat", client.name)
	entering <- client
	sending <- newMessage(client.name, enterMsg)
	go func() {
		for msg := range client.msg {
			fmt.Fprintf(conn, "%s\n", msg)
		}
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		sending <- newMessage(client.name, input.Text())
	}
}

type message struct {
	who string
	msg string
}

func newMessage(who, msg string) message {
	return message{who, msg}
}

type client struct {
	name string
	msg  chan string
}

var (
	connection = make(map[string]client)
	entering   = make(chan client)
	sending    = make(chan message)
	leaving    = make(chan client)
)

func broadcaster() {
	for {
		select {
		case client := <-entering:
			connection[client.name] = client
		case message := <-sending:
			for _, client := range connection {
				if message.who != client.name {
					client.msg <- fmt.Sprintf("%s: %s", message.who, message.msg)
				}
			}
		case client := <-leaving:
			delete(connection, client.name)
		}
	}
}

func ChatServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	fmt.Println("start server")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}
