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
type ChatHistory map[string][]Message
type AgentNameToViewPort map[string]*viewport.Model

type MessageRole int

const (
	RoleUSER MessageRole = iota
	RoleLLM
	RoleSYSTEM
)

type Message struct {
	Role    MessageRole
	Content string
}

type model struct {
	tuiHeight           int
	tuiWidth            int
	agents              NameToApiKey
	agentViewportHeight int
	agentViewportWidth  int
	agentViewports      AgentNameToViewPort
	inputTextAreaHeight int
	inputTextAreaWidth  int
	inputTextArea       textarea.Model
	chatHistory         ChatHistory
	agentsNameInOrder   []string
}

func renderHistory(history []Message) string {
	var b strings.Builder

	for _, msg := range history {
		switch msg.Role {
		case RoleUSER:
			b.WriteString("You: ")
		case RoleLLM:
			b.WriteString("LLM: ")
		case RoleSYSTEM:
			b.WriteString("System: ")
		}

		b.WriteString(msg.Content)
		b.WriteString("\n")
	}

	return b.String()
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Enter your prompt..."
	ta.Prompt = "| "
	ta.Focus()
	ta.SetHeight(3)
	ta.SetWidth(60)

	agentsNameInOrder := []string{"chatGPT", "claude"}

	return model{
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
		chatHistory: map[string][]Message{
			"chatGPT": {{Role: RoleLLM, Content: "hi , my name is GPT"}},
			"claude":  {{Role: RoleLLM, Content: "hi , my name is CLAUDE"}, {Role: RoleLLM, Content: "I don't like to talk much"}},
		},
		inputTextAreaHeight: 3,
		inputTextAreaWidth:  60,
		inputTextArea:       ta,
		agentViewports:      make(map[string]*viewport.Model),
		agentsNameInOrder:   agentsNameInOrder,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
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

func (m model) View() string {
	var panes []string

	for _, name := range m.agentsNameInOrder {
		vp, ok := m.agentViewports[name]
		if !ok || vp == nil {
			continue
		}
		styled := gloss.NewStyle().
			Border(gloss.NormalBorder()).
			BorderForeground(gloss.Color("#FFFFFF")).
			Render(vp.View())
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

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}
