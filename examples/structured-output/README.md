# Structured Output Examples

This directory contains examples demonstrating TARS's structured output functionality, which allows you to get consistent, typed responses from LLM providers using JSON schemas.

## Overview

Structured output enables you to:
- Get consistent, predictable responses from LLM providers
- Parse responses directly into Go structs
- Ensure type safety for LLM responses
- Build applications that rely on structured data

## Examples

### Weather Information (`weatherExample`)
Demonstrates how to extract weather data into a structured format:
- Temperature, condition, humidity, wind speed
- Automatic JSON schema generation from Go structs
- Error handling for structured responses

### Person Profile (`personExample`)
Shows how to generate structured person profiles:
- Name, age, city, job, hobbies
- Array handling in structured output
- Realistic data generation

### Product Catalog (`productExample`)
Illustrates product data extraction:
- Product name, price, category, description
- Boolean and numeric field handling
- Stock status and ratings

### Text Analysis (`analysisExample`)
Demonstrates text analysis with structured results:
- Sentiment analysis with confidence scores
- Key points extraction
- Summary and recommendations
- Array and nested object handling

### Error Handling (`errorHandlingExample`)
Shows how to handle errors when using structured output:
- Validation errors
- Invalid template handling
- Graceful error recovery

## Key Concepts

### JSON Schema Generation
TARS automatically generates JSON schemas from your Go structs using reflection:

```go
type WeatherInfo struct {
    Temperature float64 `json:"temperature"`
    Condition   string  `json:"condition"`
    Humidity    int     `json:"humidity"`
}
```

### Temperature Control
Use lower temperature values (0.2-0.3) for more consistent structured output:

```go
response, err := provider.Invoke(ctx, template,
    llm.WithStructuredOutput(&weatherInfo),
    llm.WithTemperature(0.3), // Lower for consistency
)
```

### Error Handling
Always validate structured output before using it:

```go
if weatherInfo.Temperature == 0 && weatherInfo.Condition == "" {
    log.Printf("Invalid structured output received")
    return
}
```

## Best Practices

1. **Clear Instructions**: Provide explicit instructions about the expected format
2. **Lower Temperature**: Use 0.2-0.3 for more consistent results
3. **Validation**: Always validate structured output before use
4. **Error Handling**: Handle cases where LLM doesn't return valid JSON
5. **Schema Design**: Design structs to be clear and unambiguous

## Running the Examples

```bash
# Run all structured output examples
go run structured_output_example.go

# Make sure you have API keys set
export OPENAI_API_KEY="your-api-key"
```

## Supported Providers

Structured output works with all supported providers:
- OpenAI (GPT-4, GPT-3.5-turbo)
- Anthropic (Claude models)
- OpenRouter (multiple providers)
- Ollama (local models)

## Use Cases

- **Data Extraction**: Extract structured data from unstructured text
- **Form Processing**: Parse form responses into typed structures
- **API Integration**: Generate consistent API responses
- **Content Analysis**: Analyze text with structured insights
- **Profile Generation**: Create structured user or product profiles
