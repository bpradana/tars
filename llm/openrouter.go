package llm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/bpradana/failsafe"
	"github.com/bpradana/failsafe/strategies"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/errorbank"
	"github.com/bpradana/tars/pkg/httpx"
	"github.com/bpradana/tars/template"
)

// OpenRouterProvider implements the BaseProvider interface for OpenRouter
type OpenRouterProvider struct {
	baseProvider
}

// NewOpenRouter creates a new OpenRouter provider
func NewOpenRouter(options ...LLMOption) BaseProvider {
	opts := llmOptions{
		baseURL:     "https://openrouter.ai/api/v1",
		timeout:     10 * time.Second,
		maxAttempts: 1,
		maxDelay:    0 * time.Second,
	}

	for _, option := range options {
		option(&opts)
	}

	return &OpenRouterProvider{
		baseProvider: baseProvider{
			options: opts,
			client: httpx.NewClient().
				WithBaseURL(opts.baseURL).
				WithDefaultHeaders(httpx.NewHeader().Bearer(opts.apiKey)).
				WithTimeout(opts.timeout),
			retrier: failsafe.NewRetrier(
				failsafe.WithMaxAttempts(opts.maxAttempts),
				failsafe.WithDelayStrategy(strategies.NewFixedDelay(opts.maxDelay)),
			),
		},
	}
}

// GetName returns the provider name
func (o *OpenRouterProvider) GetName() string {
	return "openrouter"
}

// Invoke implements the BaseProvider interface for OpenRouter
func (o *OpenRouterProvider) Invoke(ctx context.Context, template template.Template, options ...InvokeOption) (message.Message, error) {
	// Validate the template before processing
	if err := template.Validate(); err != nil {
		return nil, errorbank.NewMessageError("template_validation", "invalid template provided", err)
	}

	opts := invokeOptions{
		model:       "gpt-4o-mini",
		temperature: 0.7,
		maxTokens:   1000,
	}
	for _, option := range options {
		option(&opts)
	}

	// Validate required configuration
	if o.options.apiKey == "" {
		return nil, errorbank.NewValidationError("api_key", "OpenRouter API key is required", "")
	}

	resp, err := failsafe.RetryWithResult(ctx, o.retrier, func() (*httpx.Response, error) {
		return o.client.Post("/chat/completions", ChatCompletionsRequest{
			Model: opts.model,
			Messages: func() []Message {
				templateMessages := template.GetMessage()
				msgs := make([]Message, len(templateMessages))
				for i, msg := range templateMessages {
					msgs[i] = Message{
						Role:    string(msg.GetRole()),
						Content: msg.GetContent(),
					}
				}
				return msgs
			}(),
			ResponseFormat: func() *ResponseFormat {
				if opts.jsonSchema != nil {
					return &ResponseFormat{
						Type: "json_schema",
						JsonSchema: JsonSchema{
							Name:   "schema",
							Strict: true,
							Schema: opts.jsonSchema,
						},
					}
				}
				return nil
			}(),
		})
	})
	if err != nil {
		return nil, errorbank.NewMessageError("http_request", "failed to create request", err)
	}
	defer resp.Body.Close()

	var result ChatCompletionsResponse
	if err := resp.Decode(&result); err != nil {
		return nil, errorbank.NewMessageError("response_decode", "failed to decode response", err)
	}

	if len(result.Choices) == 0 {
		return nil, errorbank.NewMessageError("no_choices", "no choices in response", nil)
	}

	if opts.jsonSchema != nil {
		err = json.Unmarshal([]byte(result.Choices[0].Message.Content), opts.structuredOutput)
		if err != nil {
			return nil, errorbank.NewMessageError("json_unmarshal", "failed to unmarshal structured output", err)
		}
	}

	return message.FromAssistant(
		result.Choices[0].Message.Content,
		message.WithUsage(
			result.Usage.PromptTokens,
			result.Usage.CompletionTokens,
			result.Usage.TotalTokens,
		),
	), nil
}
