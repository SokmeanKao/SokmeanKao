package profile_test

import (
	"testing"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
)

func TestDefault_Identity(t *testing.T) {
	p := profile.Default()
	if p.Name != "Sokmean" {
		t.Fatalf("Name=%q", p.Name)
	}
	if p.Role != "Full Stack Developer" {
		t.Fatalf("Role=%q", p.Role)
	}
	if p.Location != "Cambodia" {
		t.Fatalf("Location=%q", p.Location)
	}
	if p.Status != "ONLINE" {
		t.Fatalf("Status=%q", p.Status)
	}
	if p.GitHubUser != "sokmeankao" {
		t.Fatalf("GitHubUser=%q", p.GitHubUser)
	}
	if p.Host != "MSI Laptop" {
		t.Fatalf("Host=%q", p.Host)
	}
}

func TestDefault_LanguagesHavePercents(t *testing.T) {
	p := profile.Default()
	want := map[string]int{"Java": 90, "JavaScript": 80, "HTML/CSS": 85}
	if len(p.Languages) != len(want) {
		t.Fatalf("len=%d", len(p.Languages))
	}
	for _, lang := range p.Languages {
		pct, ok := want[lang.Name]
		if !ok {
			t.Fatalf("unexpected language %q", lang.Name)
		}
		if lang.Percent != pct {
			t.Fatalf("%s percent=%d want=%d", lang.Name, lang.Percent, pct)
		}
		if lang.Status != "Active" {
			t.Fatalf("%s status=%q", lang.Name, lang.Status)
		}
	}
}

func TestDefault_StackCounts(t *testing.T) {
	p := profile.Default()
	if len(p.Frameworks) != 6 {
		t.Fatalf("frameworks=%d", len(p.Frameworks))
	}
	if len(p.Tools) != 8 {
		t.Fatalf("tools=%d", len(p.Tools))
	}
	if len(p.Databases) != 2 {
		t.Fatalf("databases=%d", len(p.Databases))
	}
	if len(p.Editors) < 2 || len(p.Browsers) < 2 {
		t.Fatalf("editors=%d browsers=%d", len(p.Editors), len(p.Browsers))
	}
}
