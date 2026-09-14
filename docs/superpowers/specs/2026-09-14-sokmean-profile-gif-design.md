# Sokmean Profile GIF — Design

**Date:** 2026-09-14  
**Repo:** `E:\Me\SokmeanKao` (GitHub profile + Go TUI)  
**Status:** Approved for implementation planning  
**Depends on:** Phase 1 TUI (`docs/superpowers/specs/2026-09-14-sokmean-tui-phase1-design.md`)

## Goal

Show a **visually running TUI** on the GitHub profile page via an animated GIF recorded from the real Go binary with VHS — not ASCII art, not SVG animation, not a fake terminal README.

Pipeline:

```text
real Go TUI binary
        ↓
VHS tape (demo.tape)
        ↓
real keyboard input
        ↓
assets/sokmean-tui.gif
        ↓
GitHub profile README
```

## Decisions (locked)

| Decision | Choice |
|---|---|
| Profile content | **B** — GIF hero + one short run line |
| Production pipeline | **B** — documented local regen; GIF committed; no CI yet |
| Recording approach | **1** — VHS drives real compiled binary; no `--demo` mode; no hand-recorded GIF |
| GIF role | UI **release artifact** — regenerate only on material visual changes |

## Non-goals

- Animated SVG / CSS fake terminals
- GitHub Actions GIF regeneration (Phase later if needed)
- `--demo` / auto-tour mode inside the TUI
- Hand-captured screen recordings as the official artifact
- Key legend, badges, or stack tables on the profile README
- Changing TUI behavior solely to work around Windows VHS/ttyd issues

## Profile README

```html
<div align="center">

<img src="./assets/sokmean-tui.gif" width="100%" alt="Sokmean TUI Demo" />

<br>

`go run ./cmd/sokmean` · run the real TUI

</div>
```

Controls and full usage remain in docs (and any longer usage notes), not on the profile face.

## Repo layout

```text
SokmeanKao/
├── assets/
│   └── sokmean-tui.gif      # committed
├── demo.tape                # VHS source of truth
├── bin/                     # gitignored; ./bin/sokmean for recording
├── cmd/sokmean/
├── internal/...
├── docs/
│   ├── profile-gif.md       # regen + Windows fallback
│   └── superpowers/...
├── .gitignore               # include /bin/
└── README.md                # GIF + one line
```

## Recording

### Commands

```bash
go build -o ./bin/sokmean ./cmd/sokmean
vhs demo.tape
```

### Constraints

| Item | Value |
|---|---|
| Output | `assets/sokmean-tui.gif` |
| Terminal feel | ~120 columns × ~34 rows (VHS Width/Height/FontSize tuned to match) |
| Duration | ~10–15 seconds |
| Page hold | 1–2 seconds after each page open |
| Binary | `./bin/sokmean` |
| Loop | End on Dashboard so the GIF loops cleanly |

### Tour (tape source of truth)

1. Launch `./bin/sokmean`, settle 1–2s (clock visible)
2. Visit each page with real keys (`j`/`Enter` and/or `1`–`7`): Languages → Frameworks → Tools → Databases → GitHub → System
3. Return to Dashboard (`1` or navigate)
4. Brief pause; end tape (VHS ends the session)

Keys: `j`/`k`, `Enter`, `1`–`7` only for navigation in the tour. Soft stubs (`r`/`/`/`?`) are optional and not required for Phase 1 GIF.

### Windows fallback (documented, not a TUI change)

```text
1. Native Windows VHS
2. WSL2
3. Linux/macOS
```

If native Windows fails due to ttyd/ConPTY, regenerate under WSL/Linux. Do **not** modify the TUI to accommodate recording.

## `docs/profile-gif.md`

Must include:

- Prerequisites: VHS, ttyd, ffmpeg
- Build + `vhs demo.tape`
- Fallback order above
- When to regenerate (layout, navigation, theme, page content, or tour flow changes — not every code change)

## Acceptance

1. Profile README is GIF hero + one run line only (as specified)
2. Committed GIF shows real TUI chrome and page transitions
3. After `go build -o ./bin/sokmean ./cmd/sokmean`, `vhs demo.tape` overwrites `assets/sokmean-tui.gif`
4. `go run ./cmd/sokmean` still launches the interactive TUI
5. `/bin/` is gitignored; GIF is tracked

## Spec self-review notes

- No TBD in Phase scope
- Consistent with Approach 1 + pipeline B + README B
- Single implementation plan: tape + docs + README + first GIF generation
- Explicit Windows fallback without TUI changes
