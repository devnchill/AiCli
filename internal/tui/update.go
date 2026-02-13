package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/devnchill/AiCli/internal/types"
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
				for _, agentStruct := range m.agents {
					agentStruct.ChatHistory = append(agentStruct.ChatHistory, types.Message{Role: types.RoleUSER, Content: text})
					agentStruct.Vp.SetContent(renderHistory(agentStruct.ChatHistory))
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

		for _, agentStruct := range m.agents {
			if agentStruct.Vp == nil {
				vp := viewport.New(m.agentViewportWidth-2, m.agentViewportHeight-2)
				agentStruct.Vp = &vp
			}
			vp := agentStruct.Vp
			vp.Width = m.agentViewportWidth - 2
			vp.Height = m.agentViewportHeight - 2
			vp.SetContent(renderHistory(agentStruct.ChatHistory))
		}
	}

	return m, cmd
}
