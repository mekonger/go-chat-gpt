package main

import (
	"context"
	"fmt"
	"github.com/PullRequestInc/go-gpt3"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	godotenv.Load()
	apiKey := os.Getenv("API_KEY")
	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	request := gpt3.CompletionRequest{
		Prompt: []string{"How many coffees should I drink per day?"},
	}
	resp, err := client.CompletionWithEngine(ctx, "gpt-3.5-turbo-instruct", request)
	if err != nil {
		fmt.Printf("%s\n", err)
	} else {
		fmt.Printf("Answer: %s\n", resp.Choices[0].Text)
	}
}
