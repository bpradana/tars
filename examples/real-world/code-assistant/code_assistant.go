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

// CodeAssistant provides code-related AI assistance
type CodeAssistant struct {
	provider llm.BaseProvider
}

// NewCodeAssistant creates a new code assistant
func NewCodeAssistant(provider llm.BaseProvider) *CodeAssistant {
	return &CodeAssistant{
		provider: provider,
	}
}

// GenerateCode generates code based on requirements
func (ca *CodeAssistant) GenerateCode(ctx context.Context, language, description, requirements string) (string, error) {
	template := template.From(
		message.FromSystem("You are an expert programmer. Write clean, well-documented, and efficient code."),
		message.FromUser("Write {{.Language}} code for: {{.Description}}. Requirements: {{.Requirements}}. Include comments and follow best practices."),
	)

	invokedTemplate := template.Invoke(struct {
		Language     string
		Description  string
		Requirements string
	}{
		Language:     language,
		Description:  description,
		Requirements: requirements,
	})

	response, err := ca.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to generate code: %v", err)
	}

	return response.GetContent(), nil
}

// DebugCode analyzes and fixes code issues
func (ca *CodeAssistant) DebugCode(ctx context.Context, language, code, errorMessage string) (string, error) {
	template := template.From(
		message.FromSystem("You are a debugging expert. Analyze code and provide fixes for issues."),
		message.FromUser("Debug this {{.Language}} code:\n\n{{.Code}}\n\nError: {{.ErrorMessage}}\n\nProvide the corrected code and explanation."),
	)

	invokedTemplate := template.Invoke(struct {
		Language     string
		Code         string
		ErrorMessage string
	}{
		Language:     language,
		Code:         code,
		ErrorMessage: errorMessage,
	})

	response, err := ca.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to debug code: %v", err)
	}

	return response.GetContent(), nil
}

// ReviewCode performs code review and suggests improvements
func (ca *CodeAssistant) ReviewCode(ctx context.Context, language, code string) (string, error) {
	template := template.From(
		message.FromSystem("You are a senior code reviewer. Analyze code for best practices, performance, and maintainability."),
		message.FromUser("Review this {{.Language}} code:\n\n{{.Code}}\n\nProvide feedback on code quality, potential issues, and improvement suggestions."),
	)

	invokedTemplate := template.Invoke(struct {
		Language string
		Code     string
	}{
		Language: language,
		Code:     code,
	})

	response, err := ca.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to review code: %v", err)
	}

	return response.GetContent(), nil
}

// ExplainCode explains how code works
func (ca *CodeAssistant) ExplainCode(ctx context.Context, language, code string) (string, error) {
	template := template.From(
		message.FromSystem("You are a programming instructor. Explain code in a clear and educational manner."),
		message.FromUser("Explain this {{.Language}} code step by step:\n\n{{.Code}}\n\nProvide a detailed explanation suitable for learning."),
	)

	invokedTemplate := template.Invoke(struct {
		Language string
		Code     string
	}{
		Language: language,
		Code:     code,
	})

	response, err := ca.provider.Invoke(ctx, invokedTemplate)
	if err != nil {
		return "", fmt.Errorf("failed to explain code: %v", err)
	}

	return response.GetContent(), nil
}

func main() {
	fmt.Println("=== Code Assistant Example ===")

	// Create provider
	provider := llm.NewOpenAI(
		llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
		llm.WithTimeout(30*time.Second),
	)

	// Create code assistant
	assistant := NewCodeAssistant(provider)
	ctx := context.Background()

	// Example 1: Generate code
	fmt.Println("\n--- Code Generation ---")
	code, err := assistant.GenerateCode(ctx,
		"Python",
		"a function to calculate fibonacci numbers",
		"should handle edge cases, be efficient, and include docstring")
	if err != nil {
		log.Printf("Error generating code: %v", err)
	} else {
		fmt.Printf("Generated Code:\n%s\n", code)
	}

	// Example 2: Debug code
	fmt.Println("\n--- Code Debugging ---")
	buggyCode := `def calculate_average(numbers):
    total = 0
    for num in numbers:
        total += num
    return total / len(numbers)  # This will fail if numbers is empty`

	debuggedCode, err := assistant.DebugCode(ctx,
		"Python",
		buggyCode,
		"ZeroDivisionError: division by zero")
	if err != nil {
		log.Printf("Error debugging code: %v", err)
	} else {
		fmt.Printf("Debugged Code:\n%s\n", debuggedCode)
	}

	// Example 3: Code review
	fmt.Println("\n--- Code Review ---")
	reviewCode := `def process_data(data):
    result = []
    for item in data:
        if item > 0:
            result.append(item * 2)
    return result`

	review, err := assistant.ReviewCode(ctx, "Python", reviewCode)
	if err != nil {
		log.Printf("Error reviewing code: %v", err)
	} else {
		fmt.Printf("Code Review:\n%s\n", review)
	}

	// Example 4: Code explanation
	fmt.Println("\n--- Code Explanation ---")
	complexCode := `def quicksort(arr):
    if len(arr) <= 1:
        return arr
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    return quicksort(left) + middle + quicksort(right)`

	explanation, err := assistant.ExplainCode(ctx, "Python", complexCode)
	if err != nil {
		log.Printf("Error explaining code: %v", err)
	} else {
		fmt.Printf("Code Explanation:\n%s\n", explanation)
	}
}
