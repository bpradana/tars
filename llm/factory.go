package llm

import "fmt"

// ProviderType represents the type of LLM provider
type ProviderType string

const (
	ProviderOpenAI     ProviderType = "openai"
	ProviderAnthropic  ProviderType = "anthropic"
	ProviderOpenRouter ProviderType = "openrouter"
	ProviderOllama     ProviderType = "ollama"
)

// NewProvider creates a new LLM provider based on the provider type
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

// GetSupportedProviders returns a list of all supported provider types
func GetSupportedProviders() []ProviderType {
	return []ProviderType{
		ProviderOpenAI,
		ProviderAnthropic,
		ProviderOpenRouter,
		ProviderOllama,
	}
}
