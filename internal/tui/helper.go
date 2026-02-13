package tui

import (
	"strings"

	"github.com/devnchill/AiCli/internal/types"
)

func renderHistory(history []types.Message) string {
	var b strings.Builder

	for _, msg := range history {
		switch msg.Role {
		case types.RoleUSER:
			b.WriteString("You: ")
		case types.RoleLLM:
			b.WriteString("LLM: ")
		case types.RoleSYSTEM:
			b.WriteString("System: ")
		}

		b.WriteString(msg.Content)
		b.WriteString("\n")
	}

	return b.String()
}
