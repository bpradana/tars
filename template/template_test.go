package template

import (
	"testing"

	"github.com/bpradana/tars/message"
)

func Test_Invoke(t *testing.T) {
	template := From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is the capital of {{country}}?"),
	).Invoke(map[string]any{
		"country": "France",
	})

	if template.GetMessage()[0].GetRole() != message.RoleSystem {
		t.Errorf("Expected role to be %s, got %s", message.RoleSystem, template.GetMessage()[0].GetRole())
	}

	if template.GetMessage()[0].GetContent() != "You are a helpful assistant." {
		t.Errorf("Expected content to be %s, got %s", "You are a helpful assistant.", template.GetMessage()[0].GetContent())
	}

	if template.GetMessage()[1].GetRole() != message.RoleUser {
		t.Errorf("Expected role to be %s, got %s", message.RoleUser, template.GetMessage()[1].GetRole())
	}

	if template.GetMessage()[1].GetContent() != "What is the capital of France?" {
		t.Errorf("Expected content to be %s, got %s", "What is the capital of France?", template.GetMessage()[1].GetContent())
	}
}

func Test_ToJSON(t *testing.T) {
	template := From(
		message.FromSystem("You are a helpful assistant."),
		message.FromUser("What is the capital of {{country}}?"),
	).Invoke(map[string]any{
		"country": "France",
	}).ToJSON()

	if template != `[{"role":"system","content":"You are a helpful assistant."},{"role":"user","content":"What is the capital of France?"}]` {
		t.Errorf("Expected %s, got %s", `[{"role":"system","content":"You are a helpful assistant."},{"role":"user","content":"What is the capital of France?"}]`, template)
	}
}
