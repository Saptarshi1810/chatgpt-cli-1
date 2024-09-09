package main

import(
	"os"
	"github.com/joho/godotenv"
	"log"
	openai "github.com/sashabaranov/go-openai"
	"fmt"
	"context"
)


func main(){
	godotenv.Load()
	
	apiKey := os.Getenv("API_KEY")
	
	if apiKey == ""{
		log.Fatalln("Missing API Key")
	}

	ctx := context.Background()
	client :=openai.NewClient(apiKey)

	resp,err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Hello!",
			},
		},
	},)


	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.Choices[0].Message.Content)
}