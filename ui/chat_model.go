package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/za4em/openchat/domain"
)

const (
	enterApiTokenView = iota
	mainView
)

const (
	listFocus = iota
	chatFocus
)

type ChatModel struct {
	ChatStore       domain.ChatStore
	view            uint
	focus           uint
	responseLoading bool
	chats           []domain.Chat
	listIndex       int
	currentChat     *domain.Chat
	textInput       textinput.Model
	error           string
}

func NewModel(chatStore domain.ChatStore) ChatModel {
	chats := chatStore.GetChats()
	var currentChat *domain.Chat
	if len(chats) == 0 {
		currentChat = nil
	} else {
		currentChat = &chats[len(chats)-1]
	}
	return ChatModel{
		ChatStore:       chatStore,
		view:            mainView,
		focus:           listFocus,
		chats:           chats,
		listIndex:       0,
		currentChat:     currentChat,
		responseLoading: false,
		textInput:       textinput.New(),
		error:           "",
	}
}

func (model ChatModel) Init() tea.Cmd {
	// todo check if API key is provided, request if needed
	return nil
}
