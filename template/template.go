package template

import (
	"encoding/json"

	"github.com/bpradana/tars/message"
)

type template struct {
	Message []message.Message
}

type Template interface {
	GetMessage() []message.Message
	Invoke(v map[string]any) Template
	ToJSON() string
}

func From(messages ...message.Message) Template {
	return template{
		Message: messages,
	}
}

func (t template) GetMessage() []message.Message {
	return t.Message
}

func (t template) Invoke(v map[string]any) Template {
	return template{
		Message: func() []message.Message {
			message := make([]message.Message, len(t.Message))
			for i, m := range t.Message {
				message[i] = m.Invoke(v)
			}
			return message
		}(),
	}
}

func (t template) ToJSON() string {
	json, err := json.Marshal(t.Message)
	if err != nil {
		return ""
	}
	return string(json)
}
