package tui

import (
	gloss "github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	switch m.Phase {
	case GreetingPhase:
		{
			return greetingPhase(m)
		}
	case ChatPhase:
		{
			return chatPhase(m)
		}
	}
	return errorPhase(m)
}

func greetingPhase(m model) string {
	return m.greetingState.greetingMessage
}

func chatPhase(m model) string {
	var panes []string

	for _, name := range m.chatState.agentsNameInOrder {
		if m.chatState.agents[name].ViewPort == nil {
			continue
		}
		styled := gloss.NewStyle().
			Border(gloss.NormalBorder()).
			BorderForeground(gloss.Color("#FFFFFF")).
			Render(m.chatState.agents[name].ViewPort.View())
		panes = append(panes, styled)
	}

	horizontalRow := gloss.JoinHorizontal(gloss.Top, panes...)
	insideView := gloss.JoinVertical(
		gloss.Left,
		horizontalRow,
		m.chatState.inputTextArea.View(),
	)

	parentContainer := gloss.NewStyle().
		Height(m.chatState.tuiHeight - 2).
		Width(m.chatState.tuiWidth - 2).
		Border(gloss.NormalBorder()).
		BorderForeground(gloss.Color("#FFFFFF"))

	return parentContainer.Render(insideView)
}

func errorPhase(m model) string { return "OOps an error occured" }
