package message

// FromSystem creates a new system message with the given content.
// System messages are used to set the behavior and context for the assistant.
// They are typically sent at the beginning of a conversation to define
// the assistant's personality, capabilities, or constraints.
//
// Example:
//
//	msg := FromSystem("You are a helpful assistant that specializes in math.")
func FromSystem(content string) Message {
	if content == "" {
		// Return a message that will fail validation rather than panic
		return &message{
			Role:    RoleSystem,
			Content: "",
		}
	}

	return &message{
		Role:    RoleSystem,
		Content: content,
	}
}

// FromUser creates a new user message with the given content.
// User messages represent input from the user that the assistant should respond to.
// These messages can contain questions, requests, or any other user input.
//
// Example:
//
//	msg := FromUser("What is the capital of France?")
func FromUser(content string) Message {
	if content == "" {
		// Return a message that will fail validation rather than panic
		return &message{
			Role:    RoleUser,
			Content: "",
		}
	}

	return &message{
		Role:    RoleUser,
		Content: content,
	}
}

// FromAssistant creates a new assistant message with the given content and optional usage information.
// Assistant messages represent responses from the AI assistant to user input.
// The usage information is typically provided by the LLM provider and tracks
// token consumption for billing and monitoring purposes.
//
// Example:
//
//	msg := FromAssistant("The capital of France is Paris.",
//	  WithUsage(10, 5, 15))
func FromAssistant(content string, options ...MessageOption) Message {
	opts := messageOptions{}
	for _, option := range options {
		option(&opts)
	}

	return &message{
		Role:    RoleAssistant,
		Content: content,
		Usage:   opts.usage,
	}
}
