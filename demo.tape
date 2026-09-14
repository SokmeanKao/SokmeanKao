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
