# TARS - Unified LLM Interface

<p align="center">
  <img src="./assets/logo.jpg" alt="TARS logo" />
</p>

TARS is a Go library that provides a unified interface for interacting with various Large Language Model (LLM) providers. It abstracts away the differences between providers like OpenAI, Anthropic, OpenRouter, and Ollama, allowing you to switch between them seamlessly.

## Features

- **Unified Interface**: Single API for multiple LLM providers
- **Template System**: Reusable conversation templates with variable substitution
- **Structured Output**: JSON schema support for consistent, typed responses
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
        message.FromUser("What is the capital of {{.Country}}?"),
    )

    // Validate the template before using it
    if err := template.Validate(); err != nil {
        log.Fatalf("Template validation failed: %v", err)
    }

    // Substitute variables and send to LLM
    prompt := template.Invoke(struct {
        Country string
    } {
        Country: "France",
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
    message.FromUser("Hello, {{.Name}}! How are you today?"),
)

// Validate a template
if err := template.Validate(); err != nil {
    log.Fatal(err)
}

// Substitute variables
prompt := template.Invoke(struct {
    Name string
} {
    Name: "Alice",
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

### Structured Output

TARS supports structured output using JSON schemas, allowing you to get consistent, typed responses from LLM providers. This is useful for applications that need to process LLM responses programmatically.

#### Basic Structured Output

```go
// Define a struct for structured output
type WeatherInfo struct {
    Temperature float64 `json:"temperature"`
    Condition   string  `json:"condition"`
    Humidity    int     `json:"humidity"`
    WindSpeed   float64 `json:"wind_speed"`
    Description string  `json:"description"`
}

// Create a template
template := template.From(
    message.FromSystem("You are a weather assistant. Provide weather information in the requested format."),
    message.FromUser("What's the weather like in Paris today?"),
)

// Use structured output
var weatherInfo WeatherInfo
response, err := provider.Invoke(context.Background(), template,
    llm.WithStructuredOutput(&weatherInfo),
    llm.WithTemperature(0.3), // Lower temperature for more consistent structured output
)
if err != nil {
    log.Fatal(err)
}

// The response content will be JSON that matches the WeatherInfo struct
fmt.Printf("Temperature: %.1f°C\n", weatherInfo.Temperature)
fmt.Printf("Condition: %s\n", weatherInfo.Condition)
```

#### Complex Structured Output

```go
// Define a complex struct with arrays and nested objects
type AnalysisResult struct {
    Sentiment      string   `json:"sentiment"`
    Confidence     float64  `json:"confidence"`
    KeyPoints      []string `json:"key_points"`
    Summary        string   `json:"summary"`
    Recommendations []string `json:"recommendations"`
}

template := template.From(
    message.FromSystem("You are a text analysis assistant. Analyze the given text and provide structured insights."),
    message.FromUser("Analyze this text: 'The new smartphone features an amazing camera and great battery life.'"),
)

var analysis AnalysisResult
response, err := provider.Invoke(context.Background(), template,
    llm.WithStructuredOutput(&analysis),
    llm.WithTemperature(0.2),
)

fmt.Printf("Sentiment: %s (Confidence: %.2f)\n", analysis.Sentiment, analysis.Confidence)
fmt.Printf("Key Points: %v\n", analysis.KeyPoints)
```

#### Best Practices for Structured Output

1. **Use Lower Temperature**: Set temperature to 0.2-0.3 for more consistent structured output
2. **Clear Instructions**: Provide explicit instructions about the expected format
3. **Validation**: Always validate the structured output before using it
4. **Error Handling**: Handle cases where the LLM doesn't return valid JSON
5. **Schema Design**: Design your structs to be clear and unambiguous

```go
// Example with error handling
var result MyStruct
response, err := provider.Invoke(ctx, template,
    llm.WithStructuredOutput(&result),
    llm.WithTemperature(0.2),
)
if err != nil {
    log.Printf("Error getting structured output: %v", err)
    return
}

// Validate the structured output
if result.RequiredField == "" {
    log.Printf("Invalid structured output: missing required field")
    return
}
```

## Error Handling

The library provides comprehensive error handling with custom error types for better debugging and error management.


### Error Handling Examples

```go
response, err := provider.Invoke(context.Background(), template)
if err != nil {
    // Check for specific error types
    if errorbank.IsMessageError(err) {
        log.Printf("Message error: %v", err)
        // Handle message-specific errors
    } else if errorbank.IsTemplateError(err) {
        log.Printf("Template error: %v", err)
        // Handle template errors
    } else if errorbank.IsValidationError(err) {
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
if err := msg.Validate(); err != nil {
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
err := errorbank.NewMessageError("invoke", "failed to process template", originalErr)
// Output: [invoke] failed to process template: <original error>

// TemplateError provides variable and context information
err := errorbank.NewTemplateError("name", "variable not found", nil)
// Output: [Template] variable 'name': variable not found

// ValidationError provides field and value information
err := errorbank.NewValidationError("content", "cannot be empty", "")
// Output: [Validation] field 'content': cannot be empty (value: )
```

## Architecture

### Core Components

1. **LLM Providers** (`llm/`): Implementations for different LLM services
2. **Message System** (`message/`): Conversation message handling and templates
3. **Template System** (`template/`): Variable substitution and conversation templates
4. **HTTP Client** (`pkg/httpx/`): HTTP request handling and utilities
5. **Error Handling** (`pkg/errorbank/`): Custom error types and utilities

### Design Patterns

- **Interface Segregation**: Clean interfaces for different concerns
- **Functional Options**: Flexible configuration without complex constructors
- **Factory Pattern**: Centralized provider creation
- **Template Method**: Reusable conversation patterns
- **Error Wrapping**: Comprehensive error context and type checking


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
- [ ] Add caching capabilities
- [ ] Add monitoring and metrics
- [ ] Add configuration management
- [ ] Add plugin system 
