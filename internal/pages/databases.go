package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func Databases(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("DATABASES"))
	b.WriteString("\n\n")
	b.WriteString("NAME                STATUS\n")
	b.WriteString("────────────────────────────\n\n")
	for _, db := range p.Databases {
		b.WriteString(fmt.Sprintf("%-20s● %s\n", db.Name, db.Status))
	}
	return b.String()
}
