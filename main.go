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

func NewAgent(client *anthropic.Client, getUserMessage func() (string, bool)) {
    return &Agent{
	client: client,
	getUserMessage: getUserMessage,
    }
}
