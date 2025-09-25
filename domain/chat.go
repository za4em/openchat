package domain

import (
	"github.com/google/uuid"
)

type Message struct {
	ID       string    `json:"id"`
	Text     string    `json:"text"`
	Response *Response `json:"response"`
}

type Response struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type Chat struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Messages []*Message `json:"messages"`
}

func NewMessage(input string) *Message {
	return &Message{
		ID:   uuid.NewString(),
		Text: input,
	}
}

func NewChat(message *Message) *Chat {
	return &Chat{
		ID:       uuid.NewString(),
		Name:     "New chat",
		Messages: []*Message{message},
	}
}

type ChatStore interface {
	GetChats() []*Chat
	GetChat(ID string) *Chat
	CreateChat(input string) (*Chat, error)
	SendMessage(input string, chat *Chat) error
}
