# Sokmean TUI Phase 1 Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Build a real alternate-screen Bubble Tea v2 terminal app in `SokmeanKao` with responsive layout, navigation, clock, and static profile pages — no GitHub HTTP.

**Architecture:** Thin `cmd/sokmean` boots a central `internal/app.Model` that owns state/keys/lifecycle. `internal/ui` owns theme chrome. `internal/pages` are pure string renderers. `internal/profile` holds static data from the former README. No page is a Bubble Tea sub-model in Phase 1.

**Tech Stack:** Go 1.22+, `charm.land/bubbletea/v2`, `charm.land/lipgloss/v2`. No Bubbles v2 required. No `internal/github/`.

**Spec:** `docs/superpowers/specs/2026-09-14-sokmean-tui-phase1-design.md`

## Global Constraints

- Module path: `github.com/SokmeanKao/SokmeanKao`
- Binary entry: `cmd/sokmean/main.go`
- Import Bubble Tea as `tea "charm.land/bubbletea/v2"`; Lip Gloss as `"charm.land/lipgloss/v2"`
- `View() tea.View` with `v.AltScreen = true` (v2 declarative API)
- Keys via `tea.KeyPressMsg` + `msg.String()`
- Theme tokens exact: background `#0D1117`, panel `#111827`, border `#1F6F50`, accent `#00FF9C`, text-primary `#D1D5DB`, text-muted `#6B7280`, danger `#FF5F56`, warning `#FFBD2E`, success `#00FF9C`
- Rule: chrome = green, content = neutral, alerts = semantic
- Usable at ~80×24; every Lip Gloss `Width`/`Height` uses `max(1, …)`
- GitHub page must contain exact substring `GitHub API: NOT CONNECTED`
- Soft stubs for `r`, `/`, `?` → footer `StatusHint` only; no network/modals
- No HTTP, auth, cache, or `internal/github/`
- Replace root README with short run docs only (not a fake terminal)
- Commit after each task; do not push unless asked

---

## File Structure

| Path | Responsibility |
|---|---|
| `go.mod` / `go.sum` | Module + Charm deps |
| `.gitignore` | Ignore binaries (`sokmean`, `sokmean.exe`, `/bin/`) |
| `cmd/sokmean/main.go` | `tea.NewProgram` + Run / exit codes |
| `internal/profile/profile.go` | `Profile` types + `Default()` seed data |
| `internal/profile/profile_test.go` | Seed invariants |
| `internal/ui/styles.go` | Color tokens + shared styles |
| `internal/ui/panel.go` | Bordered panel helper + `Clamp` |
| `internal/ui/layout.go` | Region size calculations |
| `internal/ui/header.go` | Header renderer |
| `internal/ui/sidebar.go` | Sidebar renderer |
| `internal/ui/footer.go` | Footer renderer |
| `internal/ui/ui_test.go` | Layout clamp + style smoke tests |
| `internal/pages/*.go` | Seven pure page renderers |
| `internal/pages/pages_test.go` | Content assertions |
| `internal/app/page.go` | `PageID`, `Menu`, helpers |
| `internal/app/model.go` | `Model`, `New`, `Init`, `TickMsg` |
| `internal/app/keys.go` | Key handling helpers |
| `internal/app/update.go` | `Update` |
| `internal/app/view.go` | `View` composition |
| `internal/app/app_test.go` | Navigation / resize / stub tests |
| `README.md` | How to run + keys + Phase 1 note |

---

### Task 1: Module scaffold

**Files:**
- Create: `go.mod`, `.gitignore`
- Modify: (none yet)
- Test: `go list ./...` after deps (will be empty packages until later — verify module init first)

**Interfaces:**
- Consumes: none
- Produces: module `github.com/SokmeanKao/SokmeanKao` with Bubble Tea v2 + Lip Gloss v2 in `go.mod`

- [ ] **Step 1: Initialize module and ignore binaries**

Create `.gitignore`:

```gitignore
# Binaries
/sokmean
/sokmean.exe
/bin/
*.exe
*.test
*.out

# IDE / OS
.idea/
.vscode/
.DS_Store
Thumbs.db
```

