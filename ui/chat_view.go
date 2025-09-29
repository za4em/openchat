package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	sidebarStyle = lipgloss.NewStyle().Width(30).Border(lipgloss.RoundedBorder())
	chatStyle    = lipgloss.NewStyle().Width(60).Padding(1)
	inputStyle   = lipgloss.NewStyle().Width(60).Border(lipgloss.RoundedBorder())
)

func (model ChatModel) View() string {

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
		sidebarStyle.Render(model.list.View()),
		main,
	)

	return body
}
