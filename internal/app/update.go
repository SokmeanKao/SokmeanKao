package app

import (
	"time"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case TickMsg:
		m.Now = time.Time(msg)
		return m, Tick()

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			m.moveCursor(-1)

		case "down", "j":
			m.moveCursor(1)

		case "enter":
			m.openCursorPage()

		case "1":
			m.selectPage(PageDashboard)
		case "2":
			m.selectPage(PageLanguages)
		case "3":
			m.selectPage(PageFrameworks)
		case "4":
			m.selectPage(PageTools)
		case "5":
			m.selectPage(PageDatabases)
		case "6":
			m.selectPage(PageGitHub)
		case "7":
			m.selectPage(PageSystem)

		case "r":
			m.StatusHint = "refresh: Phase 2"
		case "/":
			m.StatusHint = "search: Phase 2"
		case "?":
			m.StatusHint = "help: Phase 2"
		}
	}

	return m, nil
}
