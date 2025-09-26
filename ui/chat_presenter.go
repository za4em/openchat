package ui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (model ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
		cmd  tea.Cmd
	)

	model.textInput, cmd = model.textInput.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		//todo add model.view handling
		switch model.focus {
		case listFocus:
			switch key {
			case "q":
				return model, tea.Quit
			case "n":
				model.currentChat = nil
				model.textInput.SetValue("")
				model.textInput.Focus()
				model.focus = chatFocus
				// show input

			case "up", "k":
				if model.listIndex > 0 {
					model.listIndex--
				}

			case "down", "j":
				if model.listIndex < len(model.chats)-1 {
					model.listIndex++
				}

			case "enter":
				model.currentChat = model.chats[model.listIndex]
				model.textInput.SetValue("")
				model.textInput.Focus()
				model.focus = chatFocus
				// show input
			}

		case chatFocus:
			switch key {
			case "ctrl+q", "ctrl+left", "ctrl+h":
				model.focus = listFocus
				model.textInput.Blur()
			case "enter":
				input := model.textInput.Value()
				if len(input) != 0 {
					var err error
					if model.currentChat != nil {
						err = model.ChatStore.SendMessage(input, model.currentChat)
					} else {
						_, err = model.ChatStore.CreateChat(input)
					}
					if err != nil {
						log.Fatal(err)
						model.error = err.Error()
						//todo handle error
					} else {
						model.textInput.SetValue("")
					}
				}
			}
		}
	}
	return model, tea.Batch(cmds...)
}
