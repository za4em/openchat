package api

import (
	"github.com/za4em/openchat/config"
	"resty.dev/v3"
)

type APIConfig struct {
	config.Config
}

type CreateChat struct {
	Input  string `json:"input"`
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

type ContinueChat struct {
	CreateChat
	PreviousResponseId int `json:"previous_response_id"`
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

/*
 * type can be:
 * "message" - final answer, has content
 * "reasoning" - reasoning process, has summary
 */
type Output struct {
	Content []Content `json:"content,omitempty"`
	ID      string    `json:"id"`
	Role    string    `json:"role,omitempty"`
	Type    string    `json:"type"`
	Status  string    `json:"status"`
	Summary []Content `json:"summary,omitempty"`
}

// https://docs.x.ai/docs/api-reference#create-new-response
func (config APIConfig) CreateChat(request CreateChat) (*ChatResponse, error) {
	client := resty.New()
	defer client.Close()

	var response ChatResponse
	_, err := client.R().
		SetHeader("Authorization", "Bearer "+config.API_KEY).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(config.API_URL + "/responses")

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (config APIConfig) ContinueChat(responseID string, request ContinueChat) (*ChatResponse, error) {
	client := resty.New()
	defer client.Close()
	client.SetHeader("Authorization", "Bearer "+config.API_KEY)
	client.SetHeader("Content-Type", "application/json")

	var response ChatResponse
	_, err := client.R().
		SetPathParam("id", responseID).
		SetHeader("Authorization", "Bearer "+config.API_KEY).
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&response).
		Post(config.API_URL + "/responses/{id}")

	if err != nil {
		return nil, err
	}

	return &response, nil
}
