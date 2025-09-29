package domain

import (
	"github.com/google/uuid"
)

const ChatTitleLength = 30
const ChatDescriptionLength = 30

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
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

func NewMessage(input string) *Message {
	return &Message{
		ID:   uuid.NewString(),
		Text: input,
	}
}

func NewChat(message Message) *Chat {
	var name string
	if len(message.Text) > ChatTitleLength {
		name = message.Text[:ChatTitleLength]
	} else {
		name = message.Text
	}
	return &Chat{
		ID:       uuid.NewString(),
		Name:     name,
		Messages: []Message{message},
	}
}

type ChatStore interface {
	GetChats() []Chat
	CreateChat(input string) (*Chat, error)
	SendMessage(input string, chat *Chat) error
}



func (message Message) FilterValue() string {
	return message.Text + " " + message.Response.Text
}

func (message Message) Title() string       { return message.Text }
func (message Message) Description() string { return message.Response.Text }
