package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func Dashboard(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("OVERVIEW"))
	b.WriteString("\n\n")
	b.WriteString(fmt.Sprintf("Developer     %s\n", p.Name))
	b.WriteString(fmt.Sprintf("Role          %s\n", p.Role))
	b.WriteString(fmt.Sprintf("Location      %s\n", p.Location))
	b.WriteString(fmt.Sprintf("Status        ● %s\n\n", p.Status))
	b.WriteString("Development Activity\n\n")
	for _, lang := range p.Languages {
		b.WriteString(fmt.Sprintf("%-12s %s  %d%%\n", lang.Name, SkillBar(lang.Percent, 20), lang.Percent))
	}
	b.WriteString("\nCurrent mode\n\n")
	for _, mode := range p.Modes {
		b.WriteString("> " + mode + "\n")
	}
	return b.String()
}
