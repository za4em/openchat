package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/za4em/openchat/domain"
)

const (
	enterApiTokenView = iota
	mainView
)

type ChatModel struct {
	ChatStore       domain.ChatStore
	view            uint
	responseLoading bool
	chats           []*domain.Chat
	currentChat     *domain.Chat
}

func NewModel(chatStore domain.ChatStore) ChatModel {
	chats := chatStore.GetChats()
	return ChatModel{
		ChatStore:       chatStore,
		view:            mainView,
		chats:           chats,
		currentChat:     chats[len(chats)-1],
		responseLoading: false,
	}
}

func (model ChatModel) Init() tea.Cmd {
	// todo check if API key is provided, request if needed
	return nil
}
