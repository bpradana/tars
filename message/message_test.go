package message

import "testing"

func Test_FromSystem(t *testing.T) {
	message := FromSystem("You are a helpful assistant.")
	if message.GetRole() != RoleSystem {
		t.Errorf("Expected role to be %s, got %s", RoleSystem, message.GetRole())
	}
	if message.GetContent() != "You are a helpful assistant." {
		t.Errorf("Expected content to be %s, got %s", "You are a helpful assistant.", message.GetContent())
	}
}

func Test_FromUser(t *testing.T) {
	message := FromUser("What is the capital of France?")
	if message.GetRole() != RoleUser {
		t.Errorf("Expected role to be %s, got %s", RoleUser, message.GetRole())
	}
	if message.GetContent() != "What is the capital of France?" {
		t.Errorf("Expected content to be %s, got %s", "What is the capital of France?", message.GetContent())
	}
}

func Test_FromAssistant(t *testing.T) {
	message := FromAssistant("The capital of France is Paris.")
	if message.GetRole() != RoleAssistant {
		t.Errorf("Expected role to be %s, got %s", RoleAssistant, message.GetRole())
	}
	if message.GetContent() != "The capital of France is Paris." {
		t.Errorf("Expected content to be %s, got %s", "The capital of France is Paris.", message.GetContent())
	}
	if message.GetUsage().PromptTokens != 0 {
		t.Errorf("Expected usage to be %d, got %d", 0, message.GetUsage().PromptTokens)
	}
	if message.GetUsage().CompletionTokens != 0 {
		t.Errorf("Expected usage to be %d, got %d", 0, message.GetUsage().CompletionTokens)
	}
}

func Test_Invoke(t *testing.T) {
	message := FromUser("What is the capital of {{country}}?").Invoke(map[string]any{
		"country": "France",
	})
	if message.GetRole() != RoleUser {
		t.Errorf("Expected role to be %s, got %s", RoleUser, message.GetRole())
	}
	if message.GetContent() != "What is the capital of France?" {
		t.Errorf("Expected content to be %s, got %s", "What is the capital of France?", message.GetContent())
	}
}

func Test_ToJSON(t *testing.T) {
	message := FromUser("What is the capital of {{country}}?").Invoke(map[string]any{
		"country": "France",
	}).ToJSON()

	if message != `{"role":"user","content":"What is the capital of France?"}` {
		t.Errorf("Expected %s, got %s", `{"role":"user","content":"What is the capital of France?"}`, message)
	}
}
