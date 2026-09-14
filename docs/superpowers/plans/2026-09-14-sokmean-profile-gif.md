# Sokmean Profile GIF Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Put an animated GIF of the real Sokmean TUI on the GitHub profile README, generated from a VHS tape that drives the compiled binary.

**Architecture:** Keep the Phase 1 Go TUI unchanged. Add `demo.tape` as the recording source of truth, commit `assets/sokmean-tui.gif` as a UI release artifact, slim `README.md` to GIF + one run line, and document regen + Windows fallback in `docs/profile-gif.md`.

**Tech Stack:** Existing Bubble Tea TUI; Charmbracelet VHS + ttyd + ffmpeg for GIF generation.

**Spec:** `docs/superpowers/specs/2026-09-14-sokmean-profile-gif-design.md`

## Global Constraints

- Profile README must be exactly the centered GIF + `` `go run ./cmd/sokmean` · run the real TUI `` (no key legend / badges / stack tables)
- Output path: `assets/sokmean-tui.gif`
- Record via compiled binary: `go build -o ./bin/sokmean ./cmd/sokmean` then `vhs demo.tape`
- Tape launches `./bin/sokmean` with real keys (`j`/`k`, `Enter`, `1`–`7`); no `--demo` mode; no hand-recorded official GIF
- Tour ~10–15s; page holds 1–2s; end on Dashboard for clean loop
- Terminal feel ~120×34 (tune VHS Width/Height/FontSize)
- GIF is a release artifact — regenerate only on material visual changes
- Windows fallback order: Native VHS → WSL2 → Linux/macOS; never change TUI to fix recording
- `/bin/` gitignored; GIF committed
- Do not push unless asked

---

## File Structure

| Path | Responsibility |
|---|---|
| `.gitignore` | Add `/bin/` |
| `demo.tape` | VHS tour: launch binary, cycle pages, return Dashboard |
| `assets/sokmean-tui.gif` | Committed profile animation |
| `docs/profile-gif.md` | Regen instructions + fallback + when-to-regen |
| `README.md` | Profile face: GIF + one run line |

---

### Task 1: Ignore `bin/` build output

**Files:**
- Modify: `.gitignore`
- Test: `git check-ignore -v bin/sokmean` (after creating path)

**Interfaces:**
- Consumes: existing `.gitignore`
- Produces: `/bin/` ignored so recording binaries are never committed

- [ ] **Step 1: Update `.gitignore`**

Ensure `.gitignore` contains (keep existing binary rules; add `/bin/`):

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

- [ ] **Step 2: Verify ignore**

```bash
mkdir -p bin
touch bin/sokmean
git check-ignore -v bin/sokmean
rm -f bin/sokmean
```

Expected: a line showing `.gitignore` matches `bin/sokmean`.

- [ ] **Step 3: Commit**

```bash
git add .gitignore
git commit -m "$(cat <<'EOF'
chore: gitignore bin/ for VHS recording builds

EOF
)"
```

---

### Task 2: Author `demo.tape`

**Files:**
- Create: `demo.tape`
- Test: string checks via shell (no Go unit test required)

**Interfaces:**
- Consumes: `./bin/sokmean` must exist before `vhs` runs (Task 4)
- Produces: deterministic tour script writing `assets/sokmean-tui.gif`

- [ ] **Step 1: Write failing tape presence check**

```bash
test -f demo.tape && echo "exists" || echo "missing"
```

Expected: `missing` (before creating the file).

- [ ] **Step 2: Create `demo.tape`**

Create `demo.tape` with this exact structure (adjust FontSize by ±2 only if columns look wrong after first render):

