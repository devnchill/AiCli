package agent

import (
	"context"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
)

type Role int

const (
	RoleUSER Role = iota
	RoleLLM
	RoleSystem
)

type Message struct {
	Role Role
	Text string
}

type Provider interface {
	SendPrompt(ctx context.Context, prompt string) (string, error)
	UpdateHistory(role Role, content string)
}

type Agent struct {
	Name          string
	UIChatHistory []Message
	ViewPort      *viewport.Model
	Loading       bool
	Provider      Provider
}

func CreateProviders() (*GeminiProvider, *ClaudeProvider) {
	ctx := context.Background()
	geminiProvider, err := NewGeminiProvider(ctx, os.Getenv("GEMINI_API_KEY"))
	if err != nil {
		log.Fatal(err)
	}
	claudeProvider := NewClaudeProvider(ctx, os.Getenv("CLAUDE_API_KEY"))
	return geminiProvider, claudeProvider
}

func (a *Agent) SendPrompt(ctx context.Context, prompt string) (response string, err error) {
	return a.Provider.SendPrompt(ctx, prompt)
}

func (a *Agent) UpdateHistory(role Role, content string) {
	a.UIChatHistory = append(a.UIChatHistory, Message{Role: role, Text: content})
	if role == RoleSystem {
		return
	}
	a.Provider.UpdateHistory(role, content)
}
