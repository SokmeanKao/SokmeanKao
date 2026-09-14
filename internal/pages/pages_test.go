package pages_test

import (
	"strings"
	"testing"

	"github.com/SokmeanKao/SokmeanKao/internal/pages"
	"github.com/SokmeanKao/SokmeanKao/internal/profile"
)

func TestGitHub_NotConnected(t *testing.T) {
	out := pages.GitHub(profile.Default(), 60, 20)
	if !strings.Contains(out, "GitHub API: NOT CONNECTED") {
		t.Fatalf("missing NOT CONNECTED: %q", out)
	}
	if !strings.Contains(out, "sokmeankao") {
		t.Fatal("missing user")
	}
}

func TestDashboard_ShowsOverview(t *testing.T) {
	out := pages.Dashboard(profile.Default(), 60, 20)
	for _, part := range []string{"OVERVIEW", "Sokmean", "Cambodia", "Java", "90%"} {
		if !strings.Contains(out, part) {
			t.Fatalf("missing %q in %q", part, out)
		}
	}
}

func TestSkillBar_Length(t *testing.T) {
	bar := pages.SkillBar(80, 20)
	if strings.Count(bar, "█")+strings.Count(bar, "░") != 20 {
		t.Fatalf("bar=%q", bar)
	}
}

func TestAllPages_NonEmpty(t *testing.T) {
	p := profile.Default()
	fns := []func(profile.Profile, int, int) string{
		pages.Dashboard, pages.Languages, pages.Frameworks,
		pages.Tools, pages.Databases, pages.GitHub, pages.System,
	}
	for i, fn := range fns {
		if fn(p, 50, 15) == "" {
			t.Fatalf("page %d empty", i)
		}
	}
}
