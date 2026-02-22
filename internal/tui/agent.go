package tui

import (
	"context"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/devnchill/AiCli/internal/providers"
)

type Message struct {
	Role providers.Role
	Text string
}

type Provider interface {
	SendPrompt(ctx context.Context, prompt string) (string, error)
	UpdateHistory(role providers.Role, content string)
}

type Agent struct {
	Name          string
	UIChatHistory []Message
	ViewPort      *viewport.Model
	Loading       bool
	Provider      Provider
}

func (a *Agent) SendPrompt(ctx context.Context, prompt string) (response string, err error) {
	return a.Provider.SendPrompt(ctx, prompt)
}

func (a *Agent) UpdateHistory(role providers.Role, content string) {
	a.UIChatHistory = append(a.UIChatHistory, Message{Role: role, Text: content})
	if role == providers.RoleSystem {
		return
	}
	a.Provider.UpdateHistory(role, content)
}