```tape
# Sokmean TUI profile demo — source of truth for assets/sokmean-tui.gif
# Regenerate: go build -o ./bin/sokmean ./cmd/sokmean && vhs demo.tape

Output assets/sokmean-tui.gif

Set Shell "bash"
Set FontSize 16
Set Width 1200
Set Height 680
Set Padding 12
Set Framerate 24
Set TypingSpeed 0
Set Theme "Catppuccin Mocha"

# Launch real binary (build first; on Windows use ./bin/sokmean.exe if needed)
Type "./bin/sokmean"
Enter
Sleep 2s

# Languages
Type "2"
Sleep 1500ms

# Frameworks
Type "3"
Sleep 1500ms

# Tools
Type "4"
Sleep 1500ms

# Databases
Type "5"
Sleep 1500ms

# GitHub (shows NOT CONNECTED)
Type "6"
Sleep 2s

# System
Type "7"
Sleep 1500ms

# Return to Dashboard for clean loop
Type "1"
Sleep 2s
```

Notes for the implementer:

- Prefer digit jumps (`2`–`7`, then `1`) so each page is opened without relying on Enter after cursor moves — matches Phase 1 number-key behavior.
- Total sleep budget ≈ 2+1.5×5+2+1.5+2 ≈ 14.5s → within 10–15s target with typing overhead.
- If VHS rejects `Set Shell "bash"` on a host, remove that line and document the host default shell in `docs/profile-gif.md`.
- If Windows requires `.exe`, change the `Type` line to `"./bin/sokmean.exe"` and document both in `docs/profile-gif.md` (do not add a second tape).

- [ ] **Step 3: Verify tape invariants**

```bash
grep -q 'Output assets/sokmean-tui.gif' demo.tape
grep -q '\./bin/sokmean' demo.tape
grep -E -q 'Type "2"|Type "3"|Type "4"|Type "5"|Type "6"|Type "7"|Type "1"' demo.tape
```

Expected: all greps exit 0.

- [ ] **Step 4: Commit**

```bash
git add demo.tape
git commit -m "$(cat <<'EOF'
feat: add VHS tape for profile TUI demo GIF

EOF
)"
```

---

### Task 3: Regen docs + profile README

**Files:**
- Create: `docs/profile-gif.md`
- Modify: `README.md` (replace current run-docs face with GIF hero)
- Test: content assertions via shell/`rg`

**Interfaces:**
- Consumes: GIF path `./assets/sokmean-tui.gif` (file may land in Task 4)
- Produces: profile face + regen documentation

- [ ] **Step 1: Write `docs/profile-gif.md`**

```markdown
# Profile GIF regeneration

The GitHub profile README shows `assets/sokmean-tui.gif`, recorded from the real TUI with [VHS](https://github.com/charmbracelet/vhs).

The GIF is a **UI release artifact**. Regenerate it only when visual behavior changes materially (layout, navigation, theme, page content, or `demo.tape` tour) — not for every code change.

## Prerequisites

- [VHS](https://github.com/charmbracelet/vhs)
- [ttyd](https://github.com/tsl0922/ttyd)
- [ffmpeg](https://ffmpeg.org)

Install examples:

```bash
# Windows (Scoop)
scoop install vhs ffmpeg
# ttyd may need a Windows-compatible build; see fallbacks below

# macOS
brew install vhs

# Go (any OS)
go install github.com/charmbracelet/vhs@latest
```

## Regenerate

From the repo root:

```bash
go build -o ./bin/sokmean ./cmd/sokmean
# Windows if needed:
# go build -o ./bin/sokmean.exe ./cmd/sokmean
# and set demo.tape Type line to "./bin/sokmean.exe"

vhs demo.tape
```

This overwrites `assets/sokmean-tui.gif`.

## Windows fallback order

If native Windows VHS fails (blank frames, freeze, ttyd/ConPTY errors):

1. Native Windows VHS (retry with working ttyd/ffmpeg on PATH)
2. WSL2 — build and run `vhs demo.tape` inside WSL
3. Linux or macOS

Do **not** change the TUI solely to make recording work. Fix the recording environment instead.

## Verify

- Open `assets/sokmean-tui.gif` and confirm page transitions and chrome render
- Confirm README still references `./assets/sokmean-tui.gif`
- `go run ./cmd/sokmean` still launches the interactive app
```

- [ ] **Step 2: Replace `README.md`**

Replace the entire file with:

