package tui

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/devnchill/AiCli/internal/agent"
	"github.com/devnchill/AiCli/internal/types"
)

type model struct {
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

func InitialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Enter your prompt..."
	ta.Prompt = "| "
	ta.Focus()
	ta.SetHeight(3)
	ta.SetWidth(60)

	agentsNameInOrder := []string{"chatGPT", "claude"}

	return model{
		agents: map[string]*agent.Agent{
			"chatGPT": &agent.Agent{Name: "CHATGPT", API_KEY: "GPT_KEY", LOADING: false, ChatHistory: []types.Message{{Role: types.RoleLLM, Content: "hi,myself gpt"}}, Vp: &viewport.Model{}},
			"claude":  &agent.Agent{Name: "CLAUDE", API_KEY: "CLAUDE_KEY", LOADING: false, ChatHistory: []types.Message{{Role: types.RoleLLM, Content: "hi,myself claude"}}, Vp: &viewport.Model{}},
		},
		inputTextAreaHeight: 3,
		inputTextAreaWidth:  60,
		inputTextArea:       ta,
		agentsNameInOrder:   agentsNameInOrder,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}
