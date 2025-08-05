# Tars - Unified LLM Interface

Tars is a Go library that provides a unified interface for interacting with various Large Language Model (LLM) providers. It abstracts away the differences between providers like OpenAI, Anthropic, OpenRouter, and Ollama, allowing you to switch between them seamlessly.

## Features

- **Unified Interface**: Single API for multiple LLM providers
- **Template System**: Reusable conversation templates with variable substitution
- **Type Safety**: Strong typing throughout the codebase
- **Error Handling**: Comprehensive error types and handling
- **Validation**: Built-in validation for messages and templates
- **Extensible**: Easy to add new LLM providers
- **Production Ready**: Comprehensive documentation and testing

## Supported Providers

- **OpenAI**: GPT-4, GPT-3.5-turbo, and other OpenAI models
- **Anthropic**: Claude-3, Claude-2, and other Anthropic models
- **OpenRouter**: Access to multiple providers through a unified API
- **Ollama**: Local models like Llama, Mistral, and others

## Installation

```bash
go get github.com/bpradana/tars
```

## Quick Start

```go
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
    // Create an OpenAI provider
    provider := llm.NewOpenAI(
        llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
        llm.WithTimeout(30*time.Second),
    )

    // Create a conversation template
    template := template.From(
        message.FromSystem("You are a helpful assistant."),
        message.FromUser("What is the capital of {{country}}?"),
    )

    // Validate the template before using it
    if err := template.Validate(); err != nil {
        log.Fatalf("Template validation failed: %v", err)
    }

    // Substitute variables and send to LLM
    prompt := template.Invoke(map[string]any{
        "country": "France",
    })

    response, err := provider.Invoke(context.Background(), prompt)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Response: %s\n", response.GetContent())
}
```

## Usage

### Creating LLM Providers

```go
// OpenAI
provider := llm.NewOpenAI(
    llm.WithAPIKey("your-api-key"),
)

// Anthropic
provider := llm.NewAnthropic(
    llm.WithAPIKey("your-api-key"),
)

// OpenRouter
provider := llm.NewOpenRouter(
    llm.WithAPIKey("your-api-key"),
)

// Ollama
provider := llm.NewOllama(
    llm.WithBaseURL("http://localhost:11434"),
)
```

### Using the Factory Pattern

```go
provider, err := llm.NewProvider(llm.ProviderOpenAI,
    llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
    llm.WithTimeout(30*time.Second),
)
if err != nil {
    log.Fatal(err)
}
```

### Creating Messages

```go
// System message (sets behavior)
systemMsg := message.FromSystem("You are a helpful assistant.")
if err := systemMsg.Validate(); err != nil {
    log.Fatal(err)
}

// User message
userMsg := message.FromUser("What is the weather like?")
if err := userMsg.Validate(); err != nil {
    log.Fatal(err)
}

// Assistant message (typically from LLM response)
assistantMsg := message.FromAssistant("The weather is sunny today.",
    message.WithUsage(10, 5, 15), // prompt, completion, total tokens
)
if err := assistantMsg.Validate(); err != nil {
    log.Fatal(err)
}
```

### Working with Templates

```go
// Create a template
template := template.From(
    message.FromSystem("You are a helpful assistant."),
    message.FromUser("Hello, {{name}}! How are you today?"),
)

// Validate a template
err := template.Validate()
if err != nil {
    log.Fatal(err)
}

// Substitute variables
prompt := template.Invoke(map[string]any{
    "name": "Alice",
})

// Send to LLM
response, err := provider.Invoke(context.Background(), prompt)
```

### Customizing Requests

```go
response, err := provider.Invoke(context.Background(), template,
    llm.WithModel("gpt-4"),
    llm.WithTemperature(0.7),
    llm.WithMaxTokens(1000),
)
```

## Error Handling

The library provides comprehensive error handling with custom error types for better debugging and error management.


### Error Handling Examples

```go
response, err := provider.Invoke(context.Background(), template)
if err != nil {
    // Check for specific error types
    if message.IsMessageError(err) {
        log.Printf("Message error: %v", err)
        // Handle message-specific errors
    } else if message.IsTemplateError(err) {
        log.Printf("Template error: %v", err)
        // Handle template errors
    } else if message.IsValidationError(err) {
        log.Printf("Validation error: %v", err)
        // Handle validation errors
    } else {
        log.Printf("Unexpected error: %v", err)
        // Handle other errors
    }
    return
}
```

### Validation

```go
// Validate messages
msg := message.FromUser("")
if err := msg.(interface{ Validate() error }).Validate(); err != nil {
    log.Printf("Message validation failed: %v", err)
}

// Validate templates
template := template.From()
if err := template.Validate(); err != nil {
    log.Printf("Template validation failed: %v", err)
}
```

### Error Context

All custom errors provide rich context information:

```go
// MessageError provides operation and message context
err := message.NewMessageError("invoke", "failed to process template", originalErr)
// Output: [invoke] failed to process template: <original error>

// TemplateError provides variable and context information
err := message.NewTemplateError("name", "variable not found", nil)
// Output: [Template] variable 'name': variable not found

// ValidationError provides field and value information
err := message.NewValidationError("content", "cannot be empty", "")
// Output: [Validation] field 'content': cannot be empty (value: )
```

## Architecture

### Core Components

1. **LLM Providers** (`llm/`): Implementations for different LLM services
2. **Message System** (`message/`): Conversation message handling and templates
3. **Template System** (`template/`): Variable substitution and conversation templates
4. **HTTP Client** (`pkg/httpx/`): HTTP request handling and utilities
5. **Error Handling** (`message/errors.go`): Custom error types and utilities

### Design Patterns

- **Interface Segregation**: Clean interfaces for different concerns
- **Functional Options**: Flexible configuration without complex constructors
- **Factory Pattern**: Centralized provider creation
- **Template Method**: Reusable conversation patterns
- **Error Wrapping**: Comprehensive error context and type checking

## Testing

Run the test suite:

```bash
go test ./...
```

Run tests with coverage:

```bash
go test -cover ./...
```

Run specific error handling tests:

```bash
go test ./message -v
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For questions, issues, or contributions, please open an issue on GitHub.

## Roadmap

- [ ] Add more LLM providers
- [ ] Implement streaming responses
- [ ] Add rate limiting
- [ ] Add retry mechanisms
- [ ] Add caching capabilities
- [ ] Add monitoring and metrics
- [ ] Add CLI tool
- [ ] Add configuration management
- [ ] Add plugin system 