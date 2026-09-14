# Sokmean TUI profile demo — source of truth for assets/sokmean-tui.gif
# Regenerate: see docs/profile-gif.md
#
# Linux/macOS:
#   go build -o ./bin/sokmean ./cmd/sokmean
#   vhs demo.tape
#
# Windows (PowerShell shell + .exe):
#   go build -o ./bin/sokmean.exe ./cmd/sokmean
#   # temporarily set: Set Shell "powershell"
#   # temporarily set: Type "./bin/sokmean.exe"

Output assets/sokmean-tui.gif

Set Shell "powershell"
Set FontSize 16
Set Width 1200
Set Height 680
Set Padding 12
Set Framerate 24
Set TypingSpeed 0
Set Theme "Catppuccin Mocha"

Type "./bin/sokmean.exe"
Enter
Sleep 2s

Type "2"
Sleep 1500ms

Type "3"
Sleep 1500ms

Type "4"
Sleep 1500ms

Type "5"
Sleep 1500ms

Type "6"
Sleep 2s

Type "7"
Sleep 1500ms

Type "1"
Sleep 2s
