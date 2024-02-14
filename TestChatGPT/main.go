package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
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

	scannerbanner := bufio.NewScanner(banner)

	for scannerbanner.Scan() {
		printColorText(scannerbanner.Text())
	}

	if err := scannerbanner.Err(); err != nil {
		fmt.Println("Error: ", err)
	}

	godotenv.Load(".env")
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	scanner := bufio.NewScanner(os.Stdin)

	var message []openai.ChatCompletionMessage
	fmt.Print("\x1b[35muser> \x1b[0m")
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Input error: ", err)
			fmt.Println()
			return
		}

		userMessage := openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: scanner.Text(),
		}
		messages := append(message, userMessage)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)
		if err != nil {
			fmt.Printf("Chat completion error: %v\n", err)
		}
		botRespone := resp.Choices[0].Message
		//messages = append(messages, botRespone)

		fmt.Println("\x1b[34;1mcomp> ", botRespone.Content, "\x1b[0m")
		fmt.Println()
		fmt.Print("\x1b[35muser> \x1b[0m")
	}
}
