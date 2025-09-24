package api

import (
	"github.com/za4em/openchat/config"
	"resty.dev/v3"
)

type XaiApi struct {
	config.Config
}

func CreateXaiApi(config config.Config) *XaiApi {
	return &XaiApi{
		Config: config,
	}
}

type ChatRequest struct {
	Input  string `json:"input"`
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

func (api *XaiApi) CreateChatRequest(input string) *ChatRequest {
	return &ChatRequest{
		Input:  input,
		Model:  api.DefaultModel,
		Stream: api.DefaultStream,
	}
}

type ContinueChatRequest struct {
	ChatRequest
	PreviousResponseId string `json:"previous_response_id"`
}

func (api *XaiApi) CreateContinueChatRequest(input string, responseId string) *ContinueChatRequest {
	return &ContinueChatRequest{
		ChatRequest:        *api.CreateChatRequest(input),
		PreviousResponseId: responseId,
	}
}

/*
 * type can be:
 * "message" - final answer, has content
 * "reasoning" - reasoning process, has summary
 * "function_call" - a tool call
 */
type Output struct {
	Content []Content `json:"content,omitempty"`
	ID      string    `json:"id"`
	Role    string    `json:"role,omitempty"`
	Type    string    `json:"type"`
	Status  string    `json:"status"`
	Summary []Content `json:"summary,omitempty"`
}

type ChatResponse struct {
	CreatedAt          int      `json:"created_at"`
	ID                 string   `json:"id"`
	Output             []Output `json:"output"`
	PreviousResponseID any      `json:"previous_response_id"`
	Status             string   `json:"status"`
}

type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// https://docs.x.ai/docs/api-reference#create-new-response
func (config XaiApi) CreateChat(request *ChatRequest) (*ChatResponse, error) {
	client := resty.New()
	defer client.Close()

	var response ChatResponse
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+config.API_KEY).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(config.API_URL + "/responses")

	return &response, err
}

func (config XaiApi) ContinueChat(request *ContinueChatRequest) (*ChatResponse, error) {
	client := resty.New()
	defer client.Close()
	client.SetHeader("Authorization", "Bearer "+config.API_KEY)
	client.SetHeader("Content-Type", "application/json")

	var response ChatResponse
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+config.API_KEY).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(config.API_URL + "/responses/{id}")

	return &response, err
}
