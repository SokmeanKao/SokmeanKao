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
