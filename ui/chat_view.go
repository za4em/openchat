package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle  = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	sidebarStyle = lipgloss.NewStyle().Width(30).Border(lipgloss.RoundedBorder())
	chatStyle    = lipgloss.NewStyle().Width(60).Padding(1)
	inputStyle   = lipgloss.NewStyle().Width(60).Border(lipgloss.RoundedBorder())
)

func (model ChatModel) View() string {
	header := headerStyle.Render("OPENCHAT") + "\n\n"

	sidebar := ""
	for i, chat := range model.chats {
		suffix := " "
		if i == model.listIndex {
			suffix = ">"
		}
		sidebar += suffix + chat.Name + "\n"
	}

	chat := ""
	if model.currentChat != nil {
		for _, message := range model.currentChat.Messages {
			chat += message.Text + "\n"
			if message.Response != nil {
				chat += message.Response.Text + "\n"
			}
		}
	}

	main := lipgloss.JoinVertical(
		lipgloss.Bottom,
		chatStyle.Render(chat),
		inputStyle.Render(model.textInput.View()),
	)

	body := lipgloss.JoinHorizontal(
		lipgloss.Left,
		sidebarStyle.Render(sidebar),
		main,
	)

	return lipgloss.JoinVertical(
		lipgloss.Top,
		header,
		body,
	)
}
