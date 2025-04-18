package main

import (
    "bufio"
    "context"
    "fmt"
    "os"

    "github.com/anthropics/anthropic-sdk-go"
)

type Agent struct {
    client *anthropic.Client
    getUserMessage func() (string, bool)
}

func main() {
    client := anthropic.NewClient()

    scanner := bufio.NewScanner(os.Stdin)
    getUserMessage := func() (string, bool) {
	if !scanner.Scan() {
	    return "", false
	}
	return scanner.Text(), true
    }

    agent := NewAgent(&client, getUserMessage)
    err := agent.Run(context.TODO())
    if err != nil {
	fmt.Printf("Error: %s\n", err.Error())
    }
}

func (a *Agent) Run(ctx context.Context) error {
	conversation := []anthropic.MessageParam{}

	fmt.Println("Chat with Claude (use 'ctrl-c' to quit)")

	for {
		fmt.Print("\u001b[94mYou\u001b[0m: ")
		userInput, ok := a.getUserMessage()
		if !ok {
			break
		}

		userMessage := anthropic.NewUserMessage(anthropic.NewTextBlock(userInput))
		conversation = append(conversation, userMessage)

		message, err := a.runInference(ctx, conversation)
		if err != nil {
			return err
		}
		conversation = append(conversation, message.ToParam())

		for _, content := range message.Content {
			switch content.Type {
			case "text":
				fmt.Printf("\u001b[93mClaude\u001b[0m: %s\n", content.Text)
			}
		}
	}

	return nil
}

func NewAgent(client *anthropic.Client, getUserMessage func() (string, bool)) {
    return &Agent{
	client: client,
	getUserMessage: getUserMessage,
    }
}
