package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	sidebarWidth int = 20
	chatWidth    int = 60
)

var (
	sidebarStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
	chatStyle    = lipgloss.NewStyle().Padding(1, 1, 1, 1)
	inputStyle   = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())
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
		lipgloss.Left,
		chatStyle.Render(chat),
		inputStyle.Render(model.textInput.View()),
	)

	body := lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		sidebarStyle.Render(model.sidebar.View()),
		main,
	)

	return body
}
