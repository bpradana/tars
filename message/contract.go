package message

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Message interface {
	GetRole() RoleType
	GetContent() string
	GetUsage() usage
	Invoke(v map[string]any) Message
	ToJSON() string
}

type usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

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

func (m message) Invoke(v map[string]any) Message {
	return message{
		Role: m.Role,
		Content: func() string {
			for k, v := range v {
				m.Content = strings.ReplaceAll(m.Content, fmt.Sprintf("{{%s}}", k), fmt.Sprintf("%v", v))
			}
			return m.Content
		}(),
	}
}

func (m message) ToJSON() string {
	json, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(json)
}