Run from repo root `E:/Me/SokmeanKao`:

```bash
go mod init github.com/SokmeanKao/SokmeanKao
go get charm.land/bubbletea/v2@latest
go get charm.land/lipgloss/v2@latest
go mod tidy
```

Expected: `go.mod` contains both `charm.land/bubbletea/v2` and `charm.land/lipgloss/v2`; `go.sum` exists.

- [ ] **Step 2: Verify Go toolchain resolves imports**

```bash
go list -m charm.land/bubbletea/v2 charm.land/lipgloss/v2
```

Expected: two module lines with versions (not errors).

- [ ] **Step 3: Commit**

```bash
git add go.mod go.sum .gitignore
git commit -m "$(cat <<'EOF'
chore: initialize Go module with Charm v2 deps

EOF
)"
```

---

### Task 2: Static profile package

**Files:**
- Create: `internal/profile/profile.go`, `internal/profile/profile_test.go`
- Test: `internal/profile/profile_test.go`

**Interfaces:**
- Consumes: none
- Produces:
  - `type Skill struct { Name string; Status string; Percent int }`
  - `type NamedStatus struct { Name string; Status string }`
  - `type Profile struct { Name, Role, Location, Status, GitHubUser, Host string; Languages []Skill; Frameworks, Databases []NamedStatus; Tools, Editors, Browsers, Modes []string }`
  - `func Default() Profile`

- [ ] **Step 1: Write failing tests**

```go
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
```

- [ ] **Step 2: Run tests — expect FAIL**

```bash
go test ./internal/profile/ -v
```

Expected: FAIL (package/types undefined).

- [ ] **Step 3: Implement `profile.go`**

```go
package profile

type Skill struct {
	Name     string
	Status   string
	Percent  int
}

type NamedStatus struct {
	Name   string
	Status string
}

type Profile struct {
	Name       string
	Role       string
	Location   string
	Status     string
	GitHubUser string
	Host       string
	Languages  []Skill
	Frameworks []NamedStatus
	Tools      []string
	Databases  []NamedStatus
	Editors    []string
	Browsers   []string
	Modes      []string
}

func Default() Profile {
	return Profile{
		Name:       "Sokmean",
		Role:       "Full Stack Developer",
		Location:   "Cambodia",
		Status:     "ONLINE",
		GitHubUser: "sokmeankao",
		Host:       "MSI Laptop",
		Languages: []Skill{
			{Name: "Java", Status: "Active", Percent: 90},
			{Name: "JavaScript", Status: "Active", Percent: 80},
			{Name: "HTML/CSS", Status: "Active", Percent: 85},
		},
		Frameworks: []NamedStatus{
			{Name: "Spring", Status: "Ready"},
			{Name: "Spring Boot", Status: "Ready"},
			{Name: "Next.js", Status: "Ready"},
			{Name: "Docker", Status: "Ready"},
			{Name: "JWT", Status: "Ready"},
			{Name: "Yarn", Status: "Ready"},
		},
		Tools: []string{
			"GitHub", "Notion", "DigitalOcean", "Google Cloud",
			"Vercel", "Hostinger", "RabbitMQ", "Jenkins",
		},
		Databases: []NamedStatus{
			{Name: "PostgreSQL", Status: "READY"},
			{Name: "MySQL", Status: "READY"},
		},
		Editors:  []string{"IntelliJ IDEA", "VS Code", "CodeSandbox"},
		Browsers: []string{"Brave", "Chrome"},
		Modes:    []string{"Building systems", "Learning", "Shipping software"},
	}
}
```

- [ ] **Step 4: Run tests — expect PASS**

```bash
go test ./internal/profile/ -v
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add internal/profile/
git commit -m "$(cat <<'EOF'
feat: add static profile seed from former README

EOF
)"
```

---

### Task 3: UI tokens, clamp, panel, layout math

**Files:**
- Create: `internal/ui/styles.go`, `internal/ui/panel.go`, `internal/ui/layout.go`, `internal/ui/ui_test.go`
- Test: `internal/ui/ui_test.go`

