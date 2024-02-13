package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Enter your name: ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	fmt.Println("Message: ")
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		msg = fmt.Sprintf("%s: %s\n", nameInput, strings.TrimSpace(msg))
		connection.Write([]byte(msg))
	}
	connection.Close()

}
