package llm

import (
	"context"
	"time"

	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/httpx"
	"github.com/bpradana/tars/template"
	"github.com/pkg/errors"
)

// OpenRouterProvider implements the BaseProvider interface for OpenRouter
type OpenRouterProvider struct {
	baseProvider
}

// NewOpenRouter creates a new OpenRouter provider
func NewOpenRouter(options ...LLMOption) BaseProvider {
	opts := llmOptions{
		baseURL: "https://openrouter.ai/api/v1",
		timeout: 10 * time.Second,
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
		},
	}
}

// GetName returns the provider name
func (o *OpenRouterProvider) GetName() string {
	return "openrouter"
}

// Invoke implements the BaseProvider interface for OpenRouter
func (o *OpenRouterProvider) Invoke(ctx context.Context, template template.Template, options ...InvokeOption) (message.Message, error) {
	opts := invokeOptions{
		model:       "google/gemini-2.0-flash-001",
		temperature: 0.7,
		maxTokens:   1000,
	}
	for _, option := range options {
		option(&opts)
	}

	resp, err := o.client.Post("/chat/completions", ChatCompletionsRequest{
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
	if err != nil {
		return nil, errors.Wrap(err, "failed to create request")
	}
	defer resp.Body.Close()

	var result ChatCompletionsResponse
	if err := resp.Decode(&result); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	if len(result.Choices) == 0 {
		return nil, errors.New("no choices in response")
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
