package providers

import (
	"context"
	"strings"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type ClaudeProvider struct {
	APIKey         string
	APIChatHistory []anthropic.MessageParam
	Client         *anthropic.Client
}

func NewClaudeProvider(ctx context.Context, apiKey string) *ClaudeProvider {
	client := anthropic.NewClient(
		option.WithAPIKey("ANTHROPIC_API_KEY"),
	)
	return &ClaudeProvider{
		APIKey: apiKey,
		Client: &client,
	}
}

func (c *ClaudeProvider) UpdateHistory(role Role, content string) {
	if role == RoleUSER {
		c.APIChatHistory = append(c.APIChatHistory, anthropic.NewUserMessage(anthropic.NewTextBlock(content)))
	} else {
		c.APIChatHistory = append(c.APIChatHistory, anthropic.NewAssistantMessage(anthropic.NewTextBlock(content)))
	}
}

func (c *ClaudeProvider) SendPrompt(ctx context.Context, prompt string) (string, error) {
	tempHistory := append(c.APIChatHistory, anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)))
	response, err := c.Client.Messages.New(ctx, anthropic.MessageNewParams{
		Messages:  tempHistory,
		MaxTokens: 1024,
		Model:     anthropic.ModelClaudeOpus4_6,
	})
	if err != nil {
		return "", err
	}
	var result strings.Builder
	for _, block := range response.Content {
		if block.Text != "" {
			result.WriteString(block.Text)
		}
	}
	return result.String(), nil
}
