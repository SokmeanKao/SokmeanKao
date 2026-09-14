package ui

func Footer(termW int, statusHint string) string {
	keys := " ↑↓/jk navigate   enter select   1-7 pages   r/? stubs   q quit "
	line := DimStyle.Render(keys)
	height := 1
	if statusHint != "" {
		line += "\n" + DimStyle.Render(" "+statusHint+" ")
		height = 2
	}
	return Panel(line, Max(1, termW-2), height)
}
