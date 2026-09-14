package ui

import "strings"

func Sidebar(width, height, cursor int, items []string) string {
	var b strings.Builder
	b.WriteString(TitleStyle.Render("NAVIGATION"))
	b.WriteString("\n\n")
	for i, item := range items {
		if i == cursor {
			b.WriteString(SelectedStyle.Render(" > " + item + " "))
		} else {
			b.WriteString(NormalStyle.Render("   " + item))
		}
		b.WriteString("\n")
	}
	return Panel(b.String(), width, Max(1, height-2))
}
