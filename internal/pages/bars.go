package pages

import "strings"

func SkillBar(percent, width int) string {
	if width < 1 {
		width = 1
	}
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	filled := percent * width / 100
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}
