package message

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/bpradana/tars/pkg/errorbank"
)

// Message represents a conversation message with role, content, and usage information.
// It provides methods for template variable substitution and JSON serialization.
type Message interface {
	GetRole() RoleType
	GetContent() string
	GetUsage() usage
	Invoke(v map[string]any) Message
	ToJSON() string
	Validate() error
}

// usage tracks token usage information for LLM requests
type usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// message implements the Message interface
type message struct {
	Role    RoleType
	Content string
	Usage   usage
}

func (m message) GetRole() RoleType {
	return m.Role
}

func (m message) GetContent() string {
	return m.Content
}

func (m message) GetUsage() usage {
	return m.Usage
}

// Invoke performs template variable substitution on the message content.
// It creates a new message with substituted content without modifying the original.
func (m message) Invoke(v map[string]any) Message {
	if len(v) == 0 {
		return m
	}

	content := m.Content
	for k, val := range v {
		placeholder := fmt.Sprintf("{{%s}}", k)
		content = strings.ReplaceAll(content, placeholder, fmt.Sprintf("%v", val))
	}

	return message{
		Role:    m.Role,
		Content: content,
		Usage:   m.Usage,
	}
}

// ToJSON serializes the message to JSON string format.
// Returns an empty string if serialization fails.
func (m message) ToJSON() string {
	json, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(json)
}

// Validate checks if the message is valid and returns an error if not.
// This method can be used to validate messages before sending to LLM providers.
func (m message) Validate() error {
	if m.Role == "" {
		return errorbank.NewValidationError("role", "cannot be empty", m.Role)
	}

	if m.Content == "" {
		return errorbank.NewValidationError("content", "cannot be empty", m.Content)
	}

	// Validate role type
	switch m.Role {
	case RoleSystem, RoleUser, RoleAssistant:
		// Valid role
	default:
		return errorbank.NewValidationError("role", "invalid role type", m.Role)
	}

	return nil
}
