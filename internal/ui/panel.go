package ui

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Panel(content string, width, height int) string {
	return PanelStyle.
		Width(Max(1, width)).
		Height(Max(1, height)).
		Render(content)
}
