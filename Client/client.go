package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func printColorText(text string) {
	fmt.Println("\x1b[36;1m" + text + "\x1b[0m")
}

func main() {

	banner, err := os.Open("banner.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer banner.Close()

	scanner := bufio.NewScanner(banner)

	for scanner.Scan() {
		printColorText(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	connection, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("\x1b[33;4;1mNhập tên của bạn:\x1b[0m ")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = strings.TrimSpace(nameInput)

	go onMessage(connection)
	fmt.Println("--------------Bắt đầu cuộc trò chuyện--------------")
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}
		currentTime := time.Now()

		formattedTime := currentTime.Format("Monday, January 2, 2006 15:04:05")

		msg = fmt.Sprintf("\x1b[32;1m%s\x1b[0m: %s\n - %s", nameInput, strings.TrimSpace(msg), formattedTime)
		fmt.Println("\x1b[37;2m\x1b[3m\x1b[2m" + formattedTime + "\x1b[0m")
		connection.Write([]byte(msg))
	}
	connection.Close()
}

func onMessage(conn net.Conn) {
	for {
		fmt.Print("Chat: ")
		reader := bufio.NewReader(conn)

		msg, _ := reader.ReadString('\n')
		fmt.Print(msg)
	}
}
