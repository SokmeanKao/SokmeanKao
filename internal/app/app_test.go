package app_test

import (
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/app"
)

func mustKey(t *testing.T, s string) tea.KeyPressMsg {
	t.Helper()
	var k tea.Key
	switch s {
	case "up":
		k.Code = tea.KeyUp
	case "down":
		k.Code = tea.KeyDown
	case "enter":
		k.Code = tea.KeyEnter
	case "ctrl+c":
		k.Code = 'c'
		k.Mod = tea.ModCtrl
	default:
		if len(s) == 0 {
			t.Fatal("empty key")
		}
		k.Text = s
		k.Code = rune(s[0])
	}
	msg := tea.KeyPressMsg(k)
	if msg.String() != s {
		t.Fatalf("mustKey(%q) String()=%q", s, msg.String())
	}
	return msg
}

func TestNew_Defaults(t *testing.T) {
	m := app.New()
	if m.ActivePage != app.PageDashboard || m.Cursor != 0 {
		t.Fatalf("page=%v cursor=%d", m.ActivePage, m.Cursor)
	}
	if m.Profile.Name != "Sokmean" {
		t.Fatal("profile not seeded")
	}
}

func TestUpdate_JKAndEnter(t *testing.T) {
	m := app.New()
	m.Width, m.Height = 80, 24

	mod, _ := m.Update(mustKey(t, "j"))
	m = mod.(app.Model)
	if m.Cursor != 1 || m.ActivePage != app.PageDashboard {
		t.Fatalf("after j: cursor=%d page=%v", m.Cursor, m.ActivePage)
	}

	mod, _ = m.Update(mustKey(t, "enter"))
	m = mod.(app.Model)
	if m.ActivePage != app.PageLanguages {
		t.Fatalf("page=%v", m.ActivePage)
	}
}

func TestUpdate_NumberKeys(t *testing.T) {
	m := app.New()
	mod, _ := m.Update(mustKey(t, "6"))
	m = mod.(app.Model)
	if m.Cursor != 5 || m.ActivePage != app.PageGitHub {
		t.Fatalf("cursor=%d page=%v", m.Cursor, m.ActivePage)
	}
}

func TestUpdate_Quit(t *testing.T) {
	m := app.New()
	_, cmd := m.Update(mustKey(t, "q"))
	if cmd == nil {
		t.Fatal("expected Quit cmd")
	}
}

func TestUpdate_ResizeDoesNotResetPage(t *testing.T) {
	m := app.New()
	m.Cursor = 3
	m.ActivePage = app.PageTools
	mod, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	m = mod.(app.Model)
	if m.Width != 100 || m.Height != 40 {
		t.Fatalf("size=%dx%d", m.Width, m.Height)
	}
	if m.Cursor != 3 || m.ActivePage != app.PageTools {
		t.Fatal("resize reset navigation")
	}
}

func TestUpdate_SoftStubs(t *testing.T) {
	m := app.New()
	mod, _ := m.Update(mustKey(t, "r"))
	m = mod.(app.Model)
	if m.StatusHint == "" {
		t.Fatal("expected status hint for r")
	}
	mod, _ = m.Update(mustKey(t, "/"))
	m = mod.(app.Model)
	if m.StatusHint == "" {
		t.Fatal("expected status hint for /")
	}
}

func TestUpdate_TickAdvancesNow(t *testing.T) {
	m := app.New()
	ts := time.Date(2026, 9, 14, 7, 31, 0, 0, time.UTC)
	mod, cmd := m.Update(app.TickMsg(ts))
	m = mod.(app.Model)
	if !m.Now.Equal(ts) {
		t.Fatalf("now=%v", m.Now)
	}
	if cmd == nil {
		t.Fatal("expected re-tick cmd")
	}
}