**Interfaces:**
- Consumes: none
- Produces:
  - Colors/styles in `styles.go` (`Accent`, `TitleStyle`, `SelectedStyle`, `NormalStyle`, `DimStyle`, `PanelStyle`)
  - `func Clamp(v, min, max int) int`
  - `func Max(a, b int) int`
  - `func Panel(content string, width, height int) string`
  - `type Regions struct { HeaderH, FooterH, BodyH, SidebarW, ContentW int }`
  - `func ComputeRegions(termW, termH int) Regions`

- [ ] **Step 1: Write failing layout tests**

```go
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
```

- [ ] **Step 2: Run tests — expect FAIL**

```bash
go test ./internal/ui/ -v
```

Expected: FAIL (undefined).

- [ ] **Step 3: Implement UI helpers**

`internal/ui/styles.go`:

```go
package ui

import "charm.land/lipgloss/v2"

var (
	ColorBackground  = lipgloss.Color("#0D1117")
	ColorPanel       = lipgloss.Color("#111827")
	ColorBorder      = lipgloss.Color("#1F6F50")
	ColorAccent      = lipgloss.Color("#00FF9C")
	ColorTextPrimary = lipgloss.Color("#D1D5DB")
	ColorTextMuted   = lipgloss.Color("#6B7280")
	ColorDanger      = lipgloss.Color("#FF5F56")
	ColorWarning     = lipgloss.Color("#FFBD2E")
	ColorSuccess     = lipgloss.Color("#00FF9C")
)

var (
	TitleStyle = lipgloss.NewStyle().Foreground(ColorAccent).Bold(true)
	DimStyle   = lipgloss.NewStyle().Foreground(ColorTextMuted)
	NormalStyle = lipgloss.NewStyle().Foreground(ColorTextPrimary)
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#000000")).
			Background(ColorAccent).
			Bold(true)
	PanelStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorBorder).
			Background(ColorPanel)
)
```

`internal/ui/panel.go`:

```go
package ui

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Clamp(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

func Panel(content string, width, height int) string {
	return PanelStyle.
		Width(Max(1, width)).
		Height(Max(1, height)).
		Render(content)
}
```

`internal/ui/layout.go`:

```go
package ui

type Regions struct {
	HeaderH  int
	FooterH  int
	BodyH    int
	SidebarW int
	ContentW int
}

// ComputeRegions derives frame sizes for a terminal.
// Header/footer content targets ~3 rows including border; body min 10 when possible.
func ComputeRegions(termW, termH int) Regions {
	headerH := 3
	footerH := 3
	bodyH := Max(1, termH-headerH-footerH)
	if termH >= 16 {
		bodyH = Max(10, bodyH)
		// Recompute if Max(10) exceeded terminal — fall back to remaining.
		if headerH+footerH+bodyH > termH {
			bodyH = Max(1, termH-headerH-footerH)
		}
	}

	sidebarW := 22
	if termW < 60 {
		sidebarW = Max(12, termW/3)
	}
	contentW := Max(1, termW-sidebarW)
	// Leave slight slack for join; keep both >=1.
	if sidebarW+contentW > termW {
		contentW = Max(1, termW-sidebarW)
	}

	return Regions{
		HeaderH:  headerH,
		FooterH:  footerH,
		BodyH:    bodyH,
		SidebarW: sidebarW,
		ContentW: contentW,
	}
}
```

- [ ] **Step 4: Run tests — expect PASS**

```bash
go test ./internal/ui/ -v
```

Expected: PASS. If Lip Gloss border height accounting causes `TestPanel_NonEmpty` only issues, keep assertion as non-empty.

- [ ] **Step 5: Commit**

```bash
git add internal/ui/
git commit -m "$(cat <<'EOF'
feat: add UI theme tokens, panel helper, and region math

EOF
)"
```

---

### Task 4: Header, sidebar, footer chrome

