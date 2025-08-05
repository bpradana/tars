package llm

import "time"

type llmOptions struct {
	baseURL string
	apiKey  string
	timeout time.Duration
}

type LLMOption func(*llmOptions)

func WithBaseURL(baseURL string) LLMOption {
	return func(llm *llmOptions) {
		llm.baseURL = baseURL
	}
}

func WithAPIKey(apiKey string) LLMOption {
	return func(llm *llmOptions) {
		llm.apiKey = apiKey
	}
}

func WithTimeout(timeout time.Duration) LLMOption {
	return func(llm *llmOptions) {
		llm.timeout = timeout
	}
}

type invokeOptions struct {
	model       string
	temperature float64
	maxTokens   int
}

type InvokeOption func(*invokeOptions)

func WithModel(model string) InvokeOption {
	return func(llm *invokeOptions) {
		llm.model = model
	}
}

func WithTemperature(temperature float64) InvokeOption {
	return func(llm *invokeOptions) {
		llm.temperature = temperature
	}
}

func WithMaxTokens(maxTokens int) InvokeOption {
	return func(llm *invokeOptions) {
		llm.maxTokens = maxTokens
	}
}
