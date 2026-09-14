package pages

import (
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func Tools(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("TOOLS"))
	b.WriteString("\n\n")
	for _, tool := range p.Tools {
		b.WriteString(tool + "\n")
	}
	return b.String()
}
