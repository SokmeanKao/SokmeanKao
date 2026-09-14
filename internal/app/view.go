package app

import tea "charm.land/bubbletea/v2"

// View is implemented fully in view.go composition; stub satisfies tea.Model for Update tests.
func (m Model) View() tea.View {
	v := tea.NewView("Initializing...")
	v.AltScreen = true
	return v
}
