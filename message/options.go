package message

// messageOptions contains configuration options for message creation.
// This struct is used internally to collect options before creating a message.
type messageOptions struct {
	usage usage
}

// MessageOption is a function type that modifies message options.
// It follows the functional options pattern for flexible message configuration.
type MessageOption func(*messageOptions)

// WithUsage sets the token usage information for a message.
// This is typically used when creating assistant messages to track
// token consumption from LLM providers for billing and monitoring.
//
// Parameters:
//   - promptTokens: Number of tokens in the input/prompt
//   - completionTokens: Number of tokens in the response/completion
//   - totalTokens: Total number of tokens used in the request
//
// Example:
//
//	msg := FromAssistant("Response content",
//	  WithUsage(100, 50, 150))
func WithUsage(promptTokens int, completionTokens int, totalTokens int) MessageOption {
	return func(m *messageOptions) {
		m.usage = usage{
			PromptTokens:     promptTokens,
			CompletionTokens: completionTokens,
			TotalTokens:      totalTokens,
		}
	}
}
