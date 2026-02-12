package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	displayHeight int
	displayWidth  int
	agents        map[string]string
	paneHeight    int
	paneWidth     int
	chatHistory   []string
	currentMsg    string
}

func initialModel() model {
	return model{
		currentMsg:  "",
		chatHistory: []string{""},
	}
}

func (m model) Init() tea.Cmd { return nil }

func apiCall(s string) {
	panic("unimplemented")
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
	return "Hello"
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error :%v", err)
		os.Exit(1)
	}
}
