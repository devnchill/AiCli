package agent

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/devnchill/AiCli/internal/types"
)

type Agent struct {
	Name        string
	API_KEY     string
	LOADING     bool
	ChatHistory []types.Message
	Vp          *viewport.Model
}
