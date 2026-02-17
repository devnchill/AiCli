package tui

import (
	"github.com/devnchill/AiCli/internal/agent"
	"strings"
)

func renderHistory(history []agent.Message) string {
	var b strings.Builder

	for _, msg := range history {
		switch msg.Role {
		case agent.RoleUSER:
			b.WriteString("You: ")
		case agent.RoleLLM:
			b.WriteString("LLM: ")
		case agent.RoleSystem:
			b.WriteString("System: ")
		}

		b.WriteString(msg.Text)
		b.WriteString("\n")
	}

	return b.String()
}
