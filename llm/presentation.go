package llm

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Refusal string `json:"refusal"`
}

type Choice struct {
	Message      Message        `json:"message"`
	LogProbs     map[string]any `json:"logprobs"`
	FinishReason string         `json:"finish_reason"`
	Index        int            `json:"index"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type JsonSchema struct {
	Name   string         `json:"name"`
	Strict bool           `json:"strict"`
	Schema map[string]any `json:"schema"`
}

type ResponseFormat struct {
	Type       string     `json:"type"`
	JsonSchema JsonSchema `json:"json_schema"`
}

type ChatCompletionsRequest struct {
	Model          string          `json:"model"`
	Messages       []Message       `json:"messages"`
	ResponseFormat *ResponseFormat `json:"response_format,omitempty"`
}

type ChatCompletionsResponse struct {
	ID                string   `json:"id"`
	Choices           []Choice `json:"choices"`
	Provider          string   `json:"provider"`
	Model             string   `json:"model"`
	Object            string   `json:"object"`
	Created           int      `json:"created"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Usage             Usage    `json:"usage"`
}
