# Profile GIF regeneration

The GitHub profile README shows `assets/sokmean-tui.gif`, recorded from the real TUI with [VHS](https://github.com/charmbracelet/vhs).

The GIF is a **UI release artifact**. Regenerate it only when visual behavior changes materially (layout, navigation, theme, page content, or `demo.tape` tour) — not for every code change.

## Prerequisites

- [VHS](https://github.com/charmbracelet/vhs)
- [ttyd](https://github.com/tsl0922/ttyd)
- [ffmpeg](https://ffmpeg.org)

Install examples:

```bash
# Windows (winget / Scoop)
winget install charmbracelet.vhs
winget install tsl0922.ttyd
winget install Gyan.FFmpeg

# macOS
brew install vhs

# Go (any OS)
go install github.com/charmbracelet/vhs@latest
```

### VHS v0.12 note (ffmpeg never runs)

VHS v0.12 cancels the recording context before `ffmpeg` starts, so the CLI can print `Creating …gif` and exit 0 **without writing a file**. If that happens, build a one-line patched VHS that calls `Render(context.Background())` after teardown (see Charmbracelet `evaluator.go`), or use a fixed upstream release when available.

## Regenerate

### Windows (this repo’s current tape)

`demo.tape` is configured for PowerShell + `./bin/sokmean.exe`:

```powershell
go build -o ./bin/sokmean.exe ./cmd/sokmean
vhs demo.tape
```

Do **not** use `Set Shell "bash"` on Windows unless a real Linux distro provides `/bin/bash` (docker-desktop-only WSL will fail).

### Linux / macOS

Adjust the tape temporarily:

- `Set Shell "bash"` (or remove `Set Shell`)
- `Type "./bin/sokmean"`

Then:

```bash
go build -o ./bin/sokmean ./cmd/sokmean
vhs demo.tape
```

This overwrites `assets/sokmean-tui.gif`.

## Windows fallback order

If native Windows VHS fails (blank frames, freeze, ttyd/ConPTY errors):

1. Native Windows VHS (working ttyd + ffmpeg; patched VHS if needed)
2. WSL2 with a real distro (Ubuntu) — not docker-desktop alone
3. Linux or macOS

Do **not** change the TUI solely to make recording work. Fix the recording environment instead.

## Verify

- Open `assets/sokmean-tui.gif` and confirm page transitions and chrome render
- Confirm README still references `./assets/sokmean-tui.gif`
- `go run ./cmd/sokmean` still launches the interactive app