**Files:**
- Create: `internal/ui/header.go`, `internal/ui/sidebar.go`, `internal/ui/footer.go`
- Modify: `internal/ui/ui_test.go` (add chrome tests)
- Test: `internal/ui/ui_test.go`

**Interfaces:**
- Consumes: `Panel`, styles, `Max`
- Produces:
  - `func Header(termW int, clock string) string`
  - `func Sidebar(width, height, cursor int, items []string) string`
  - `func Footer(termW int, statusHint string) string`

- [ ] **Step 1: Write failing chrome tests**

Add to `ui_test.go`:

```go
func TestHeader_ContainsBrandAndClock(t *testing.T) {
	out := ui.Header(80, "07:31:00")
	if !containsAll(out, "SOKMEAN", "ONLINE", "07:31:00") {
		t.Fatalf("header=%q", out)
	}
}

func TestSidebar_HighlightsCursor(t *testing.T) {
	items := []string{"Dashboard", "Languages"}
	out := ui.Sidebar(22, 12, 1, items)
	if !containsAll(out, "NAVIGATION", "Languages") {
		t.Fatalf("sidebar=%q", out)
	}
}

func TestFooter_ShowsKeysAndHint(t *testing.T) {
	out := ui.Footer(80, "refresh: Phase 2")
	if !containsAll(out, "navigate", "quit", "refresh: Phase 2") {
		t.Fatalf("footer=%q", out)
	}
}

func containsAll(s string, parts ...string) bool {
	for _, p := range parts {
		if !stringContains(s, p) {
			return false
		}
	}
	return true
}

func stringContains(s, sub string) bool {
	return len(sub) == 0 || (len(s) >= len(sub) && (s == sub || len(s) > 0 && containsAt(s, sub)))
}

func containsAt(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
```

Prefer `strings.Contains` from stdlib instead of hand-rolled helpers — use `strings.Contains` in the real test file.

- [ ] **Step 2: Run tests — expect FAIL**

```bash
go test ./internal/ui/ -v
```

Expected: FAIL on new Header/Sidebar/Footer symbols.

- [ ] **Step 3: Implement chrome**

`header.go`:

```go
package ui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

func Header(termW int, clock string) string {
	left := TitleStyle.Render(" SOKMEAN // DEV MONITOR")
	right := NormalStyle.Render(fmt.Sprintf("● ONLINE  %s ", clock))
	inner := Max(1, termW-2)
	space := Max(1, inner-lipgloss.Width(left)-lipgloss.Width(right))
	line := left + strings.Repeat(" ", space) + right
	return Panel(line, Max(1, termW-2), 1)
}
```

`sidebar.go`:

```go
package ui

import "strings"

func Sidebar(width, height, cursor int, items []string) string {
	var b strings.Builder
	b.WriteString(TitleStyle.Render("NAVIGATION"))
	b.WriteString("\n\n")
	for i, item := range items {
		if i == cursor {
			b.WriteString(SelectedStyle.Render(" > " + item + " "))
		} else {
			b.WriteString(NormalStyle.Render("   " + item))
		}
		b.WriteString("\n")
	}
	return Panel(b.String(), width, Max(1, height-2))
}
```

`footer.go`:

```go
package ui

import "strings"

func Footer(termW int, statusHint string) string {
	keys := " ↑↓/jk navigate   enter select   1-7 pages   r/? stubs   q quit "
	line := DimStyle.Render(keys)
	if statusHint != "" {
		line += "\n" + DimStyle.Render(" "+statusHint+" ")
	}
	return Panel(strings.TrimRight(line, "\n"), Max(1, termW-2), Max(1, 1+boolToInt(statusHint != "")))
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}
```

If panel height math is awkward, keep footer height fixed at content lines and let `Panel` clip — prefer readable footer over exact row accounting.

- [ ] **Step 4: Run tests — expect PASS**

```bash
go test ./internal/ui/ -v
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add internal/ui/
git commit -m "$(cat <<'EOF'
feat: add header, sidebar, and footer chrome renderers

EOF
)"
```

---

### Task 5: Pure page renderers

