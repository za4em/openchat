package domain

import (
	"github.com/google/uuid"
)

const ChatTitleLength = 30
const ChatDescriptionLength = 30

type Role string

const (
	System    Role = "system"
	Developer Role = "developer"
	User      Role = "user"
	Assistant Role = "assistant"
	Tool      Role = "tool"
)

type Message struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
	Text string `json:"text"`
}

type Chat struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Messages []Message `json:"messages"`
}

func NewMessage(role Role, text string) *Message {
	return &Message{
		ID:   uuid.NewString(),
		Role: User,
		Text: text,
	}
}

func NewChat(input string) *Chat {
	name := input
	if len(input) > ChatTitleLength {
		name = input[:ChatTitleLength]
	}
	return &Chat{
		ID:       uuid.NewString(),
		Name:     name,
		Messages: []Message{},
	}
}

type ChatStore interface {
	GetChats() []Chat
	CreateChat(input string) (*Chat, error)
	SendMessage(input string, chat *Chat) error
}

func (message Message) FilterValue() string {
	return message.Text
}

func (message Message) Title() string       { return message.Text }
func (message Message) Description() string { return message.Text }
