package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bpradana/tars/llm"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/template"
)

// ChatBot represents a simple conversational AI
type ChatBot struct {
	provider   llm.BaseProvider
	template   template.Template
	history    []message.Message
	maxHistory int
}

// NewChatBot creates a new chat bot instance
func NewChatBot(provider llm.BaseProvider) *ChatBot {
	return &ChatBot{
		provider:   provider,
		maxHistory: 10, // Keep last 10 messages for context
		template: template.From(
			message.FromSystem("You are a helpful and friendly AI assistant. Keep responses concise and engaging."),
		),
	}
}

// AddMessage adds a message to the conversation history
func (cb *ChatBot) AddMessage(msg message.Message) {
	cb.history = append(cb.history, msg)

	// Keep only the last maxHistory messages
	if len(cb.history) > cb.maxHistory {
		cb.history = cb.history[len(cb.history)-cb.maxHistory:]
	}
}

// BuildTemplate creates a template with conversation history
func (cb *ChatBot) BuildTemplate(userInput string) template.Template {
	messages := []message.Message{
		cb.template.GetMessage()[0], // System message
	}

	// Add conversation history
	messages = append(messages, cb.history...)

	// Add current user input
	messages = append(messages, message.FromUser(userInput))

	return template.From(messages...)
}

// SendMessage sends a message and returns the response
func (cb *ChatBot) SendMessage(ctx context.Context, userInput string) (string, error) {
	// Build template with history
	template := cb.BuildTemplate(userInput)

	// Validate template
	if err := template.Validate(); err != nil {
		return "", fmt.Errorf("template validation failed: %v", err)
	}

	// Send to LLM
	response, err := cb.provider.Invoke(ctx, template)
	if err != nil {
		return "", fmt.Errorf("failed to get response: %v", err)
	}

	// Add messages to history
	cb.AddMessage(message.FromUser(userInput))
	cb.AddMessage(response)

	return response.GetContent(), nil
}

func main() {
	fmt.Println("=== Chat Bot Example ===")
	fmt.Println("Type 'quit' to exit")
	fmt.Println()

	// Create provider
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Create chat bot
	bot := NewChatBot(provider)

	// Start interactive chat
	scanner := bufio.NewScanner(os.Stdin)
	ctx := context.Background()

	for {
		fmt.Print("You: ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "quit" {
			fmt.Println("Goodbye!")
			break
		}

		if userInput == "" {
			continue
		}

		// Get response from bot
		response, err := bot.SendMessage(ctx, userInput)
		if err != nil {
			log.Printf("Error: %v", err)
			continue
		}

		fmt.Printf("Bot: %s\n\n", response)
	}
}