**Files:**
- Create: `internal/pages/dashboard.go`, `languages.go`, `frameworks.go`, `tools.go`, `databases.go`, `github.go`, `system.go`, `pages_test.go`, `bars.go` (skill bar helper)
- Test: `internal/pages/pages_test.go`

**Interfaces:**
- Consumes: `profile.Profile`, `ui.TitleStyle` / `ui.NormalStyle` / `ui.DimStyle` (or plain strings + local titles — prefer `ui` styles for accent titles)
- Produces:
  - `func Dashboard(p profile.Profile, width, height int) string`
  - `func Languages(p profile.Profile, width, height int) string`
  - `func Frameworks(p profile.Profile, width, height int) string`
  - `func Tools(p profile.Profile, width, height int) string`
  - `func Databases(p profile.Profile, width, height int) string`
  - `func GitHub(p profile.Profile, width, height int) string`
  - `func System(p profile.Profile, width, height int) string`
  - `func SkillBar(percent, width int) string`

- [ ] **Step 1: Write failing page tests**

```go
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
```

- [ ] **Step 2: Run tests — expect FAIL**

```bash
go test ./internal/pages/ -v
```

Expected: FAIL.

- [ ] **Step 3: Implement pages**

`bars.go`:

```go
package pages

import "strings"

func SkillBar(percent, width int) string {
	if width < 1 {
		width = 1
	}
	if percent < 0 {
		percent = 0
	}
	if percent > 100 {
		percent = 100
	}
	filled := percent * width / 100
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}
```

`dashboard.go` (pattern for others):

```go
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
```

Implement `Languages`, `Frameworks`, `Tools`, `Databases`, `System` similarly with titled lists from `profile`.

`github.go` **must** include the exact line:

```text
GitHub API: NOT CONNECTED
```

Also list planned metrics: Contributions, Repositories, Pull Requests, Commits, Languages, Recent Activity.

- [ ] **Step 4: Run tests — expect PASS**

```bash
go test ./internal/pages/ -v
```

Expected: PASS.

- [ ] **Step 5: Commit**

```bash
git add internal/pages/
git commit -m "$(cat <<'EOF'
feat: add static profile page renderers

EOF
)"
```

---

### Task 6: App state, keys, Update (no View yet)

**Files:**
- Create: `internal/app/page.go`, `model.go`, `keys.go`, `update.go`, `app_test.go`
- Test: `internal/app/app_test.go`

**Interfaces:**
- Consumes: `profile.Default`, Bubble Tea msgs
- Produces:
  - `type PageID int` + consts `PageDashboard` … `PageSystem`
  - `var Menu []string`
  - `func (id PageID) Valid() bool`
  - `type Model struct { Width, Height, Cursor int; ActivePage PageID; Now time.Time; Profile profile.Profile; StatusHint string }`
  - `func New() Model`
  - `func (m Model) Init() tea.Cmd`
  - `func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd)`
  - `type TickMsg time.Time`
  - `func Tick() tea.Cmd`

- [ ] **Step 1: Write failing Update tests**

```go
package app_test

import (
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/app"
)

func TestNew_Defaults(t *testing.T) {
	m := app.New()
	if m.ActivePage != app.PageDashboard || m.Cursor != 0 {
		t.Fatalf("page=%v cursor=%d", m.ActivePage, m.Cursor)
	}
	if m.Profile.Name != "Sokmean" {
		t.Fatal("profile not seeded")
	}
}

```

**Key construction (Task 6):** After `go get`, inspect the installed `charm.land/bubbletea/v2` `Key` / `KeyPressMsg` types and write `mustKey(t, s string) tea.KeyPressMsg` such that `mustKey(t, s).String() == s` for: `j`, `k`, `up`, `down`, `enter`, `1`–`7`, `q`, `r`, `/`, `?`, `ctrl+c`. Fatal in the helper if `String()` mismatches. Prefer setting `Code` for specials (`tea.KeyEnter`, `tea.KeyUp`, …) and `Text` for printable runes.

Behavioral tests:

