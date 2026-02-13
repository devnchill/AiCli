package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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

func InitialModel() model {
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
