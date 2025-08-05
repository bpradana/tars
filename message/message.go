package message

func FromSystem(content string) Message {
	return &message{
		Role:    RoleSystem,
		Content: content,
	}
}

func FromUser(content string) Message {
	return &message{
		Role:    RoleUser,
		Content: content,
	}
}

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