```go
func TestUpdate_JKAndEnter(t *testing.T) {
	m := app.New()
	m.Width, m.Height = 80, 24

	mod, _ := m.Update(mustKey(t, "j"))
	m = mod.(app.Model)
	if m.Cursor != 1 || m.ActivePage != app.PageDashboard {
		t.Fatalf("after j: cursor=%d page=%v", m.Cursor, m.ActivePage)
	}

	mod, _ = m.Update(mustKey(t, "enter"))
	m = mod.(app.Model)
	if m.ActivePage != app.PageLanguages {
		t.Fatalf("page=%v", m.ActivePage)
	}
}

func TestUpdate_NumberKeys(t *testing.T) {
	m := app.New()
	mod, _ := m.Update(mustKey(t, "6"))
	m = mod.(app.Model)
	if m.Cursor != 5 || m.ActivePage != app.PageGitHub {
		t.Fatalf("cursor=%d page=%v", m.Cursor, m.ActivePage)
	}
}

func TestUpdate_Quit(t *testing.T) {
	m := app.New()
	_, cmd := m.Update(mustKey(t, "q"))
	if cmd == nil {
		t.Fatal("expected Quit cmd")
	}
}

func TestUpdate_ResizeDoesNotResetPage(t *testing.T) {
	m := app.New()
	m.Cursor = 3
	m.ActivePage = app.PageTools
	mod, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 40})
	m = mod.(app.Model)
	if m.Width != 100 || m.Height != 40 {
		t.Fatalf("size=%dx%d", m.Width, m.Height)
	}
	if m.Cursor != 3 || m.ActivePage != app.PageTools {
		t.Fatal("resize reset navigation")
	}
}

func TestUpdate_SoftStubs(t *testing.T) {
	m := app.New()
	mod, _ := m.Update(mustKey(t, "r"))
	m = mod.(app.Model)
	if m.StatusHint == "" {
		t.Fatal("expected status hint for r")
	}
	mod, _ = m.Update(mustKey(t, "/"))
	m = mod.(app.Model)
	if m.StatusHint == "" {
		t.Fatal("expected status hint for /")
	}
}

func TestUpdate_TickAdvancesNow(t *testing.T) {
	m := app.New()
	ts := time.Date(2026, 9, 14, 7, 31, 0, 0, time.UTC)
	mod, cmd := m.Update(app.TickMsg(ts))
	m = mod.(app.Model)
	if !m.Now.Equal(ts) {
		t.Fatalf("now=%v", m.Now)
	}
	if cmd == nil {
		t.Fatal("expected re-tick cmd")
	}
}
```

- [ ] **Step 2: Run tests — expect FAIL**

```bash
go test ./internal/app/ -v
```

Expected: FAIL.

- [ ] **Step 3: Implement app state + Update**

`page.go`:

```go
package app

type PageID int

const (
	PageDashboard PageID = iota
	PageLanguages
	PageFrameworks
	PageTools
	PageDatabases
	PageGitHub
	PageSystem
)

var Menu = []string{
	"Dashboard",
	"Languages",
	"Frameworks",
	"Tools",
	"Databases",
	"GitHub",
	"System",
}

func (id PageID) Valid() bool {
	return id >= PageDashboard && id <= PageSystem
}
```

`model.go`:

```go
package app

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/profile"
)

type TickMsg time.Time

type Model struct {
	Width      int
	Height     int
	Cursor     int
	ActivePage PageID
	Now        time.Time
	Profile    profile.Profile
	StatusHint string
}

func New() Model {
	return Model{
		ActivePage: PageDashboard,
		Now:        time.Now(),
		Profile:    profile.Default(),
	}
}

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func (m Model) Init() tea.Cmd {
	return Tick()
}
```

`keys.go` + `update.go`: handle `WindowSizeMsg`, `TickMsg`, `KeyPressMsg` cases for navigation, numbers `1`–`7`, quit, and stubs:

| Key | StatusHint |
|---|---|
| `r` | `refresh: Phase 2` |
| `/` | `search: Phase 2` |
| `?` | `help: Phase 2` |

