package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bpradana/tars/llm"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/template"
)

func main() {
	// Example 1: Basic OpenAI usage
	fmt.Println("=== Basic OpenAI Usage ===")
	basicOpenAIExample()

	// Example 2: Using the factory pattern
	fmt.Println("\n=== Factory Pattern Example ===")
	factoryPatternExample()

	// Example 3: Template system
	fmt.Println("\n=== Template System Example ===")
	templateSystemExample()

	// Example 4: Error handling
	fmt.Println("\n=== Error Handling Example ===")
	errorHandlingExample()
}

func basicOpenAIExample() {
	// Create an OpenAI provider
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Create a simple conversation template
	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is the capital of France?"),
	)

	// Validate the template
	if err := template.Validate(); err != nil {
		log.Printf("Template validation failed: %v", err)
		return
	}

	// Send to LLM
	ctx := context.Background()
	response, err := provider.Invoke(ctx, template)
	if err != nil {
		log.Printf("Error invoking LLM: %v", err)
		return
	}

	fmt.Printf("Response: %s\n", response.GetContent())
}

func factoryPatternExample() {
	// Use the factory pattern to create a provider
	provider, err := llm.NewProvider(llm.ProviderOpenAI,
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Printf("Failed to create provider: %v", err)
		return
	}

	// Create a template with multiple messages
	template := template.From(
		message.FromSystem("You are a coding assistant. Provide concise answers."),
		message.FromUser("Write a simple 'Hello, World!' program in Go."),
	)

	ctx := context.Background()
	response, err := provider.Invoke(ctx, template)
	if err != nil {
		log.Printf("Error invoking LLM: %v", err)
		return
	}

	fmt.Printf("Response: %s\n", response.GetContent())
}

func templateSystemExample() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Create a template with variables
	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("Hello, {{name}}! How is the weather in {{city}} today?"),
	)

	// Substitute variables
	invokedTemplate := template.Invoke(map[string]any{
		"name": "Alice",
		"city": "Paris",
	})

	ctx := context.Background()
	response, err := provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		log.Printf("Error invoking LLM: %v", err)
		return
	}

	fmt.Printf("Response: %s\n", response.GetContent())
}

func errorHandlingExample() {
	// Example with invalid API key to demonstrate error handling
	provider := llm.NewOpenAI(
		llm.WithAPIKey("invalid-key"),
		llm.WithTimeout(5*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is 2+2?"),
	)

	ctx := context.Background()
	_, err := provider.Invoke(ctx, template)
	if err != nil {
		fmt.Printf("Expected error occurred: %v\n", err)
	} else {
		fmt.Println("Unexpected success with invalid API key")
	}
}
