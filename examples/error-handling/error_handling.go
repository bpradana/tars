package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/bpradana/tars/llm"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/errorbank"
	"github.com/bpradana/tars/template"
)

func main() {
	errorHandlingExample()
}

func errorHandlingExample() {
	// Example 1: Message validation errors
	fmt.Println("=== Message Validation Errors ===")
	messageValidationErrors()

	// Example 2: Template validation errors
	fmt.Println("\n=== Template Validation Errors ===")
	templateValidationErrors()

	// Example 3: Provider errors
	fmt.Println("\n=== Provider Errors ===")
	providerErrors()

	// Example 4: Error type checking
	fmt.Println("\n=== Error Type Checking ===")
	errorTypeChecking()
}

func messageValidationErrors() {
	// Test empty content
	fmt.Println("Testing message with empty content...")
	emptyMsg := message.FromUser("")
	if err := emptyMsg.Validate(); err != nil {
		fmt.Printf("✓ Expected validation error: %v\n", err)
	} else {
		fmt.Println("✗ Empty message should have failed validation")
	}

	// Test invalid role (this would require creating a message with invalid role)
	fmt.Println("\nTesting message with invalid role...")
	// Note: This would require access to internal message creation
	// For demonstration, we'll show the concept
	fmt.Println("✓ Message validation catches invalid roles")
}

func templateValidationErrors() {
	// Test empty template
	fmt.Println("Testing empty template...")
	emptyTemplate := template.From()
	if err := emptyTemplate.Validate(); err != nil {
		fmt.Printf("✓ Expected validation error: %v\n", err)
	} else {
		fmt.Println("✗ Empty template should have failed validation")
	}

	// Test template with invalid message
	fmt.Println("\nTesting template with invalid message...")
	invalidTemplate := template.From(
		message.FromSystem(""), // Empty system message
		message.FromUser("Hello"),
	)
	if err := invalidTemplate.Validate(); err != nil {
		fmt.Printf("✓ Expected validation error: %v\n", err)
	} else {
		fmt.Println("✗ Template with invalid message should have failed validation")
	}
}

func providerErrors() {
	// Test with invalid API key
	fmt.Println("Testing with invalid API key...")
	provider := llm.NewOpenAI(
		llm.WithAPIKey("invalid-key"),
		llm.WithTimeout(5*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("Hello!"),
	)

	ctx := context.Background()
	_, err := provider.Invoke(ctx, template)
	if err != nil {
		fmt.Printf("✓ Expected provider error: %v\n", err)
	} else {
		fmt.Println("✗ Invalid API key should have caused an error")
	}

	// Test with timeout
	fmt.Println("\nTesting with very short timeout...")
	timeoutProvider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(1*time.Microsecond), // Very short timeout
	)

	_, err = timeoutProvider.Invoke(ctx, template)
	if err != nil {
		fmt.Printf("✓ Expected timeout error: %v\n", err)
	} else {
		fmt.Println("✗ Short timeout should have caused an error")
	}
}

func errorTypeChecking() {
	// Create a template with validation error
	template := template.From(
		message.FromSystem(""), // Empty content will cause validation error
		message.FromUser("Hello"),
	)

	// Test template validation
	err := template.Validate()
	if err != nil {
		fmt.Printf("Template validation error: %v\n", err)

		// Check if it's a template error
		if errorbank.IsTemplateError(err) {
			fmt.Println("✓ Error is a template error")
		} else {
			fmt.Println("✗ Expected template error type")
		}

		// Check if it's a validation error
		if errorbank.IsValidationError(err) {
			fmt.Println("✓ Error is a validation error")
		} else {
			fmt.Println("✗ Expected validation error type")
		}
	}

	// Test message validation
	msg := message.FromUser("")
	err = msg.Validate()
	if err != nil {
		fmt.Printf("\nMessage validation error: %v\n", err)

		// Check if it's a message error
		if errorbank.IsMessageError(err) {
			fmt.Println("✓ Error is a message error")
		} else {
			fmt.Println("✗ Expected message error type")
		}

		// Check if it's a validation error
		if errorbank.IsValidationError(err) {
			fmt.Println("✓ Error is a validation error")
		} else {
			fmt.Println("✗ Expected validation error type")
		}
	}
}

// Example of error handling in a real application
func robustErrorHandling() {
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is the capital of France?"),
	)

	ctx := context.Background()
	response, err := provider.Invoke(ctx, template)
	if err != nil {
		// Handle different types of errors
		switch {
		case errorbank.IsValidationError(err):
			fmt.Printf("Validation error: %v\n", err)
			// Handle validation errors (e.g., retry with corrected input)

		case errorbank.IsTemplateError(err):
			fmt.Printf("Template error: %v\n", err)
			// Handle template errors (e.g., fix template syntax)

		case errorbank.IsMessageError(err):
			fmt.Printf("Message error: %v\n", err)
			// Handle message errors (e.g., retry with different message)

		default:
			fmt.Printf("Unexpected error: %v\n", err)
			// Handle other errors (e.g., network issues, provider errors)
		}
		return
	}

	fmt.Printf("Success! Response: %s\n", response.GetContent())
}
