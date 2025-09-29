package ui

import (
	"github.com/charmbracelet/bubbles/list"
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
	list            list.Model
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
	lastMessage := chat.Messages[len(chat.Messages)-1]
	lastResponse := lastMessage.Response
	responseText := "Loading response"
	if lastResponse != nil {
		responseText = lastResponse.Text
	}
	responseTextTrimmed := responseText
	if len(responseText) > domain.ChatDescriptionLength {
		responseTextTrimmed = responseText[:30]
	}
	return responseTextTrimmed
}

func NewModel(chatStore domain.ChatStore) ChatModel {
	chats := chatStore.GetChats()
	items := make([]list.Item, len(chats))
	for i := range chats {
		items[i] = ChatItem{Chat: &chats[i]}
	}
	listModel := list.New(items, list.NewDefaultDelegate(), 0, 0)
	listModel.Title = "OPENCHAT"
	listModel.SetShowPagination(false)

	return ChatModel{
		ChatStore:       chatStore,
		view:            mainView,
		focus:           listFocus,
		chats:           chats,
		list:            listModel,
		currentChat:     nil,
		responseLoading: false,
		textInput:       textinput.New(),
		error:           "",
	}
}

func (model ChatModel) Init() tea.Cmd {
	// todo check if API key is provided, request if needed
	return nil
}
