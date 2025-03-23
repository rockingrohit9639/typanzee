package typanzee

import "github.com/charmbracelet/lipgloss"

// Colors
const (
	wrongCharColor   = lipgloss.Color("#FF0000")
	correctCharColor = lipgloss.Color("#008000")
	currTypingColor  = lipgloss.Color("")
)

// Styles
var (
	wrongChar   = lipgloss.NewStyle().Foreground(wrongCharColor)
	correctChar = lipgloss.NewStyle().Foreground(correctCharColor)

	cursorStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderRightForeground(correctCharColor).
			PaddingBottom(0)
)
