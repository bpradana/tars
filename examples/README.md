# Tars Examples

This directory contains comprehensive examples demonstrating how to use the Tars library for unified LLM interactions.

## Directory Structure

```
examples/
├── README.md                       # This file
├── basic/                          # Basic usage examples
│   └── basic_usage.go              # Comprehensive basic examples
├── multi-provider/                 # Multi-provider examples
│   └── multi_provider_example.go
├── templates/                      # Template system examples
│   └── advanced_templates.go
├── error-handling/                 # Error handling examples
│   └── error_handling.go
└── real-world/                     # Real-world application examples
    ├── chat-bot
    │   └── chat_bot.go             # Interactive chat bot
    ├── content-generator
    │   └── content_generator.go    # Content generation system
    ├── code-assistant
    │   └── code_assistant.go       # AI-powered code assistant
    └── README.md               # Real-world examples documentation
```

## Quick Start

### Prerequisites

1. Set up your API keys as environment variables:
   ```bash
   export OPENAI_API_KEY="your-openai-api-key"
   export ANTHROPIC_API_KEY="your-anthropic-api-key"
   export OPENROUTER_API_KEY="your-openrouter-api-key"
   ```

2. For Ollama examples, ensure Ollama is running locally:
   ```bash
   ollama serve
   ```

### Running Examples

#### Basic Examples
```bash
# Comprehensive basic examples
go run basic/basic_usage.go
```

#### Multi-Provider Examples
```bash
# Compare different providers
go run multi-provider/multi_provider_example.go
```

#### Template Examples
```bash
# Advanced template features
go run templates/advanced_templates.go
```

#### Error Handling Examples
```bash
# Error handling patterns
go run error-handling/error_handling.go
```

## Example Categories

### Basic Examples (`basic/`)
- Simple provider creation
- Basic template usage
- Factory pattern usage
- Template variable substitution
- Error handling basics

### Multi-Provider Examples (`multi-provider/`)
- Comparing responses from different providers
- Provider-specific configurations
- Fallback strategies
- Provider selection logic

### Template Examples (`templates/`)
- Complex template creation
- Template validation
- Template composition
- JSON serialization
- Variable substitution patterns

### Error Handling Examples (`error-handling/`)
- Message validation errors
- Template validation errors
- Provider errors
- Error type checking
- Robust error handling patterns

### Real-World Examples (`real-world/`)
- Interactive chat bot with conversation history
- Content generation system for blogs, emails, and social media
- AI-powered code assistant for development tasks
- Integration patterns and best practices

## Environment Setup

### Required Environment Variables

```bash
# OpenAI
export OPENAI_API_KEY="sk-..."

# Anthropic
export ANTHROPIC_API_KEY="sk-ant-..."

# OpenRouter
export OPENROUTER_API_KEY="sk-or-..."

# Optional: Custom base URLs
export OLLAMA_BASE_URL="http://localhost:11434"
```

### Optional Setup

For local development with Ollama:
```bash
# Install Ollama
curl -fsSL https://ollama.ai/install.sh | sh

# Pull a model
ollama pull llama2

# Start Ollama server
ollama serve
```

## Common Patterns

### Creating a Provider
```go
// OpenAI
provider := llm.NewOpenAI(
    llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
    llm.WithTimeout(30*time.Second),
)

// Using factory pattern
provider, err := llm.NewProvider(llm.ProviderOpenAI,
    llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
)
```

### Creating Templates
```go
// Simple template
template := template.From(
    message.FromSystem("You are a helpful assistant."),
    message.FromUser("Hello, {{name}}!"),
)

// With variable substitution
invokedTemplate := template.Invoke(map[string]any{
    "name": "Alice",
})
```

### Error Handling
```go
response, err := provider.Invoke(ctx, template)
if err != nil {
    // Handle different error types
    switch {
    case message.IsValidationError(err):
        // Handle validation errors
    case message.IsTemplateError(err):
        // Handle template errors
    default:
        // Handle other errors
    }
    return
}
```

## Contributing

When adding new examples:

1. Create a new file in the appropriate directory
2. Use descriptive function names
3. Include comprehensive comments
4. Handle errors appropriately
5. Update this README if adding new categories

## Troubleshooting

### Common Issues

1. **API Key Errors**: Ensure environment variables are set correctly
2. **Timeout Errors**: Increase timeout values for slower connections
3. **Validation Errors**: Check template and message content
4. **Network Errors**: Verify internet connection and API endpoints

### Debug Mode

For debugging, you can enable verbose logging by setting:
```bash
export TARS_DEBUG=1
```

## Next Steps

After running the examples, try:

1. Modifying templates with your own variables
2. Testing with different providers
3. Implementing error handling in your applications
4. Creating custom templates for your use cases
5. Exploring the real-world examples for integration patterns 