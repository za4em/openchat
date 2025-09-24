package store

import (
	"log"

	"github.com/za4em/openchat/datasource/api"
	"github.com/za4em/openchat/datasource/storage"
	"github.com/za4em/openchat/domain"
)

type ChatStore struct {
	Api     *api.XaiApi
	Storage *storage.ChatStorage
}

func (store *ChatStore) GetChats() map[string]*domain.Chat {
	return store.Storage.Chats
}

func (store *ChatStore) GetChat(ID string) *domain.Chat {
	return store.Storage.Chats[ID]
}

// todo only create chat if there is no other empty chats
func (store *ChatStore) CreateChat(input string) (*domain.Chat, error) {
	message := domain.NewMessage(input)
	chat := domain.NewChat(message)
	error := store.Storage.Save(chat)
	if error != nil {
		return nil, domain.ErrorStorageFailure(error)
	}

	request := api.CreateChatRequest{
		Input:  input,
		Model:  store.Api.DefaultModel,
		Stream: store.Api.DefaultStream,
	}
	response, error := store.Api.CreateChat(request)
	if error != nil {
		log.Println(error)
		return nil, domain.ErrorUnexpectedAPIResponse(error)
	}

	message.Response = &domain.Response{
		ID:   response.ID,
		Text: getContentFromResponse(response),
	}

	error = store.Storage.Save(chat)
	return chat, error
}

func (store *ChatStore) SendMessage(input string, chat *domain.Chat) error {
	lastResponse := chat.Messages[len(chat.Messages)-1].Response
	if lastResponse == nil {
		return domain.ErrorUnableToSendMessage("Wait for the response or delete previous message")
	}

	message := domain.NewMessage(input)
	chat.Messages = append(chat.Messages, message)
	error := store.Storage.Save(chat)
	if error != nil {
		return domain.ErrorStorageFailure(error)
	}

	request := api.ContinueChatRequest{
		CreateChatRequest: api.CreateChatRequest{
			Input:  input,
			Model:  store.Api.DefaultModel,
			Stream: store.Api.DefaultStream,
		},
		PreviousResponseId: lastResponse.ID,
	}
	response, error := store.Api.ContinueChat(request)
	if error != nil {
		log.Println(error)
		return domain.ErrorUnexpectedAPIResponse(error)
	}

	message.Response = &domain.Response{
		ID:   response.ID,
		Text: getContentFromResponse(response),
	}

	error = store.Storage.Save(chat)
	return error
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
