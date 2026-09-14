package ui_test

import (
	"testing"

	"github.com/SokmeanKao/SokmeanKao/internal/ui"
)

func TestComputeRegions_80x24(t *testing.T) {
	r := ui.ComputeRegions(80, 24)
	if r.HeaderH < 1 || r.FooterH < 1 {
		t.Fatalf("header/footer missing: %+v", r)
	}
	if r.BodyH < 10 {
		t.Fatalf("BodyH=%d want >=10", r.BodyH)
	}
	if r.SidebarW < 10 {
		t.Fatalf("SidebarW=%d", r.SidebarW)
	}
	if r.ContentW < 30 {
		t.Fatalf("ContentW=%d", r.ContentW)
	}
	if r.SidebarW+r.ContentW > 80 {
		t.Fatalf("widths overflow: %+v", r)
	}
}

func TestComputeRegions_TinyTerminalNeverZero(t *testing.T) {
	r := ui.ComputeRegions(40, 12)
	if r.BodyH < 1 || r.SidebarW < 1 || r.ContentW < 1 {
		t.Fatalf("zero region: %+v", r)
	}
}

func TestClamp(t *testing.T) {
	if ui.Clamp(5, 1, 3) != 3 {
		t.Fatal("upper")
	}
	if ui.Clamp(0, 1, 3) != 1 {
		t.Fatal("lower")
	}
	if ui.Clamp(2, 1, 3) != 2 {
		t.Fatal("mid")
	}
}

func TestPanel_NonEmpty(t *testing.T) {
	out := ui.Panel("hello", 20, 5)
	if out == "" {
		t.Fatal("empty panel")
	}
}
