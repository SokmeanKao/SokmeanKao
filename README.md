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
