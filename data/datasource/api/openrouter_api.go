package api

import (
	"github.com/za4em/openchat/config"
	"resty.dev/v3"
)

type OpenRouterApi struct {
	config.Config
}

func NewOpenRouterApi(config config.Config) *OpenRouterApi {
	return &OpenRouterApi{
		Config: config,
	}
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream,omitempty"`
}

type ChatCompletionResponse struct {
	ID      string   `json:"id"`
	Choices []Choice `json:"choices"`
	Usage   Usage    `json:"usage"`
	// For now we set time.Now()
	// Created int      `json:"created"`
}

type Choice struct {
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func (api *OpenRouterApi) NewChatCompletionRequest(messages []Message) *ChatCompletionRequest {
	return &ChatCompletionRequest{
		Model:    api.DefaultModel,
		Messages: messages,
		Stream:   api.DefaultStream,
	}
}

func (api *OpenRouterApi) SendMessage(request *ChatCompletionRequest) (*ChatCompletionResponse, error) {
	client := resty.New()
	defer client.Close()

	var response ChatCompletionResponse
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+api.API_KEY).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(api.API_URL + "/chat/completions")

	return &response, err
}
