package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var (
	model       string
	temperature float32
)

var chatCmd = &cobra.Command{
	Use:   "chat [prompt]",
	Short: "Start a chat session with the AI",
	Long:  `Starts an interactive chat session. Provide an initial prompt as an argument.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			fmt.Println("Error: OPENAI_API_KEY environment variable not set.")
			os.Exit(1)
		}

		client := openai.NewClient(apiKey)
		prompt := strings.Join(args, " ")

		req := openai.ChatCompletionRequest{
			Model:       model,
			Temperature: temperature,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		}

		resp, err := client.CreateChatCompletion(context.Background(), req)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			return
		}

		if len(resp.Choices) > 0 {
			fmt.Println(resp.Choices[0].Message.Content)
		} else {
			fmt.Println("No response from the AI.")
		}
	},
}

func init() {
	rootCmd.AddCommand(chatCmd)
	chatCmd.Flags().StringVarP(&model, "model", "m", openai.GPT3Dot5Turbo, "The model to use for the chat")
	chatCmd.Flags().Float32VarP(&temperature, "temperature", "t", 0.7, "The temperature for the AI response")
}
