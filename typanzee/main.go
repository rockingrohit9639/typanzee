package typanzee

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	text  string // the text to render in terminal
	input string // input from user
}

func initModel() model {
	return model{
		text:  "This is typanzee",
		input: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Exit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		// Handle the backspace to remove last character typed
		case tea.KeyBackspace.String():
			if len(m.input) > 0 {
				m.input = m.input[:len(m.input)-1]
			}

		// Handle user input
		default:
			if len(m.input) < len(m.text) {
				m.input += msg.String()
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	width, height := TerminalSize()

	var renderedText strings.Builder
	for i, char := range m.text {
		targetChar := string(char)

		// Typing the current character
		if i == len(m.input) {
			renderedText.WriteString(cursorStyle.Render(targetChar))
		} else if i < len(m.input) {
			inputChar := string(m.input[i])

			// Check if user is entering the correct character
			if inputChar == targetChar {
				renderedText.WriteString(correctChar.Render(inputChar))
			} else {
				renderedText.WriteString(wrongChar.Render(inputChar))
			}
		} else {
			renderedText.WriteString(string(char))
		}

	}

	centered := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center)

	return centered.Render(renderedText.String())
}

func Start() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
