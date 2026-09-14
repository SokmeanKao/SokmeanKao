package pages

import (
	"fmt"
	"strings"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func GitHub(p profile.Profile, width, height int) string {
	_ = width
	_ = height
	var b strings.Builder
	b.WriteString(ui.TitleStyle.Render("GITHUB"))
	b.WriteString("\n\n")
	b.WriteString("User\n\n")
	b.WriteString(fmt.Sprintf("    %s\n\n", p.GitHubUser))
	b.WriteString("API\n\n")
	b.WriteString("    GitHub API: NOT CONNECTED\n\n")
	b.WriteString("Planned metrics\n\n")
	for _, metric := range []string{
		"Contributions",
		"Repositories",
		"Pull Requests",
		"Commits",
		"Languages",
		"Recent Activity",
	} {
		b.WriteString("    " + metric + "\n")
	}
	return b.String()
}
