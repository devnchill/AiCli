package tui

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/devnchill/AiCli/internal/providers"
	"github.com/joho/godotenv"
)

func renderHistory(history []Message) string {
	var b strings.Builder

	for _, msg := range history {
		switch msg.Role {
		case providers.RoleUSER:
			b.WriteString("You: ")
		case providers.RoleLLM:
			b.WriteString("LLM: ")
		case providers.RoleSystem:
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
