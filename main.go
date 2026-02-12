package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	lipGloss "github.com/charmbracelet/lipgloss"
)

type model struct {
	displayHeight int
	displayWidth  int
	agents        map[string]string
	paneHeight    int
	paneWidth     int
	chatHistory   map[string]string
	currentMsg    string
}

func initialModel() model {
	return model{
		currentMsg: "",
		agents: map[string]string{
			"chatGPT": "OPEN_AI_APIKEY",
			"claude":  "CLAUDE_PIKEY",
		},
	}
}

func (m model) Init() tea.Cmd { return nil }

func apiCall(s string) {
	fmt.Println("Not implemented yet")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			{
				fmt.Println("exiting gracefully")
				return m, tea.Quit
			}
		case "enter":
			{
				// send msg to agents and get the response
				apiCall(m.currentMsg)
				m.chatHistory = append(m.chatHistory, m.currentMsg)
				m.currentMsg = ""
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	var s strings.Builder
	for key := range m.agents {
		s.WriteString(key + " ")
	}
	for _, chat := range m.chatHistory {
		s.WriteString(chat)
	}
	s.WriteString(">" + m.currentMsg)
	return s.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error :%v", err)
		os.Exit(1)
	}
}
