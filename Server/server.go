package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns   []net.Conn
	connCh  = make(chan net.Conn)
	closeCh = make(chan net.Conn)
	msgCh   = make(chan string)
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			conn, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}

			conns = append(conns, conn)
			connCh <- conn

		}
	}()

	for {
		select {
		case conn := <-connCh:
			go onMessage(conn)
		case msg := <-msgCh:
			fmt.Print(msg)
		case conn := <-closeCh:
			fmt.Println("Client exit", conn)
			// removeConn(conn)
		}
	}

}

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			break
		}

		msgCh <- msg

	}
	closeCh <- conn
}
