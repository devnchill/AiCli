package tui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/devnchill/AiCli/internal/agent"
	"github.com/joho/godotenv"
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

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func createTextArea() textarea.Model {
	ta := textarea.New()
	ta.Placeholder = "Enter your prompt..."
	ta.Prompt = "| "
	ta.Focus()
	ta.SetHeight(3)
	ta.SetWidth(150)
	return ta
}
