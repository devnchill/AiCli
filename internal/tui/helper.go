package tui

import "strings"

func renderHistory(history []Message) string {
	var b strings.Builder

	for _, msg := range history {
		switch msg.Role {
		case RoleUSER:
			b.WriteString("You: ")
		case RoleLLM:
			b.WriteString("LLM: ")
		case RoleSYSTEM:
			b.WriteString("System: ")
		}

		b.WriteString(msg.Content)
		b.WriteString("\n")
	}

	return b.String()
}
