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
}

func NewModel(chatStore domain.ChatStore) ChatModel {
	return ChatModel{
		ChatStore:       chatStore,
		view:            mainView,
		responseLoading: false,
	}
}

func (model ChatModel) Init() tea.Cmd {
	// todo check if API key is provided, request if needed
	return nil
}
