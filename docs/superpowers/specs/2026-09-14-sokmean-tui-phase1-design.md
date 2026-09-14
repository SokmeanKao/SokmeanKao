# Sokmean TUI — Phase 1 Design

**Date:** 2026-09-14  
**Repo:** `E:\Me\SokmeanKao` (GitHub profile repo becomes the Go app)  
**Module:** `github.com/SokmeanKao/SokmeanKao`  
**Binary:** `sokmean` (`cmd/sokmean`)  
**Status:** Approved for implementation planning

## Goal

Ship a real alternate-screen terminal application (Bubble Tea v2 + Lip Gloss v2), not a README that looks like a terminal. Phase 1 validates the TUI foundation only. Live GitHub data is Phase 2.

## Decisions (locked)

| Decision | Choice |
|---|---|
| Location | Inside `SokmeanKao` — README only documents how to run |
| Scope | Phase 1 only (no HTTP / auth / cache / refresh networking) |
| Theme | Terminal green monitor (`chrome = green`, `content = neutral`, `alerts = semantic`) |
| Architecture | Central `app.Model` + pure page renderers; lightweight `PageID` abstraction |
| Page sub-models | Not used in Phase 1. A page becomes a Bubble Tea sub-model only when it needs independent interactive state |

## Non-goals (Phase 1)

- `internal/github/` package
- GitHub HTTP client, API calls, authentication, caching
- Networked `r` refresh
- Search modal (`/`) or help modal (`?`) — keys may soft-stub with footer hints only
- Scrollable viewports / Bubbles tables (unless forced by a layout bug)
- Nested Bubble Tea page models

## Architecture

```text
cmd/sokmean/main.go
        │
        ▼
internal/app          Model, PageID, keys, Init / Update / View
        │
        ├── internal/ui         layout, styles, header, sidebar, footer, panel
        ├── internal/pages      pure renderers (no keys / timers / cmds / HTTP)
        └── internal/profile    static Profile data from former README
```

### Ownership boundaries

| Package | Owns | Must not own |
|---|---|---|
| `cmd/sokmean` | Program bootstrap | UI logic |
| `internal/app` | Application state, keyboard, lifecycle | Lip Gloss chrome details / page copy |
| `internal/ui` | Visual primitives and theme tokens | Page-specific content |
| `internal/pages` | Page-specific presentation | Keys, timers, commands, network |
| `internal/profile` | Static data model | Rendering |

### Lifecycle

```text
Bubble Tea event
      │
      ▼
 app.Update()  → keyboard | resize | clock
      │
      ▼
 app.View()
      │
      ├── ui.Header()
      ├── ui.Sidebar()
      ├── pages.<Active>()
      └── ui.Footer()
      │
      ▼
 Alternate-screen terminal
```

### Central model (Phase 1)

```go
type Model struct {
	Width      int
	Height     int
	Cursor     int
	ActivePage PageID
	Now        time.Time
	Profile    profile.Profile
	StatusHint string // optional footer stub for r / / / ?
}
```

### Page IDs

```go
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
```

Page renderers stay pure:

```go
func Dashboard(p profile.Profile, width, height int) string
```

### Phase 2 extension (documented, not built)

`app.Model` may later hold `GitHub github.State` without converting every page into a nested model. Only GitHub (and similar) becomes stateful when loading / error / selection / scroll appear.

## Theme tokens

```text
background      #0D1117
panel           #111827
border          #1F6F50
accent          #00FF9C
text-primary    #D1D5DB
text-muted      #6B7280
danger          #FF5F56
warning         #FFBD2E
success         #00FF9C
```

Rules:

- Green for active navigation, borders, status, titles, selected states
- Body content stays gray/white
- Semantic colors only for alerts

## Layout & resize

```text
┌─ header (full width) ─────────────────────────────────────────┐
├─ sidebar ─┬─ content ─────────────────────────────────────────┤
├─ footer (full width) ─────────────────────────────────────────┘
```

| Region | Rule |
|---|---|
| Header | ~3 rows: brand left; `● ONLINE` + `HH:MM:SS` right |
| Footer | ~3 rows: key hints (+ optional status stub) |
| Sidebar | ~20 content columns (~22 with border); clamp only if terminal very narrow |
| Content | Remaining width; Lip Gloss width/height clip |
| Body height | `height - header - footer`; minimum body 10 rows |

