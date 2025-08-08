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
	advancedTemplateExample()
}

func advancedTemplateExample() {
	// Example 1: Complex template with multiple variables
	fmt.Println("=== Complex Template Example ===")
	complexTemplateExample()

	// Example 2: Template validation
	fmt.Println("\n=== Template Validation Example ===")
	templateValidationExample()

	// Example 3: Template composition
	fmt.Println("\n=== Template Composition Example ===")
	templateCompositionExample()

	// Example 4: JSON serialization
	fmt.Println("\n=== JSON Serialization Example ===")
	jsonSerializationExample()
}

func complexTemplateExample() {
	// Create a complex template with multiple variables
	template := template.From(
		message.FromSystem("You are a travel assistant. Provide helpful information about destinations."),
		message.FromUser("I'm planning a trip to {{.Destination}} for {{.Duration}} days. I'm interested in {{.Interests}}. Can you suggest an itinerary?"),
		message.FromAssistant("I'd be happy to help you plan your trip to {{.Destination}}! Let me create a {{.Duration}}-day itinerary focused on {{.Interests}}."),
		message.FromUser("What's the best time to visit {{.Destination}}?"),
	)

	// Substitute multiple variables
	invokedTemplate := template.Invoke(struct {
		Destination string
		Duration    string
		Interests   string
	}{
		Destination: "Japan",
		Duration:    "7",
		Interests:   "culture, food, and technology",
	})

	// Validate the template
	if err := invokedTemplate.Validate(); err != nil {
		fmt.Printf("Template validation failed: %v\n", err)
		return
	}

	// Send to LLM
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	ctx := context.Background()
	response, err := provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		fmt.Printf("Error invoking LLM: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", response.GetContent())
}

func templateValidationExample() {
	// Example 1: Valid template
	fmt.Println("Testing valid template...")
	validTemplate := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("Hello!"),
	)

	if err := validTemplate.Validate(); err != nil {
		fmt.Printf("Unexpected validation error: %v\n", err)
	} else {
		fmt.Println("✓ Valid template passed validation")
	}

	// Example 2: Empty template
	fmt.Println("\nTesting empty template...")
	emptyTemplate := template.From()

	if err := emptyTemplate.Validate(); err != nil {
		fmt.Printf("✓ Expected validation error: %v\n", err)
	} else {
		fmt.Println("✗ Empty template should have failed validation")
	}

	// Example 3: Template with invalid message
	fmt.Println("\nTesting template with invalid message...")
	invalidTemplate := template.From(
		message.FromSystem(""), // Empty content
		message.FromUser("Hello!"),
	)

	if err := invalidTemplate.Validate(); err != nil {
		fmt.Printf("✓ Expected validation error: %v\n", err)
	} else {
		fmt.Println("✗ Template with invalid message should have failed validation")
	}
}

func templateCompositionExample() {
	// Create base templates
	baseSystemPrompt := template.From(
		message.FromSystem("You are a helpful assistant."),
	)

	greetingTemplate := template.From(
		message.FromUser("Hello, {{.Name}}! How are you today?"),
	)

	questionTemplate := template.From(
		message.FromUser("What's the weather like in {{.City}}?"),
	)

	// Compose templates by combining them
	composedTemplate := template.From(
		baseSystemPrompt.GetMessage()[0], // Get the system message
		greetingTemplate.GetMessage()[0], // Get the greeting message
		questionTemplate.GetMessage()[0], // Get the question message
	)

	// Substitute variables
	invokedTemplate := composedTemplate.Invoke(struct {
		Name string
		City string
	}{
		Name: "Alice",
		City: "Paris",
	})

	// Validate and use
	if err := invokedTemplate.Validate(); err != nil {
		fmt.Printf("Template validation failed: %v\n", err)
		return
	}

	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	ctx := context.Background()
	response, err := provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		fmt.Printf("Error invoking LLM: %v\n", err)
		return
	}

	fmt.Printf("Response: %s\n", response.GetContent())
}

func jsonSerializationExample() {
	// Create a template
	template := template.From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("Hello, {{.Name}}!"),
	)

	// Substitute variables
	invokedTemplate := template.Invoke(struct {
		Name string
	}{
		Name: "Bob",
	})

	// Serialize to JSON
	jsonStr := invokedTemplate.ToJSON()
	fmt.Printf("Template JSON: %s\n", jsonStr)

	// Also serialize individual messages
	messages := invokedTemplate.GetMessage()
	for i, msg := range messages {
		msgJSON := msg.ToJSON()
		fmt.Printf("Message %d JSON: %s\n", i+1, msgJSON)
	}
}
