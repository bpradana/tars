package message

type RoleType string

const (
	RoleSystem    RoleType = "system"    // System role is used to set the behavior of the assistant.
	RoleUser      RoleType = "user"      // User role is used to send a message to the assistant.
	RoleAssistant RoleType = "assistant" // Assistant role is used to send a message from the assistant.
)
