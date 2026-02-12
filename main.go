package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type NameToApiKey map[string]string
type ChatHistory map[string][]string

type model struct {
	displayHeight int
	displayWidth  int
	paneHeight    int
	paneWidth     int
	agents        NameToApiKey
	chatHistory   ChatHistory
	input         textarea.Model
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Prompt> "
	ta.FocusedStyle.Placeholder = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF"))
	ta.SetWidth(60)
	ta.SetHeight(3)
	ta.Focus()

	return model{
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
		chatHistory: map[string][]string{
			"chatGPT": {"hi , my name is GPT", "This is my second message"},
			"claude":  {"hi , my name is CLAUDE", "I don't like to talk much"},
		},
		input: ta,
	}
}

// func apiCall(text string) {
// 	fmt.Println("API CALL:", text)
// }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			text := strings.TrimSpace(m.input.Value())
			if text != "" {
				m.chatHistory["user"] = append(m.chatHistory["user"], text)
				// apiCall(text)
				m.input.Reset()
			}
		}
	case tea.WindowSizeMsg:
		m.displayHeight = msg.Height - 4
		m.displayWidth = msg.Width - 2
		m.paneHeight = m.displayHeight - m.input.Height()
		m.paneWidth = (m.displayWidth / len(m.agents)) - 2
		m.input.SetWidth(m.displayWidth - 4)
	}

	return m, cmd
}

func (m model) View() string {
	var agentViews []string
	for agentName := range m.agents {
		historyText := strings.Join(m.chatHistory[agentName], "\n> ")
		content := fmt.Sprintf("Agent: %s\n> %s", agentName, historyText)

		agentPane := lipgloss.NewStyle().
			Height(m.paneHeight).
			Width(m.paneWidth).
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("#FFFFFF")).
			Padding(1, 1).
			Render(content)

		agentViews = append(agentViews, agentPane)
	}

	horizontalRow := lipgloss.JoinHorizontal(lipgloss.Top, agentViews...)

	insideView := lipgloss.JoinVertical(
		lipgloss.Left,
		horizontalRow,
		m.input.View(),
	)

	parentContainer := lipgloss.NewStyle().
		Height(m.displayHeight).
		Width(m.displayWidth).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#FFFFFF"))

	return parentContainer.Render(insideView)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
