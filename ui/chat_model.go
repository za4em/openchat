package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/za4em/openchat/domain"
)

type view int

const (
	enterApiToken view = iota
	main
)

type focus int

const (
	sidebar focus = iota
	chat
)

type ChatModel struct {
	ChatStore       domain.ChatStore
	view            view
	focus           focus
	responseLoading bool
	chats           []domain.Chat
	sidebar         list.Model
	currentChat     *domain.Chat
	textInput       textinput.Model
	error           string
}

type ChatItem struct {
	*domain.Chat
}

func (chat ChatItem) FilterValue() string {
	return chat.Name + " " + chat.Messages[0].Text
}

func (chat ChatItem) Title() string { return chat.Name }
func (chat ChatItem) Description() string {
	responseText := "Loading response"
	lastMessage := chat.Messages[len(chat.Messages)-1]
	if len(chat.Messages) != 0 {
		responseText = lastMessage.Text
	}
	responseTextTrimmed := responseText
	if len(responseText) > domain.ChatDescriptionLength {
		responseTextTrimmed = responseText[:30]
	}
	return responseTextTrimmed
}

func NewModel(chatStore domain.ChatStore) ChatModel {
	chats, err := chatStore.GetChats()
	if err != nil {
		chats = []domain.Chat{}
	}
	items := chatsToListItem(chats)
	listModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "OPENCHAT"
	listModel.SetShowPagination(false)
	model := ChatModel{
		ChatStore:       chatStore,
		view:            main,
		focus:           sidebar,
		chats:           chats,
		sidebar:         listModel,
		currentChat:     nil,
		responseLoading: false,
		textInput:       textinput.New(),
		error:           err.Error(),
	}
	model.textInput.Width = chatWidth

	return model
}

func chatsToListItem(chats []domain.Chat) []list.Item {
	items := make([]list.Item, len(chats))
	for i := range chats {
		items[i] = ChatItem{Chat: &chats[i]}
	}
	return items
}

func (model ChatModel) Init() tea.Cmd {
	// todo check if API key is provided, request if needed
	return nil
}
