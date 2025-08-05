package message

type messageOptions struct {
	usage usage
}

type MessageOption func(*messageOptions)

func WithUsage(promptTokens int, completionTokens int, totalTokens int) MessageOption {
	return func(m *messageOptions) {
		m.usage = usage{
			PromptTokens:     promptTokens,
			CompletionTokens: completionTokens,
			TotalTokens:      totalTokens,
		}
	}
}
