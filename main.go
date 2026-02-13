package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	gloss "github.com/charmbracelet/lipgloss"
)

type NameToApiKey map[string]string
type ChatHistory map[string][]string
type AgentNameToViewPort map[string]viewport.Model

type model struct {
	tuiHeight           int
	tuiWidth            int
	agents              NameToApiKey
	agentViewportHeight int
	agentViewportWidth  int
	agentViewports      AgentNameToViewPort
	inputTextArea       textarea.Model
	chatHistory         ChatHistory
}

func initialModel() model {
	return model{
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
		chatHistory: map[string][]string{
			"chatGPT": {"hi , my name is GPT", "This is my second message"},
			"claude":  {"hi , my name is CLAUDE", "I don't like to talk much"},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

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
				m.chatHistory["user"] = append(m.chatHistory["user"], text)
				// apiCall(text)
				m.inputTextArea.Reset()
			}
		}
	case tea.WindowSizeMsg:
		m.tuiHeight = msg.Height - 4
		m.tuiWidth = msg.Width - 2
		m.agentViewportHeight = m.tuiHeight - m.inputTextArea.Height()
		m.agentViewportWidth = (m.tuiWidth / len(m.agents)) - 2
		m.inputTextArea.SetWidth(m.tuiWidth - 4)
	}

	return m, cmd
}

func (m model) View() string {
	var agentViews []string
	for agentName := range m.agents {
		historyText := strings.Join(m.chatHistory[agentName], "\n> ")
		content := fmt.Sprintf("Agent: %s\n> %s", agentName, historyText)

		agentPane := gloss.NewStyle().
			Height(m.agentViewportHeight).
			Width(m.agentViewportWidth).
			Border(gloss.NormalBorder()).
			BorderForeground(gloss.Color("#FFFFFF")).
			Padding(1, 1).
			Render(content)

		agentViews = append(agentViews, agentPane)
	}

	horizontalRow := gloss.JoinHorizontal(gloss.Top, agentViews...)

	insideView := gloss.JoinVertical(
		gloss.Left,
		horizontalRow,
		m.inputTextArea.View(),
	)

	parentContainer := gloss.NewStyle().
		Height(m.tuiHeight).
		Width(m.tuiWidth).
		Border(gloss.NormalBorder()).
		BorderForeground(gloss.Color("#FFFFFF"))

	return parentContainer.Render(insideView)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
