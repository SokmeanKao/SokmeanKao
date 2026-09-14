package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func Frameworks(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("FRAMEWORKS"))
	b.WriteString("\n\n")
	b.WriteString("FRAMEWORK           STATUS\n")
	b.WriteString("────────────────────────────\n\n")
	for _, fw := range p.Frameworks {
		b.WriteString(fmt.Sprintf("%-20s● %s\n", fw.Name, fw.Status))
	}
	return b.String()
}
