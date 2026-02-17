package tui

import (
	gloss "github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	var panes []string

	for _, name := range m.agentsNameInOrder {
		if m.agents[name].ViewPort == nil {
			continue
		}
		styled := gloss.NewStyle().
			Border(gloss.NormalBorder()).
			BorderForeground(gloss.Color("#FFFFFF")).
			Render(m.agents[name].ViewPort.View())
		panes = append(panes, styled)
	}

	horizontalRow := gloss.JoinHorizontal(gloss.Top, panes...)
	insideView := gloss.JoinVertical(
		gloss.Left,
		horizontalRow,
		m.inputTextArea.View(),
	)

	parentContainer := gloss.NewStyle().
		Height(m.tuiHeight - 2).
		Width(m.tuiWidth - 2).
		Border(gloss.NormalBorder()).
		BorderForeground(gloss.Color("#FFFFFF"))

	return parentContainer.Render(insideView)
}