Clamp cursor to `0..len(Menu)-1`. On digit keys, set cursor and `ActivePage`. On Enter, `ActivePage = PageID(Cursor)`.

Return `tea.Quit` for `q` and `ctrl+c`.

- [ ] **Step 4: Run tests — expect PASS**

```bash
go test ./internal/app/ -v
```

Expected: PASS. Fix `mustKey` construction against the real v2 API if needed.

- [ ] **Step 5: Commit**

```bash
git add internal/app/
git commit -m "$(cat <<'EOF'
feat: add app model, navigation, clock tick, and key stubs

EOF
)"
```

---

### Task 7: View composition + main entrypoint

**Files:**
- Create: `internal/app/view.go`, `cmd/sokmean/main.go`
- Modify: `internal/app/app_test.go` (optional View smoke test)
- Test: unit smoke + manual run

**Interfaces:**
- Consumes: `ui.ComputeRegions`, `ui.Header/Sidebar/Footer/Panel`, `pages.*`
- Produces:
  - `func (m Model) View() tea.View`
  - `main` that runs `tea.NewProgram(app.New())`

- [ ] **Step 1: Write View smoke test**

```go
func TestView_InitializingAndFrame(t *testing.T) {
	m := app.New()
	v := m.View()
	// tea.View should be usable; when Width==0 content mentions Initializing
	// After setting size, content should include brand and menu labels.
	m.Width, m.Height = 80, 24
	v = m.View()
	_ = v
	if !v.AltScreen {
		t.Fatal("AltScreen must be true")
	}
}
```

Inspect `tea.View` fields on the installed version — assert `AltScreen == true`. For content, if `View` stores string in a field (e.g. access via rendering helper or exported content field), assert it contains `SOKMEAN` and `NAVIGATION`. If content is unexported, skip string assert and rely on manual run.

- [ ] **Step 2: Run test — expect FAIL**

```bash
go test ./internal/app/ -run TestView -v
```

Expected: FAIL (View undefined).

- [ ] **Step 3: Implement `view.go` and `main.go`**

`view.go` outline:

```go
func (m Model) View() tea.View {
	if m.Width == 0 || m.Height == 0 {
		v := tea.NewView("Initializing...")
		v.AltScreen = true
		return v
	}
	r := ui.ComputeRegions(m.Width, m.Height)
	header := ui.Header(m.Width, m.Now.Format("15:04:05"))
	sidebar := ui.Sidebar(r.SidebarW, r.BodyH, m.Cursor, Menu)
	content := ui.Panel(m.renderPage(r.ContentW, r.BodyH), Max(1, r.ContentW-2 /* adjust */), Max(1, r.BodyH-2))
	// Prefer: pages render raw text; wrap with ui.Panel using ContentW and BodyH consistently with Sidebar.
	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
	footer := ui.Footer(m.Width, m.StatusHint)
	screen := lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
	v := tea.NewView(screen)
	v.AltScreen = true
	return v
}

func (m Model) renderPage(width, height int) string {
	switch m.ActivePage {
	case PageDashboard:
		return pages.Dashboard(m.Profile, width, height)
	case PageLanguages:
		return pages.Languages(m.Profile, width, height)
	case PageFrameworks:
		return pages.Frameworks(m.Profile, width, height)
	case PageTools:
		return pages.Tools(m.Profile, width, height)
	case PageDatabases:
		return pages.Databases(m.Profile, width, height)
	case PageGitHub:
		return pages.GitHub(m.Profile, width, height)
	case PageSystem:
		return pages.System(m.Profile, width, height)
	default:
		return pages.Dashboard(m.Profile, width, height)
	}
}
```

Tune panel width/height so `JoinHorizontal` fits `m.Width` without overflow (use `r.SidebarW` and `r.ContentW` as the widths passed into `ui.Sidebar` / `ui.Panel`).

`cmd/sokmean/main.go`:

```go
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/SokmeanKao/SokmeanKao/internal/app"
)

func main() {
	p := tea.NewProgram(app.New())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "TUI error: %v\n", err)
		os.Exit(1)
	}
}
```

- [ ] **Step 4: Unit tests + build**

