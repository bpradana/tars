package template

import (
	"encoding/json"
	"fmt"

	"github.com/bpradana/tars/message"
	"github.com/bpradana/tars/pkg/errorbank"
)

// template represents a conversation template that can be used with LLM providers.
// It contains a sequence of messages that form a conversation context.
type template struct {
	Message []message.Message
}

// Template defines the interface for conversation templates.
// Templates can be used to create reusable conversation patterns
// and perform variable substitution across multiple messages.
type Template interface {
	// GetMessage returns the list of messages in the template
	GetMessage() []message.Message

	// Invoke performs variable substitution on all messages in the template.
	// It creates a new template with substituted content without modifying the original.
	// The variables map should contain key-value pairs where keys correspond
	// to placeholder names in the template (e.g., "{{name}}").
	Invoke(v map[string]any) Template

	// ToJSON serializes the template to JSON string format.
	// Returns an empty string if serialization fails.
	ToJSON() string

	// Validate checks if the template is valid and returns an error if not.
	// This method validates all messages in the template.
	Validate() error
}

// From creates a new template from a sequence of messages.
// This is the primary constructor for creating conversation templates.
//
// Example:
//
//	template := From(
//	  message.FromSystem("You are a helpful assistant."),
//	  message.FromUser("Hello, {{name}}!"),
//	)
func From(messages ...message.Message) Template {
	return template{
		Message: messages,
	}
}

// GetMessage returns the list of messages in the template
func (t template) GetMessage() []message.Message {
	return t.Message
}

// Invoke performs variable substitution on all messages in the template.
// It creates a new template with substituted content without modifying the original.
// If the variables map is empty or nil, the original template is returned unchanged.
//
// Example:
//
//	result := template.Invoke(map[string]any{
//	  "name": "Alice",
//	  "city": "Paris",
//	})
func (t template) Invoke(v map[string]any) Template {
	if len(v) == 0 {
		return t
	}

	return template{
		Message: func() []message.Message {
			messages := make([]message.Message, len(t.Message))
			for i, m := range t.Message {
				messages[i] = m.Invoke(v)
			}
			return messages
		}(),
	}
}

// ToJSON serializes the template to JSON string format.
// Returns an empty string if serialization fails.
func (t template) ToJSON() string {
	json, err := json.Marshal(t.Message)
	if err != nil {
		return ""
	}
	return string(json)
}

// Validate checks if the template is valid and returns an error if not.
// This method validates all messages in the template.
func (t template) Validate() error {
	if len(t.Message) == 0 {
		return errorbank.NewValidationError("messages", "template cannot be empty", t.Message)
	}

	for i, msg := range t.Message {
		if err := msg.Validate(); err != nil {
			return errorbank.NewTemplateError(fmt.Sprintf("message[%d]", i), "validation failed", err)
		}
	}

	return nil
}
