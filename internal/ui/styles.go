package ui

import "charm.land/lipgloss/v2"

var (
	ColorBackground  = lipgloss.Color("#0D1117")
	ColorPanel       = lipgloss.Color("#111827")
	ColorBorder      = lipgloss.Color("#1F6F50")
	ColorAccent      = lipgloss.Color("#00FF9C")
	ColorTextPrimary = lipgloss.Color("#D1D5DB")
	ColorTextMuted   = lipgloss.Color("#6B7280")
	ColorDanger      = lipgloss.Color("#FF5F56")
	ColorWarning     = lipgloss.Color("#FFBD2E")
	ColorSuccess     = lipgloss.Color("#00FF9C")
)

var (
	TitleStyle = lipgloss.NewStyle().Foreground(ColorAccent).Bold(true)
	DimStyle   = lipgloss.NewStyle().Foreground(ColorTextMuted)
	NormalStyle = lipgloss.NewStyle().Foreground(ColorTextPrimary)
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(ColorAccent).
			Bold(true)
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Background(ColorPanel)
)
