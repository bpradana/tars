package llm

import (
	"context"
	"time"

	"github.com/bpradana/failsafe"
	"github.com/bpradana/failsafe/strategies"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/errorbank"
	"github.com/bpradana/tars/pkg/httpx"
	"github.com/bpradana/tars/template"
)

// AnthropicProvider implements the BaseProvider interface for Anthropic
type AnthropicProvider struct {
	baseProvider
}

// NewAnthropic creates a new Anthropic provider
func NewAnthropic(options ...LLMOption) BaseProvider {
	opts := llmOptions{
		baseURL:     "https://api.anthropic.com",
		timeout:     10 * time.Second,
		maxAttempts: 1,
		maxDelay:    0 * time.Second,
	}

	for _, option := range options {
		option(&opts)
	}

	return &AnthropicProvider{
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
func (a *AnthropicProvider) GetName() string {
	return "anthropic"
}

// Invoke implements the BaseProvider interface for Anthropic
func (a *AnthropicProvider) Invoke(ctx context.Context, template template.Template, options ...InvokeOption) (message.Message, error) {
	// Validate the template before processing
	if err := template.Validate(); err != nil {
		return nil, errorbank.NewMessageError("template_validation", "invalid template provided", err)
	}

	opts := invokeOptions{
		model:       "claude-3-5-sonnet-20240620",
		temperature: 0.7,
		maxTokens:   1000,
	}
	for _, option := range options {
		option(&opts)
	}

	// Validate required configuration
	if a.options.apiKey == "" {
		return nil, errorbank.NewValidationError("api_key", "Anthropic API key is required", "")
	}

	resp, err := failsafe.RetryWithResult(ctx, a.retrier, func() (*httpx.Response, error) {
		return a.client.Post("/chat/completions", ChatCompletionsRequest{
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

	return message.FromAssistant(
		result.Choices[0].Message.Content,
		message.WithUsage(
			result.Usage.PromptTokens,
			result.Usage.CompletionTokens,
			result.Usage.TotalTokens,
		),
	), nil
}