**Hard requirements:**

- Usable at roughly **80×24**
- Every `Width`/`Height` uses `max(1, …)` — never feed Lip Gloss zeros
- On `WindowSizeMsg`: update dimensions only; do not reset cursor/page
- Each `View` paints one full frame with alt-screen enabled
- Repeated resize + navigation must not corrupt the screen (ghost borders / split lines)

**Clock:** `tea.Tick` every 1s updates `Now` and re-arms the tick; does not reset navigation state.

## Navigation & keys

Menu order: Dashboard, Languages, Frameworks, Tools, Databases, GitHub, System.

| Key | Action |
|---|---|
| `↑` / `k` | Move cursor up (clamp) |
| `↓` / `j` | Move cursor down (clamp) |
| `Enter` | Set `ActivePage` from cursor |
| `1`…`7` | Jump cursor + page |
| `q` / `ctrl+c` | Quit |
| `r` / `/` / `?` | Soft stub — footer hint only (`refresh: Phase 2`, etc.); no networking/modals |

Sidebar: selected row uses accent background + dark text. Number keys keep cursor and page in sync; arrow/`j`/`k` move cursor until Enter.

## Profile data

Static seed from the former cyberpunk README (not live):

- Identity: Sokmean · Full Stack Developer · Cambodia · ONLINE · `sokmeankao`
- Languages: Java 90%, JavaScript 80%, HTML/CSS 85% (Active)
- Frameworks: Spring, Spring Boot, Next.js, Docker, JWT, Yarn
- Tools: GitHub, Notion, DigitalOcean, Google Cloud, Vercel, Hostinger, RabbitMQ, Jenkins
- Databases: PostgreSQL, MySQL
- System: MSI Laptop; IntelliJ IDEA, VS Code, CodeSandbox; Brave, Chrome; Bubble Tea v2 / Lip Gloss v2

### Pages

| Page | Content |
|---|---|
| Dashboard | Overview fields, skill bars, mode lines |
| Languages / Frameworks / Tools / Databases | Titled lists/tables, neutral body |
| GitHub | User, planned metrics, **`GitHub API: NOT CONNECTED`** |
| System | Host/editors/browsers/engine, `● HEALTHY` |

No scroll in Phase 1 — clip to panel height.

## Errors

- Startup failure → stderr + exit 1
- Unknown keys → ignore
- Invalid sizes → clamp; never panic
- No network error paths in Phase 1

## Acceptance criteria

1. `go run ./cmd/sokmean` opens an alternate-screen TUI
2. All 7 pages reachable via `j/k`, arrows, Enter, and `1–7`
3. Clock updates without resetting page/cursor
4. Repeated resize (including ~80×24) keeps layout coherent with no screen corruption
5. GitHub page shows `GitHub API: NOT CONNECTED`
6. `q` / `ctrl+c` quit cleanly
7. Root `README.md` only documents install/run and Phase 1 scope (not a fake terminal)

## Deliverables

```text
SokmeanKao/
├── cmd/sokmean/main.go
├── internal/app/{model,update,view,keys,page}.go
├── internal/ui/{layout,styles,header,sidebar,footer,panel}.go
├── internal/pages/{dashboard,languages,frameworks,tools,databases,github,system}.go
├── internal/profile/profile.go
├── go.mod
├── .gitignore          (Go binaries)
├── README.md           (short: how to run)
└── docs/superpowers/specs/2026-09-14-sokmean-tui-phase1-design.md
```

Stack: Bubble Tea v2 + Lip Gloss v2 (`charm.land/.../v2`). Bubbles v2 not required for Phase 1.

## Phase 2 boundary (next, not now)

```text
Stable TUI shell
      ↓
internal/github/{client,models,commands,errors}.go
      ↓
GitHub page states: loading | connected | error | rate-limited | refreshing
      ↓
r = refresh
```

## Spec self-review notes

- No TBD placeholders in Phase 1 scope
- Architecture, theme, layout, keys, and data are consistent
- Scope is a single implementation plan (Phase 1 only)
- Soft stubs for `r`/`/`/`?` are explicit (footer hints, not features)
