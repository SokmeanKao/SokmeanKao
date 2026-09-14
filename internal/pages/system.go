package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func System(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("SYSTEM"))
	b.WriteString("\n\n")
	b.WriteString(fmt.Sprintf("Host          %s\n", p.Host))
	for _, editor := range p.Editors {
		b.WriteString(fmt.Sprintf("Editor        %s\n", editor))
	}
	for _, browser := range p.Browsers {
		b.WriteString(fmt.Sprintf("Browser       %s\n", browser))
	}
	b.WriteString("\n")
	b.WriteString("TUI Engine    Bubble Tea v2\n")
	b.WriteString("Renderer      Lip Gloss v2\n\n")
	b.WriteString("Status        ● HEALTHY\n")
	return b.String()
}
