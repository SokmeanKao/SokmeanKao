package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func Languages(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("LANGUAGES"))
	b.WriteString("\n\n")
	b.WriteString("LANGUAGE            STATUS\n")
	b.WriteString("────────────────────────────\n\n")
	for _, lang := range p.Languages {
		b.WriteString(fmt.Sprintf("%-20s● %s\n", lang.Name, lang.Status))
	}
	return b.String()
}
