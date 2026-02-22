package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/devnchill/AiCli/internal/agent"
)

type Phase int

const (
	GreetingPhase Phase = iota
	ChatPhase
)

type option struct {
	name       string
	isSelected bool
}

type greetingPhaseStruct struct {
	greetingMessage string
	options         []option
	cursor          int
	selectedAgents  []option
}

type chatPhaseStruct struct {
	tuiHeight           int
	tuiWidth            int
	agents              map[string]*agent.Agent
	agentsNameInOrder   []string
	agentViewportHeight int
	agentViewportWidth  int
	inputTextAreaHeight int
	inputTextAreaWidth  int
	inputTextArea       textarea.Model
}

type model struct {
	Phase         Phase
	greetingState greetingPhaseStruct
	chatState     chatPhaseStruct
}

func InitialModel() model {
	loadEnv()
	greetingState := greetingPhaseStruct{
		greetingMessage: "hello welcome to ai cli",
	}
	chatState := chatPhaseStruct{}

	return model{
		Phase:         GreetingPhase,
		greetingState: greetingState,
		chatState:     chatState,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}
