package app

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
)

type TickMsg time.Time

type Model struct {
	Width      int
	Height     int
	Cursor     int
	ActivePage PageID
	Now        time.Time
	Profile    profile.Profile
	StatusHint string
}

func New() Model {
	return Model{
		ActivePage: PageDashboard,
		Now:        time.Now(),
		Profile:    profile.Default(),
	}
}

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return Tick()
}
