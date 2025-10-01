package store

import (
	"context"
	"log"
	"time"

	"github.com/za4em/openchat/data/datasource/api"
	"github.com/za4em/openchat/data/datasource/db"
	"github.com/za4em/openchat/domain"
)

type ChatStore struct {
	Api *api.OpenRouterApi
	DB  *db.Queries
	Ctx context.Context
}

func (store *ChatStore) GetChats() ([]domain.Chat, error) {
	var chats []domain.Chat
	dbChats, err := store.DB.GetChats(store.Ctx)
	if err != nil {
		return nil, domain.ErrStorageFailure(err)
	}
	for _, dbChat := range dbChats {
		dbMessages, err := store.DB.GetMessagesByChatID(store.Ctx, dbChat.ID)
		if err != nil {
			return nil, domain.ErrStorageFailure(err)
		}
		var messages []domain.Message
		for _, dbMessage := range dbMessages {
			messages = append(messages, domain.Message{
				ID:      dbMessage.ID,
				Role:    domain.Role(dbMessage.Role),
				Text:    dbMessage.Text,
				Created: time.Unix(dbMessage.CreatedAt, 0),
			})
		}

		chats = append(chats, domain.Chat{
			ID:       dbChat.ID,
			Name:     dbChat.Name,
			Messages: messages,
			Updated:  time.Unix(dbChat.UpdatedAt, 0),
		})
	}
	return chats, nil
}

func (store *ChatStore) CreateChat(input string) ([]domain.Chat, error) {
	chat := domain.NewChat(input)
	err := store.createChat(chat)
	if err != nil {
		return nil, err
	}
	chats, err := store.SendMessage(input, chat)
	return chats, err
}

func (store *ChatStore) SendMessage(input string, chat *domain.Chat) ([]domain.Chat, error) {
	userMessage := domain.NewMessage(domain.User, input)
	chat.Messages = append(chat.Messages, *userMessage)
	err := store.saveMessages(chat, []domain.Message{*userMessage})
	if err != nil {
		return nil, err
	}

	messages := convertDomainMessagesToApi(chat.Messages)
	request := store.Api.NewChatCompletionRequest(messages)
	response, err := store.Api.SendMessage(request)
	if err != nil {
		return nil, handleApiError(err)
	}

	responseMessages := convertApiMessagesToDomain(response.Choices)
	chat.Messages = append(chat.Messages, responseMessages...)
	err = store.saveMessages(chat, responseMessages)
	if err != nil {
		return nil, err
	}
	return store.GetChats()
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

func (store *ChatStore) createChat(chat *domain.Chat) error {
	_, err := store.DB.CreateChat(store.Ctx, db.CreateChatParams{
		ID:        chat.ID,
		UpdatedAt: chat.Updated.Unix(),
		Name:      chat.Name,
	})
	if err != nil {
		return domain.ErrStorageFailure(err)
	}
	return nil
}

func (store *ChatStore) saveChat(chat *domain.Chat) error {
	err := store.DB.UpdateChat(store.Ctx, db.UpdateChatParams{
		ID:        chat.ID,
		UpdatedAt: chat.Updated.Unix(),
		Name:      chat.Name,
	})
	if err != nil {
		log.Fatal(err)
		return domain.ErrStorageFailure(err)
	}
	return nil
}

func (store *ChatStore) saveMessages(chat *domain.Chat, messages []domain.Message) error {
	var err error
	if len(messages) == 0 {
		return nil
	}
	for _, message := range messages {
		_, err = store.DB.CreateMessage(store.Ctx, db.CreateMessageParams{
			ID:        message.ID,
			ChatID:    chat.ID,
			Role:      string(message.Role),
			Text:      message.Text,
			CreatedAt: message.Created.Unix(),
		})
	}
	chat.UpdateDatetime()
	store.saveChat(chat)
	if err != nil {
		return domain.ErrStorageFailure(err)
	}
	return nil
}

func handleApiError(err error) error {
	log.Println(err)
	return domain.ErrUnexpectedAPIResponse(err)
}