```markdown
<div align="center">

<img src="./assets/sokmean-tui.gif" width="100%" alt="Sokmean TUI Demo" />

<br>

`go run ./cmd/sokmean` · run the real TUI

</div>
```

- [ ] **Step 3: Verify README shape**

```bash
rg -n "sokmean-tui.gif|go run ./cmd/sokmean|run the real TUI" README.md
rg -n "Keys|Phase 1|Requirements" README.md && exit 1 || true
```

Expected: first `rg` matches; second finds no leftover run-docs sections (command may exit 0 via `|| true` when no matches).

- [ ] **Step 4: Commit**

```bash
git add docs/profile-gif.md README.md
git commit -m "$(cat <<'EOF'
docs: profile GIF README face and regeneration guide

EOF
)"
```

---

### Task 4: Build binary and generate the GIF

**Files:**
- Create: `assets/sokmean-tui.gif` (via VHS)
- Test: file exists, non-trivial size, visual spot-check

**Interfaces:**
- Consumes: `demo.tape`, compilable `./cmd/sokmean`
- Produces: committed GIF artifact

- [ ] **Step 1: Check tooling**

```bash
command -v vhs && vhs --version
command -v ffmpeg && ffmpeg -version | head -1
command -v ttyd && ttyd --version
```

Expected: all three present. If any missing, install per `docs/profile-gif.md`. If native Windows cannot run VHS successfully, switch to WSL2/Linux **before** changing any Go code.

- [ ] **Step 2: Build binary**

```bash
mkdir -p bin assets
go build -o ./bin/sokmean ./cmd/sokmean
# If Windows produces only .exe, use:
# go build -o ./bin/sokmean.exe ./cmd/sokmean
ls -la bin/
```

Expected: executable exists under `bin/`.

- [ ] **Step 3: Record GIF**

```bash
vhs demo.tape
ls -la assets/sokmean-tui.gif
```

Expected: `assets/sokmean-tui.gif` exists and is larger than ~50KB (blank/corrupt renders are often tiny or zero). If VHS fails on Windows, stop and regenerate under WSL2; do not invent a hand-recorded stand-in.

- [ ] **Step 4: Acceptance checks**

```bash
go test ./...
# Confirm README still points at the asset
rg -n "assets/sokmean-tui.gif" README.md
# Confirm GIF is not ignored
git check-ignore -v assets/sokmean-tui.gif || echo "GIF is trackable (good)"
```

Expected: tests PASS; README references the path; `git check-ignore` does **not** report the GIF as ignored (the `|| echo` path runs).

Manual: open the GIF and confirm Dashboard→…→System→Dashboard chrome/pages are visible.

- [ ] **Step 5: Commit GIF**

```bash
git add assets/sokmean-tui.gif
git commit -m "$(cat <<'EOF'
feat: add recorded Sokmean TUI profile GIF

EOF
)"
```

---

## Spec coverage self-review

| Spec requirement | Task |
|---|---|
| README GIF + one run line | 3 |
| `demo.tape` → `assets/sokmean-tui.gif` | 2, 4 |
| Compiled `./bin/sokmean` + `vhs demo.tape` | 4 |
| Tour pages + return Dashboard | 2 |
| `docs/profile-gif.md` + Windows fallback | 3 |
| `/bin/` gitignored; GIF committed | 1, 4 |
| No TUI `--demo` / no hand GIF | all (omitted) |
| `go run ./cmd/sokmean` still works | 4 acceptance |

## Placeholder / consistency check

- Paths consistent: `./bin/sokmean`, `assets/sokmean-tui.gif`, `demo.tape`
- Windows `.exe` handled as documented tape/doc note, not a second pipeline
- No CI, no SVG, no README key legend

---

## Execution handoff

Plan complete and saved to `docs/superpowers/plans/2026-09-14-sokmean-profile-gif.md`.

**Two execution options:**

1. **Subagent-Driven (recommended)** — fresh subagent per task, review between tasks  
2. **Inline Execution** — execute tasks in this session with checkpoints  

Which approach?
