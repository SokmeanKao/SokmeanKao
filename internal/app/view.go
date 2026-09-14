package app

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/pages"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func (m Model) View() tea.View {
	if m.Width == 0 || m.Height == 0 {
		v := tea.NewView("Initializing...")
		v.AltScreen = true
		return v
	}

	r := ui.ComputeRegions(m.Width, m.Height)
	header := ui.Header(m.Width, m.Now.Format("15:04:05"))
	sidebar := ui.Sidebar(r.SidebarW, r.BodyH, m.Cursor, Menu)
	content := ui.Panel(m.renderPage(r.ContentW, r.BodyH), r.ContentW, ui.Max(1, r.BodyH-2))
	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
	footer := ui.Footer(m.Width, m.StatusHint)
	screen := lipgloss.JoinVertical(lipgloss.Left, header, body, footer)

	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m Model) renderPage(width, height int) string {
	switch m.ActivePage {
	case PageDashboard:
		return pages.Dashboard(m.Profile, width, height)
	case PageLanguages:
		return pages.Languages(m.Profile, width, height)
	case PageFrameworks:
		return pages.Frameworks(m.Profile, width, height)
	case PageTools:
		return pages.Tools(m.Profile, width, height)
	case PageDatabases:
		return pages.Databases(m.Profile, width, height)
	case PageGitHub:
		return pages.GitHub(m.Profile, width, height)
	case PageSystem:
		return pages.System(m.Profile, width, height)
	default:
		return pages.Dashboard(m.Profile, width, height)
	}
}
