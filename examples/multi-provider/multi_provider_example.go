package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bpradana/tars/llm"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/template"
)

func main() {
	multiProviderExample()
}

func multiProviderExample() {
	// Example 1: Compare responses from different providers
	fmt.Println("=== Multi-Provider Comparison ===")
	compareProviders()

	// Example 2: Provider-specific configurations
	fmt.Println("\n=== Provider-Specific Configurations ===")
	providerSpecificConfigs()

	// Example 3: Fallback strategy
	fmt.Println("\n=== Fallback Strategy ===")
	fallbackStrategy()
}

func compareProviders() {
	// Create templates for comparison
	template := template.From(
		message.FromSystem("You are a helpful assistant. Provide a brief, one-sentence answer."),
		message.FromUser("What is the meaning of life?"),
	)

	// Test with different providers
	providers := []struct {
		name     string
		provider llm.BaseProvider
	}{
		{
			name: "OpenAI",
			provider: llm.NewOpenAI(
				llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
				llm.WithTimeout(30*time.Second),
			),
		},
		{
			name: "Anthropic",
			provider: llm.NewAnthropic(
				llm.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
				llm.WithTimeout(30*time.Second),
			),
		},
		{
			name: "OpenRouter",
			provider: llm.NewOpenRouter(
				llm.WithAPIKey(os.Getenv("OPENROUTER_API_KEY")),
				llm.WithTimeout(30*time.Second),
			),
		},
		{
			name: "Ollama",
			provider: llm.NewOllama(
				llm.WithBaseURL("http://localhost:11434"),
				llm.WithTimeout(30*time.Second),
			),
		},
	}

	ctx := context.Background()
	for _, p := range providers {
		fmt.Printf("\n--- Testing %s ---\n", p.name)

		response, err := p.provider.Invoke(ctx, template)
		if err != nil {
			fmt.Printf("Error with %s: %v\n", p.name, err)
			continue
		}

		fmt.Printf("Response: %s\n", response.GetContent())
	}
}

func providerSpecificConfigs() {
	// Create a template for code generation
	template := template.From(
		message.FromSystem("You are a programming expert. Write clean, well-documented code."),
		message.FromUser("Write a function to calculate the factorial of a number in Python."),
	)

	// OpenAI with specific model and parameters
	openAIProvider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Anthropic with Claude-specific settings
	anthropicProvider := llm.NewAnthropic(
		llm.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	ctx := context.Background()

	// Test OpenAI
	fmt.Println("--- OpenAI Response ---")
	response, err := openAIProvider.Invoke(ctx, template,
		llm.WithModel("gpt-4"),
		llm.WithTemperature(0.3),
		llm.WithMaxTokens(500),
	)
	if err != nil {
		fmt.Printf("OpenAI error: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response.GetContent())
	}

	// Test Anthropic
	fmt.Println("\n--- Anthropic Response ---")
	response, err = anthropicProvider.Invoke(ctx, template,
		llm.WithModel("claude-3-sonnet-20240229"),
		llm.WithTemperature(0.3),
		llm.WithMaxTokens(500),
	)
	if err != nil {
		fmt.Printf("Anthropic error: %v\n", err)
	} else {
		fmt.Printf("Response: %s\n", response.GetContent())
	}
}

func fallbackStrategy() {
	// Create a template for a simple question
	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is the capital of Japan?"),
	)

	// Define providers in order of preference
	providers := []llm.BaseProvider{
		llm.NewOpenAI(
			llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
			llm.WithTimeout(10*time.Second),
		),
		llm.NewAnthropic(
			llm.WithAPIKey(os.Getenv("ANTHROPIC_API_KEY")),
			llm.WithTimeout(10*time.Second),
		),
		llm.NewOpenRouter(
			llm.WithAPIKey(os.Getenv("OPENROUTER_API_KEY")),
			llm.WithTimeout(10*time.Second),
		),
		llm.NewOllama(
			llm.WithBaseURL("http://localhost:11434"),
			llm.WithTimeout(10*time.Second),
		),
	}

	ctx := context.Background()
	var lastErr error

	// Try each provider until one succeeds
	for i, provider := range providers {
		fmt.Printf("Trying provider %d (%s)...\n", i+1, provider.GetName())

		response, err := provider.Invoke(ctx, template)
		if err != nil {
			lastErr = err
			fmt.Printf("Provider %s failed: %v\n", provider.GetName(), err)
			continue
		}

		fmt.Printf("Success with %s!\n", provider.GetName())
		fmt.Printf("Response: %s\n", response.GetContent())
		return
	}

	fmt.Printf("All providers failed. Last error: %v\n", lastErr)
}
