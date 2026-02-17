package tui

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/devnchill/AiCli/internal/agent"
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
	ta.SetWidth(150)

	agentsNameInOrder := []string{"gemini", "claude"}
	ctx := context.Background()
	geminiProvider, err := agent.NewGeminiProvider(ctx, os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	claudeProvider := agent.NewClaudeProvider(ctx, os.Getenv("CLAUDE_API_KEY"))

	return model{
		agents: map[string]*agent.Agent{
			"gemini": {Name: "GEMINI", Loading: false, UIChatHistory: []agent.Message{{Role: agent.RoleLLM, Text: "hi,myself gemini"}}, ViewPort: &viewport.Model{}, Provider: geminiProvider},
			"claude": {Name: "CLAUDE", Loading: false, UIChatHistory: []agent.Message{{Role: agent.RoleLLM, Text: "hi,myself claude"}}, ViewPort: &viewport.Model{}, Provider: claudeProvider},
		},
		inputTextAreaHeight: 3,
		inputTextAreaWidth:  150,
		inputTextArea:       ta,
		agentsNameInOrder:   agentsNameInOrder,
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}
