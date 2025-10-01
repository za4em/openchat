package domain

import (
	"time"

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
	ID      string
	Role    Role
	Text    string
	Created time.Time
}

type Chat struct {
	ID       string
	Name     string
	Messages []Message
	Updated  time.Time
}

func NewMessage(role Role, text string) *Message {
	return &Message{
		ID:      uuid.NewString(),
		Role:    role,
		Text:    text,
		Created: time.Now(),
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
		Updated:  time.Now(),
	}
}

func (chat *Chat) UpdateDatetime() {
	chat.Updated = time.Now()
}

type ChatStore interface {
	GetChats() ([]Chat, error)
	CreateChat(input string) ([]Chat, error)
	SendMessage(input string, chat *Chat) ([]Chat, error)
}

func (message Message) FilterValue() string {
	return message.Text
}

func (message Message) Title() string       { return message.Text }
func (message Message) Description() string { return message.Text }