```bash
go test ./...
go build -o sokmean.exe ./cmd/sokmean
```

Expected: all tests PASS; binary builds.

- [ ] **Step 5: Manual acceptance (interactive)**

```bash
go run ./cmd/sokmean
```

Checklist (from spec):

1. Alt-screen TUI opens
2. All 7 pages via `j/k`, arrows, Enter, `1–7`
3. Clock ticks without resetting page/cursor
4. Resize repeatedly including ~80×24 — no corruption
5. GitHub shows `GitHub API: NOT CONNECTED`
6. `q` / `ctrl+c` quit cleanly
7. Soft stubs show footer hints for `r` `/` `?`

- [ ] **Step 6: Commit**

```bash
git add internal/app/cmd cmd/sokmean internal/app/view.go
git commit -m "$(cat <<'EOF'
feat: wire TUI view composition and sokmean entrypoint

EOF
)"
```

(Use `git add` paths that match actual files: `internal/app/view.go` `cmd/sokmean/main.go` and any test updates.)

---

### Task 8: README as launcher docs only

**Files:**
- Modify: `README.md` (replace empty/fake-terminal profile with run docs)
- Test: manual read

**Interfaces:**
- Consumes: none
- Produces: short README — install, run, keys, Phase 1 scope note

- [ ] **Step 1: Replace README content**

```markdown
# Sokmean TUI

Interactive developer monitor for [sokmeankao](https://github.com/sokmeankao) — a real terminal UI built with Bubble Tea v2 + Lip Gloss v2.

## Requirements

- Go 1.22+
- A real terminal (Windows Terminal, macOS Terminal, iTerm2, etc.)

## Run

```bash
go run ./cmd/sokmean
```

Or build:

```bash
go build -o sokmean ./cmd/sokmean
./sokmean
```

## Keys

| Key | Action |
|-----|--------|
| `↑` / `k` | Previous item |
| `↓` / `j` | Next item |
| `Enter` | Open page |
| `1`–`7` | Jump to page |
| `r` `/` `?` | Reserved (Phase 2 stubs) |
| `q` | Quit |

## Phase 1

Static profile pages + navigation + resize + clock. Live GitHub API is **not** included yet (`GitHub API: NOT CONNECTED`).
```

- [ ] **Step 2: Commit**

```bash
git add README.md
git commit -m "$(cat <<'EOF'
docs: replace profile README with TUI run instructions

EOF
)"
```

---

## Spec coverage self-review

| Spec requirement | Task |
|---|---|
| Repo becomes Go app / module path / `cmd/sokmean` | 1, 7 |
| Bubble Tea v2 + Lip Gloss v2, alt-screen | 1, 7 |
| Package ownership boundaries | 2–7 |
| No `internal/github/` / no HTTP | all (omitted) |
| Theme tokens + chrome/content rule | 3, 4 |
| Header / sidebar / content / footer | 4, 7 |
| 80×24 + clamp + resize without reset | 3, 6, 7 |
| Clock tick | 6, 7 |
| Navigation keys + number keys + quit | 6 |
| Soft stubs `r` `/` `?` | 6, 8 |
| Static profile from former README | 2, 5 |
| GitHub `NOT CONNECTED` | 5, 7 |
| Short README run docs | 8 |
| Acceptance criteria | 7 manual + automated tests |

## Placeholder / consistency check

- Types aligned: `PageID`, `Menu`, `Model`, `profile.Profile`, `ui.Regions`
- Page function signatures consistent: `(profile.Profile, width, height int) string`
- Exact status string: `GitHub API: NOT CONNECTED`
- Key construction for tests must be verified against the resolved Bubble Tea v2 API during Task 6 (only intentional implementation-time discovery)

---

## Execution handoff

Plan complete and saved to `docs/superpowers/plans/2026-09-14-sokmean-tui-phase1.md`.

**Two execution options:**

1. **Subagent-Driven (recommended)** — fresh subagent per task, review between tasks  
2. **Inline Execution** — execute tasks in this session with checkpoints  

Which approach?
