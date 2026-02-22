package providers

import (
	"context"
	"log"
	"os"
)

type Role int

const (
	RoleUSER Role = iota
	RoleLLM
	RoleSystem
)

type Provider interface {
	SendPrompt(ctx context.Context, prompt string) (string, error)
	UpdateHistory(role Role, content string)
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
