package llm

import "fmt"

// ProviderType represents the type of LLM provider.
// This enum ensures type safety when specifying provider types
// and prevents errors from typos or invalid provider names.
type ProviderType string

const (
	// ProviderOpenAI represents the OpenAI provider.
	// Supports models like GPT-4, GPT-3.5-turbo, etc.
	ProviderOpenAI ProviderType = "openai"

	// ProviderAnthropic represents the Anthropic provider.
	// Supports models like Claude-3, Claude-2, etc.
	ProviderAnthropic ProviderType = "anthropic"

	// ProviderOpenRouter represents the OpenRouter provider.
	// Provides access to multiple LLM providers through a unified API.
	ProviderOpenRouter ProviderType = "openrouter"

	// ProviderOllama represents the Ollama provider.
	// Supports local LLM models like Llama, Mistral, etc.
	ProviderOllama ProviderType = "ollama"
)

// NewProvider creates a new LLM provider based on the provider type.
// This factory function abstracts provider creation and ensures
// consistent initialization across different provider types.
//
// The provider type determines which specific implementation is created,
// while the options allow customization of the provider's behavior.
//
// Example:
//
//	provider, err := NewProvider(ProviderOpenAI,
//	  WithAPIKey(os.Getenv("OPENAI_API_KEY")),
//	  WithTimeout(30*time.Second),
//	)
//	if err != nil {
//	  log.Fatal(err)
//	}
func NewProvider(providerType ProviderType, options ...LLMOption) (BaseProvider, error) {
	switch providerType {
	case ProviderOpenAI:
		return NewOpenAI(options...), nil
	case ProviderAnthropic:
		return NewAnthropic(options...), nil
	case ProviderOpenRouter:
		return NewOpenRouter(options...), nil
	case ProviderOllama:
		return NewOllama(options...), nil
	default:
		return nil, fmt.Errorf("unsupported provider type: %s", providerType)
	}
}

// GetSupportedProviders returns a list of all supported provider types.
// This is useful for validation, documentation, and dynamic provider selection.
//
// Example:
//
//	providers := GetSupportedProviders()
//	for _, provider := range providers {
//	  fmt.Printf("Supported provider: %s\n", provider)
//	}
func GetSupportedProviders() []ProviderType {
	return []ProviderType{
		ProviderOpenAI,
		ProviderAnthropic,
		ProviderOpenRouter,
		ProviderOllama,
	}
}
