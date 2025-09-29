package store

import (
	"log"

	"github.com/za4em/openchat/data/datasource/api"
	"github.com/za4em/openchat/data/datasource/storage"
	"github.com/za4em/openchat/domain"
)

type ChatStore struct {
	Api     *api.XaiApi
	Storage *storage.ChatStorage
}

func (store *ChatStore) GetChats() []domain.Chat {
	var chats []domain.Chat
	for _, chat := range store.Storage.Chats {
		chats = append(chats, chat)
	}
	return chats
}

func (store *ChatStore) CreateChat(input string) (*domain.Chat, error) {
	message, chat := store.createMessageAndChat(input)
	err := store.saveChat(chat)
	if err != nil {
		return nil, err
	}

	request := store.Api.NewChatRequest(input)
	response, err := store.Api.CreateChat(request)
	if err != nil {
		return nil, handleApiError(err)
	}

	message.Response = convertApiResponse(response)
	err = store.saveChat(chat)
	return chat, err
}

func (store *ChatStore) SendMessage(input string, chat *domain.Chat) error {
	lastResponse, err := getLastResponse(chat)
	if err != nil {
		return err
	}

	message := addNewMessageToChat(input, chat)
	err = store.saveChat(chat)
	if err != nil {
		return err
	}

	request := store.Api.NewContinueChatRequest(input, lastResponse.ID)
	response, err := store.Api.ContinueChat(request)
	if err != nil {
		return handleApiError(err)
	}

	message.Response = convertApiResponse(response)
	err = store.saveChat(chat)
	return err
}

func (store *ChatStore) createMessageAndChat(input string) (*domain.Message, *domain.Chat) {
	var chat *domain.Chat
	for _, c := range store.Storage.Chats {
		if len(c.Messages) == 0 {
			chat = &c
			break
		}
	}
	message := domain.NewMessage(input)
	if chat == nil {
		chat = domain.NewChat(*message)
	} else {
		chat.Messages = append(chat.Messages, *message)
	}
	return message, chat
}

func getLastResponse(chat *domain.Chat) (*domain.Response, error) {
	lastResponse := chat.Messages[len(chat.Messages)-1].Response
	if lastResponse == nil {
		return nil, domain.ErrUnableToSendMessage("Wait for the response or delete previous message")
	}
	return lastResponse, nil
}

func addNewMessageToChat(input string, chat *domain.Chat) *domain.Message {
	message := domain.NewMessage(input)
	chat.Messages = append(chat.Messages, *message)
	return message
}

func (store *ChatStore) saveChat(chat *domain.Chat) error {
	err := store.Storage.Save(chat)
	if err != nil {
		return domain.ErrStorageFailure(err)
	}
	return nil
}

func handleApiError(err error) error {
	log.Println(err)
	return domain.ErrUnexpectedAPIResponse(err)
}

func convertApiResponse(response *api.ChatResponse) *domain.Response {
	return &domain.Response{
		ID:   response.ID,
		Text: getContentFromResponse(response),
	}
}

func getContentFromResponse(response *api.ChatResponse) string {
	for _, output := range response.Output {
		if output.Type == "message" {
			for _, content := range output.Content {
				if len(content.Text) != 0 {
					return content.Text
				}
			}
		}
	}
	return "Model responded with empty string"
}
