package store

import (
	"log"

	"github.com/za4em/openchat/data/datasource/api"
	"github.com/za4em/openchat/data/datasource/storage"
	"github.com/za4em/openchat/domain"
)

type ChatStore struct {
	Api     *api.OpenRouterApi
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
	chat := domain.NewChat(input)
	err := store.SendMessage(input, chat)
	return chat, err
}

func (store *ChatStore) SendMessage(input string, chat *domain.Chat) error {
	userMessage := domain.NewMessage(domain.User, input)
	chat.Messages = append(chat.Messages, *userMessage)
	chat.UpdateDatetime()
	err := store.saveChat(chat)
	if err != nil {
		return err
	}

	messages := convertDomainMessagesToApi(chat.Messages)
	request := store.Api.NewChatCompletionRequest(messages)
	response, err := store.Api.SendMessage(request)
	if err != nil {
		return handleApiError(err)
	}

	responseMessages := convertApiMessagesToDomain(response.Choices)
	chat.Messages = append(chat.Messages, responseMessages...)
	chat.UpdateDatetime()
	err = store.saveChat(chat)
	return err
}

func convertDomainMessagesToApi(domainMessages []domain.Message) []api.Message {
	var apiMessages []api.Message
	for _, msg := range domainMessages {
		apiMessages = append(apiMessages, api.Message{Role: string(msg.Role), Content: msg.Text})
	}
	return apiMessages
}

func convertApiMessagesToDomain(apiMessages []api.Choice) []domain.Message {
	var domainMessages []domain.Message
	for _, choice := range apiMessages {
		msg := choice.Message
		domainMessages = append(domainMessages, *domain.NewMessage(domain.Role(msg.Role), msg.Content))
	}
	return domainMessages
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
