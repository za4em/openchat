package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/za4em/openchat/domain"
)

func (model ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		_, y := sidebarStyle.GetFrameSize()
		model.sidebar.SetSize(sidebarWidth, msg.Height-y)
	case tea.KeyMsg:
		key := msg.String()
		//todo add model.view handling
		switch model.focus {
		case sidebar:
			switch key {
			case "q":
				return model, tea.Quit
			case "n":
				model.currentChat = nil
				model.textInput.SetValue("")
				model.textInput.Focus()
				model.focus = chat
			case "enter":
				listIndex := model.sidebar.Index()
				model.currentChat = &model.chats[listIndex]
				model.textInput.SetValue("")
				model.textInput.Focus()
				model.focus = chat
			}
		case chat:
			switch key {
			case "ctrl+q", "ctrl+left", "ctrl+h":
				model.focus = sidebar
				model.textInput.Blur()
			case "enter":
				input := model.textInput.Value()
				if len(input) != 0 {
					var err error
					var chats []domain.Chat
					if model.currentChat != nil {
						chats, err = model.ChatStore.SendMessage(input, model.currentChat)
					} else {
						chats, err = model.ChatStore.CreateChat(input)
					}
					if err != nil {
						log.Fatal(err)
						model.error = err.Error()
						//todo handle error
					} else {
						model.textInput.SetValue("")
						model.chats = chats
						model.sidebar.SetItems(chatsToListItem(chats))
					}
				}
			}
		}
	}

	switch model.focus {
	case chat:
		model.textInput, cmd = model.textInput.Update(msg)
		cmds = append(cmds, cmd)
	case sidebar:
		model.sidebar, cmd = model.sidebar.Update(msg)
		cmds = append(cmds, cmd)
	}
	return model, tea.Batch(cmds...)
}
