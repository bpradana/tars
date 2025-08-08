# Real-World Examples

This directory contains real-world application examples that demonstrate how to use the Tars library in practical scenarios.

## Examples

### Chat Bot (`chat_bot.go`)
A conversational AI chat bot that maintains conversation history and provides interactive chat functionality.

**Features:**
- Conversation history management
- Interactive command-line interface
- Context-aware responses
- Error handling

**Usage:**
```bash
go run real-world/chat_bot.go
```

**Example interaction:**
```
You: Hello, how are you?
Bot: I'm doing well, thank you for asking! How can I help you today?

You: What's the weather like?
Bot: I don't have access to real-time weather data, but I can help you with other questions or tasks. Is there anything specific you'd like to know about?
```

### Content Generator (`content_generator.go`)
A content generation system that creates various types of content using AI.

**Features:**
- Blog post generation
- Email composition
- Social media post creation
- Product description writing

**Usage:**
```bash
go run real-world/content_generator.go
```

**Example outputs:**
- Professional blog posts on any topic
- Business emails with proper formatting
- Platform-optimized social media posts
- Compelling product descriptions

### Code Assistant (`code_assistant.go`)
An AI-powered code assistant for development tasks.

**Features:**
- Code generation from requirements
- Code debugging and error fixing
- Code review and improvement suggestions
- Code explanation and documentation

**Usage:**
```bash
go run real-world/code_assistant.go
```

**Example capabilities:**
- Generate Python functions from descriptions
- Debug code with error messages
- Review code for best practices
- Explain complex algorithms

## Common Patterns

### Template Management
All examples use the template system for consistent AI interactions:

```go
template := template.From(
    message.FromSystem("You are a helpful assistant."),
    message.FromUser("{{.UserInput}}"),
)

invokedTemplate := template.Invoke(struct {
    UserInput string
} {
    UserInput: "Hello",
}
```

### Error Handling
Robust error handling is implemented across all examples:

```go
response, err := provider.Invoke(ctx, template)
if err != nil {
    log.Printf("Error: %v", err)
    return
}
```

### Context Management
Proper context usage for timeouts and cancellation:

```go
ctx := context.Background()
// or with timeout:
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()
```

## Integration Examples

### Web Application Integration
```go
func handleChatRequest(w http.ResponseWriter, r *http.Request) {
    provider := llm.NewOpenAI(
        llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
    )
    
    bot := NewChatBot(provider)
    
    userInput := r.FormValue("message")
    response, err := bot.SendMessage(r.Context(), userInput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    json.NewEncoder(w).Encode(map[string]string{
        "response": response,
    })
}
```

### Batch Processing
```go
func processBatch(requests []Request) []Response {
    provider := llm.NewOpenAI(
        llm.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
    )
    
    generator := NewContentGenerator(provider)
    responses := make([]Response, len(requests))
    
    for i, req := range requests {
        content, err := generator.GenerateBlogPost(
            context.Background(),
            req.Topic,
            req.Style,
            req.Tone,
        )
        if err != nil {
            responses[i] = Response{Error: err.Error()}
        } else {
            responses[i] = Response{Content: content}
        }
    }
    
    return responses
}
```

## Best Practices

### 1. Template Design
- Use clear, specific system prompts
- Include relevant context in user messages
- Validate templates before use
- Use variable substitution for dynamic content

### 2. Error Handling
- Always check for errors from provider.Invoke()
- Implement retry logic for transient failures
- Provide meaningful error messages to users
- Log errors for debugging

### 3. Performance
- Use appropriate timeouts
- Implement caching for repeated requests
- Consider rate limiting for API providers
- Use context cancellation for long-running requests

### 4. Security
- Never hardcode API keys
- Use environment variables for configuration
- Validate user input before sending to AI
- Implement proper access controls

## Customization

### Adding New Content Types
```go
func (cg *ContentGenerator) GeneratePressRelease(ctx context.Context, company, announcement string) (string, error) {
    template := template.From(
        message.FromSystem("You are a PR specialist. Write professional press releases."),
        message.FromUser("Write a press release for {{.Company}} announcing {{.Announcement}}."),
    )
    
    invokedTemplate := template.Invoke(struct {
        Company      string
        Announcement string
    } {
        Company:      "Company",
        Announcement: "Announcement",
    })
    
    response, err := cg.provider.Invoke(ctx, invokedTemplate)
    if err != nil {
        return "", fmt.Errorf("failed to generate press release: %v", err)
    }
    
    return response.GetContent(), nil
}
```

### Custom Chat Bot Features
```go
func (cb *ChatBot) SetPersonality(personality string) {
    cb.template = template.From(
        message.FromSystem(fmt.Sprintf("You are a helpful assistant with the following personality: %s", personality)),
    )
}

func (cb *ChatBot) ClearHistory() {
    cb.history = nil
}
```

## Troubleshooting

### Common Issues

1. **API Key Errors**: Ensure environment variables are set correctly
2. **Timeout Errors**: Increase timeout values for complex requests
3. **Template Errors**: Validate templates before use
4. **Memory Issues**: Limit conversation history for long-running bots

### Debug Mode
Enable debug logging by setting:
```bash
export TARS_DEBUG=1
```

## Next Steps

After exploring these examples:

1. **Extend the chat bot** with additional features like file uploads or integrations
2. **Create custom content generators** for your specific use cases
3. **Build a code review system** that integrates with your development workflow
4. **Implement a multi-provider system** for redundancy and cost optimization
5. **Add streaming responses** for real-time interactions 