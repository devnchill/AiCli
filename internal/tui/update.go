package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.inputTextArea, cmd = m.inputTextArea.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			text := strings.TrimSpace(m.inputTextArea.Value())
			if text != "" {
				for name, vp := range m.agentViewports {
					m.chatHistory[name] = append(m.chatHistory[name], Message{Role: RoleUSER, Content: text})
					vp.SetContent(renderHistory(m.chatHistory[name]))
				}
				m.inputTextArea.Reset()
			}
		}
	case tea.WindowSizeMsg:
		m.tuiHeight = msg.Height
		m.tuiWidth = msg.Width

		usableHeight := m.tuiHeight - 2
		usableWidth := m.tuiWidth - 2

		m.agentViewportHeight = usableHeight - m.inputTextAreaHeight
		m.agentViewportWidth = usableWidth / len(m.agents)

		for agentName := range m.agents {
			if _, exists := m.agentViewports[agentName]; !exists {
				vp := viewport.New(m.agentViewportWidth-2, m.agentViewportHeight-2)
				m.agentViewports[agentName] = &vp
			}
			vp := m.agentViewports[agentName]
			vp.Width = m.agentViewportWidth - 2
			vp.Height = m.agentViewportHeight - 2
			vp.SetContent(renderHistory(m.chatHistory[agentName]))
		}
	}

	return m, cmd
}
