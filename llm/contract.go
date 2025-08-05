package llm

import (
	"context"

	"github.com/bpradana/failsafe"
	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/httpx"
	"github.com/bpradana/tars/template"
)

// BaseProvider defines the interface that all LLM providers must implement.
// It provides a unified way to interact with different LLM services
// regardless of their specific API implementations.
//
// The interface abstracts away the differences between providers like
// OpenAI, Anthropic, OpenRouter, and Ollama, allowing applications
// to switch between providers without changing their core logic.
type BaseProvider interface {
	// Invoke sends a template to the LLM provider and returns the response.
	// The template contains the conversation context and user input.
	// Options can be used to customize the request (model, temperature, etc.).
	//
	// The context can be used for cancellation and timeout control.
	// The template should contain the full conversation context including
	// system messages, user messages, and any previous assistant responses.
	Invoke(ctx context.Context, template template.Template, options ...InvokeOption) (message.Message, error)

	// GetName returns the provider's name for identification and logging.
	// This is useful for debugging, monitoring, and provider-specific logic.
	GetName() string
}

// baseProvider is the base struct that all providers inherit from.
// It contains common functionality and configuration that is shared
// across all LLM provider implementations.
type baseProvider struct {
	options llmOptions
	client  *httpx.Client
	retrier *failsafe.Retrier
}

// GetName returns the provider name - to be overridden by each provider.
// This default implementation should not be used in production.
func (b *baseProvider) GetName() string {
	return "base"
}

// GetOptions returns the common options for the provider.
// This allows access to the provider's configuration for debugging
// and monitoring purposes.
func (b *baseProvider) GetOptions() llmOptions {
	return b.options
}

// SetOptions sets the common options for the provider.
// This is primarily used during provider initialization and
// should not be called after the provider is in use.
func (b *baseProvider) SetOptions(options llmOptions) {
	b.options = options
}
