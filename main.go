package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type ChatMessage struct {
	User string
	Text string
}

type model struct {
	displayHeight int
	displayWidth  int
	agents        map[string]string
	paneHeight    int
	paneWidth     int
	chatHistory   []ChatMessage
	currentMsg    string
}

func initialModel() model {
	return model{
		currentMsg: "",
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
	}
}

func (m model) Init() tea.Cmd { return nil }

func apiCall(s string) {
	fmt.Println("Not implemented yet")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			{
				fmt.Println("exiting gracefully")
				return m, tea.Quit
			}
		case "enter":
			{
				// send msg to agents and get the response
				apiCall(m.currentMsg)
				m.chatHistory = append(m.chatHistory, ChatMessage{User: "user", Text: m.currentMsg})
				m.currentMsg = ""
			}
		}

	case tea.WindowSizeMsg:
		{
			m.displayHeight = msg.Height - 2
			m.displayWidth = msg.Width - 2
			m.paneHeight = m.displayHeight - 6
			m.paneWidth = (m.displayWidth / len(m.agents)) - 2
		}
	}
	return m, nil
}

func (m model) View() string {
	parentContainer := lipgloss.NewStyle().
		Height(m.displayHeight).
		Width(m.displayWidth).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#FFFFFF"))
	var agentViews []string
	for agentNames, _ := range m.agents {
		agentContent := fmt.Sprintf("Agent: %s", agentNames)
		agentPane := lipgloss.NewStyle().Height(m.paneHeight).Width(m.paneWidth).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#FFFFFF")).Padding(1, 1).Render(agentContent)
		agentViews = append(agentViews, agentPane)
	}
	horizontalRow := lipgloss.JoinHorizontal(lipgloss.Top, agentViews...)
	return parentContainer.Render(horizontalRow)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error :%v", err)
		os.Exit(1)
	}
}
