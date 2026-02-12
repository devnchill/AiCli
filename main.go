package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipgloss "github.com/charmbracelet/lipgloss"
)

type ChatMessage struct {
	User string
	Text string
}

type NameToApiKey map[string]string

type ChatHistory map[string][]string

type model struct {
	displayHeight  int
	displayWidth   int
	paneHeight     int
	paneWidth      int
	inputBarHeight int
	inputBarWidth  int
	inputBarPrompt string
	agents         NameToApiKey
	chatHistory    ChatHistory
	currentMsg     string
}

func initialModel() model {
	return model{
		currentMsg: "",
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
		chatHistory: map[string][]string{
			"chatGPT": {"hi , my name is GPT", "This is my second message", "Oh i like to keep talking", "alr i'll stop"},
			"claude":  {"hi , my name is CLAUDE", "I don't like to talk much", "Talk isn't cheap hence i cost you money", "I'm better than GPT slop."},
		},
		inputBarPrompt: "Prompt> ",
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
				m.chatHistory["user"] = append(m.chatHistory["user"], m.currentMsg)
				m.currentMsg = ""
			}
		}

	case tea.WindowSizeMsg:
		{
			m.displayHeight = msg.Height - 5
			m.displayWidth = msg.Width - 2
			m.inputBarHeight = 2
			m.inputBarWidth = m.displayWidth - 2
			m.paneHeight = m.displayHeight - m.inputBarHeight
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
	for agentNames := range m.agents {
		agentContent := fmt.Sprintf("Agent: %s\n>  %s", agentNames, strings.Join(m.chatHistory[agentNames], "\n>  "))
		agentPane := lipgloss.NewStyle().Height(m.paneHeight).Width(m.paneWidth).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#FFFFFF")).Padding(1, 1).Render(agentContent)
		agentViews = append(agentViews, agentPane)
	}
	horizontalRow := lipgloss.JoinHorizontal(lipgloss.Top, agentViews...)
	inputBar := lipgloss.NewStyle().Width(m.inputBarWidth).Border(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("#FFFFFF")).Render(m.inputBarPrompt)
	insideView := lipgloss.JoinVertical(lipgloss.Left, horizontalRow, inputBar)
	return parentContainer.Render(insideView)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error :%v", err)
		os.Exit(1)
	}
}
