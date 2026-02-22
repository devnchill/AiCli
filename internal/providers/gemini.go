package providers

import (
	"context"
	"fmt"

	"google.golang.org/genai"
)

type GeminiProvider struct {
	APIKey         string
	APIChatHistory []*genai.Content
	Client         *genai.Client
}

func NewGeminiProvider(ctx context.Context, apiKey string) (*GeminiProvider, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		return nil, err
	}

	return &GeminiProvider{
		APIKey: apiKey,
		Client: client,
	}, nil
}

func convertRoleToGenAIRole(role Role) genai.Role {
	if role == RoleLLM {
		return genai.RoleModel
	}
	return genai.RoleUser
}

func (g *GeminiProvider) UpdateHistory(role Role, content string) {
	g.APIChatHistory = append(g.APIChatHistory, genai.NewContentFromText(content, convertRoleToGenAIRole(role)))
}

func (g *GeminiProvider) SendPrompt(ctx context.Context, prompt string) (string, error) {
	chat, err := g.Client.Chats.Create(
		ctx,
		"gemini-3-flash-preview",
		nil,
		g.APIChatHistory,
	)
	if err != nil {
		return "", err
	}

	res, err := chat.SendMessage(ctx, genai.Part{Text: prompt})
	if err != nil {
		return "", err
	}

	if len(res.Candidates) == 0 ||
		len(res.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response")
	}

	reply := res.Candidates[0].Content.Parts[0].Text

	return reply, nil
}
