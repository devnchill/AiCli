package tui

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/devnchill/AiCli/internal/agent"
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
					agentStruct.UpdateHistory(agent.RoleUSER, text)
					ctx := context.Background()
					res, err := agentStruct.SendPrompt(ctx, text)
					if err != nil {
						fmt.Println(err)
						os.Exit(0)
					}
					if len(res) > 0 {
						agentStruct.UpdateHistory(agent.RoleLLM, res)
						content := renderHistory(agentStruct.UIChatHistory)
						wrapped := lipgloss.NewStyle().Width(agentStruct.ViewPort.Width).Height(agentStruct.ViewPort.Height).Render(content)
						agentStruct.ViewPort.SetContent(wrapped)
					}
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
			if agentStruct.ViewPort == nil {
				vp := viewport.New(m.agentViewportWidth-2, m.agentViewportHeight-2)
				agentStruct.ViewPort = &vp
			}
			vp := agentStruct.ViewPort
			vp.Width = m.agentViewportWidth - 2
			vp.Height = m.agentViewportHeight - 2
			m.inputTextArea.SetWidth(vp.Width)
			vp.SetContent(renderHistory(agentStruct.UIChatHistory))
		}
	}

	return m, cmd
}
