package ui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func Header(termW int, clock string) string {
	left := TitleStyle.Render(" SOKMEAN // DEV MONITOR")
	right := NormalStyle.Render(fmt.Sprintf("● ONLINE  %s ", clock))
	inner := Max(1, termW-2)
	space := Max(1, inner-lipgloss.Width(left)-lipgloss.Width(right))
	line := left + strings.Repeat(" ", space) + right
	return Panel(line, Max(1, termW-2), 1)
}
