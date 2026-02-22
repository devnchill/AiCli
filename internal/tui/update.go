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
	switch m.Phase {
	case GreetingPhase:
		{
		}
	case ChatPhase:
		{
			m.chatState.inputTextArea, cmd = m.chatState.inputTextArea.Update(msg)
			switch msg := msg.(type) {
			case tea.KeyMsg:
				switch msg.String() {
				case "ctrl+c":
					return m, tea.Quit
				case "enter":
					text := strings.TrimSpace(m.chatState.inputTextArea.Value())
					if text != "" {
						for _, agentStruct := range m.chatState.agents {
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
						m.chatState.inputTextArea.Reset()
					}
				}
			case tea.WindowSizeMsg:
				m.chatState.tuiHeight = msg.Height
				m.chatState.tuiWidth = msg.Width

				usableHeight := m.chatState.tuiHeight - 2
				usableWidth := m.chatState.tuiWidth - 2

				m.chatState.agentViewportHeight = usableHeight - m.chatState.inputTextAreaHeight
				m.chatState.agentViewportWidth = usableWidth / len(m.chatState.agents)

				for _, agentStruct := range m.chatState.agents {
					if agentStruct.ViewPort == nil {
						vp := viewport.New(m.chatState.agentViewportWidth-2, m.chatState.agentViewportHeight-2)
						agentStruct.ViewPort = &vp
					}
					vp := agentStruct.ViewPort
					vp.Width = m.chatState.agentViewportWidth - 2
					vp.Height = m.chatState.agentViewportHeight - 2
					m.chatState.inputTextArea.SetWidth(vp.Width)
					vp.SetContent(renderHistory(agentStruct.UIChatHistory))
				}
			}
		}
	}

	return m, cmd
}
