package llm

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/invopop/jsonschema"
)

// llmOptions contains configuration options for LLM providers.
// This struct is used internally to collect options during provider initialization.
type llmOptions struct {
	baseURL     string
	apiKey      string
	timeout     time.Duration
	maxAttempts int
	maxDelay    time.Duration
}

// LLMOption is a function type that modifies LLM options.
// It follows the functional options pattern for flexible provider configuration.
type LLMOption func(*llmOptions)

// WithBaseURL sets the base URL for the LLM provider's API.
// This is useful for using custom endpoints or different API versions.
//
// Example:
//
//	provider := NewOpenAI(
//	  WithBaseURL("https://api.openai.com/v1"),
//	  WithAPIKey("your-api-key"),
//	)
func WithBaseURL(baseURL string) LLMOption {
	return func(llm *llmOptions) {
		llm.baseURL = baseURL
	}
}

// WithAPIKey sets the API key for authentication with the LLM provider.
// The API key is required for all LLM providers and should be kept secure.
//
// Example:
//
//	provider := NewOpenAI(
//	  WithAPIKey(os.Getenv("OPENAI_API_KEY")),
//	)
func WithAPIKey(apiKey string) LLMOption {
	return func(llm *llmOptions) {
		llm.apiKey = apiKey
	}
}

// WithTimeout sets the timeout for HTTP requests to the LLM provider.
// This prevents requests from hanging indefinitely and allows for
// proper error handling and retry logic.
//
// Example:
//
//	provider := NewOpenAI(
//	  WithTimeout(30 * time.Second),
//	)
func WithTimeout(timeout time.Duration) LLMOption {
	return func(llm *llmOptions) {
		llm.timeout = timeout
	}
}

// WithMaxAttempts sets the maximum number of attempts for the LLM provider.
// This is useful for retrying failed requests.
//
// Example:
//
//	provider := NewOpenAI(
//	  WithMaxAttempts(3),
//	)
func WithMaxAttempts(maxAttempts int) LLMOption {
	return func(llm *llmOptions) {
		llm.maxAttempts = maxAttempts
	}
}

// WithMaxDelay sets the maximum delay for the LLM provider.
// This is useful for retrying failed requests.
//
// Example:
//
//	provider := NewOpenAI(
//	  WithMaxDelay(10 * time.Second),
//	)
func WithMaxDelay(maxDelay time.Duration) LLMOption {
	return func(llm *llmOptions) {
		llm.maxDelay = maxDelay
	}
}

// invokeOptions contains configuration options for individual LLM requests.
// These options can be customized per request to control the model's behavior.
type invokeOptions struct {
	model            string
	temperature      float64
	maxTokens        int
	structuredOutput any
	jsonSchema       map[string]any
}

// InvokeOption is a function type that modifies invoke options.
// It allows customization of individual requests without affecting
// the provider's default configuration.
type InvokeOption func(*invokeOptions)

// WithModel sets the specific model to use for the request.
// Different models have different capabilities, costs, and performance characteristics.
//
// Example:
//
//	response, err := provider.Invoke(ctx, template,
//	  WithModel("gpt-4"),
//	)
func WithModel(model string) InvokeOption {
	return func(llm *invokeOptions) {
		llm.model = model
	}
}

// WithTemperature sets the temperature (randomness) for the model's response.
// Higher values (0.8-1.0) make responses more creative and unpredictable.
// Lower values (0.0-0.3) make responses more focused and deterministic.
//
// Example:
//
//	response, err := provider.Invoke(ctx, template,
//	  WithTemperature(0.7),
//	)
func WithTemperature(temperature float64) InvokeOption {
	return func(llm *invokeOptions) {
		llm.temperature = temperature
	}
}

// WithMaxTokens sets the maximum number of tokens in the response.
// This controls the length of the generated response and can help
// manage costs and response times.
//
// Example:
//
//	response, err := provider.Invoke(ctx, template,
//	  WithMaxTokens(1000),
//	)
func WithMaxTokens(maxTokens int) InvokeOption {
	return func(llm *invokeOptions) {
		llm.maxTokens = maxTokens
	}
}

// WithStructuredOutput sets the structured output for the request.
// The structured output is a pointer to a struct that will be used to unmarshal the response.
// This is useful for returning structured data from the model.
//
// Example:
//
//	response, err := provider.Invoke(ctx, template,
//	  WithStructuredOutput(&StructuredOutput{}),
//	)
func WithStructuredOutput(structuredOutput any) InvokeOption {
	return func(llm *invokeOptions) {
		llm.structuredOutput = structuredOutput

		llm.jsonSchema = func() map[string]any {
			schema := jsonschema.Reflect(structuredOutput)
			ref := strings.Split(schema.Ref, "#/$defs/")
			schemaDefinition, _ := schema.Definitions[ref[1]].MarshalJSON()
			var jsonSchema map[string]any
			_ = json.Unmarshal(schemaDefinition, &jsonSchema)
			return jsonSchema
		}()
	}
}
