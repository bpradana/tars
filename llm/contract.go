package llm

import (
	"context"

	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/httpx"
	"github.com/bpradana/tars/template"
)

// BaseProvider is the interface that all LLM providers must implement
type BaseProvider interface {
	Invoke(ctx context.Context, template template.Template, options ...InvokeOption) (message.Message, error)
	GetName() string
}

// baseProvider is the base struct that all providers inherit from
type baseProvider struct {
	options llmOptions
	client  *httpx.Client
}

// GetName returns the provider name - to be overridden by each provider
func (b *baseProvider) GetName() string {
	return "base"
}

// GetOptions returns the common options
func (b *baseProvider) GetOptions() llmOptions {
	return b.options
}

// SetOptions sets the common options
func (b *baseProvider) SetOptions(options llmOptions) {
	b.options = options
}
