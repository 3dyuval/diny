# 🦕 diny GUI recipes

GUI and IDE integration recipes for using `diny message` with visual tools.

## Code Editors & IDEs

### VS Code
```json
// tasks.json
{
  "label": "Generate Commit Message",
  "type": "shell",
  "command": "diny message",
  "presentation": {
    "reveal": "always"
  }
}

// keybindings.json - Add keyboard shortcut
{
  "key": "cmd+shift+m",
  "command": "workbench.action.terminal.sendSequence",
  "args": {
    "text": "diny message | pbcopy\n"
  }
}

// settings.json - Custom Git commit template
{
  "git.inputValidation": "off",
  "terminal.integrated.commandsToSkipShell": ["diny.message"]
}
```

### VS Code Extension Ideas
```javascript
// Would need a VS Code extension that:
// 1. Runs 'diny message' when commit input opens
// 2. Pre-fills the SCM input box
// 3. Could be triggered via Command Palette
```

### JetBrains IDEs (IntelliJ, WebStorm, etc.)

#### External Tool
```xml
<!-- Settings > Tools > External Tools -->
<tool name="Diny Message">
  <exec>diny</exec>
  <params>message</params>
  <output>clipboard</output>
</tool>
```

#### Git Hook Integration
```bash
# .git/hooks/prepare-commit-msg
#!/bin/sh
# Auto-fill for JetBrains IDEs
if [ -z "$2" ]; then
  diny message > "$1"
fi
```

### Sublime Text
```json
// Package: Terminal
{
  "keys": ["cmd+shift+d"],
  "command": "terminal_send_text",
  "args": {
    "text": "diny message | pbcopy\n"
  }
}
```

## macOS & Apple Tools

### Xcode
```bash
# Use with Xcode's Source Control
# 1. Stage changes in Xcode
# 2. Run in Terminal:
diny message | pbcopy
# 3. Paste in Xcode's commit dialog (Cmd+V)

# Or create an Xcode Behavior:
# Preferences > Behaviors > Custom
# Run script: diny message | pbcopy
```

### Tower (Git GUI)
```bash
# Custom Command in Tower
# Settings > Integration > Custom Commands
# Name: Generate Commit Message
# Script: diny message | pbcopy
# Then use Cmd+V in commit dialog
```

### Fork (Git client)
```bash
# Custom Action in Fork
# Preferences > Custom Commands
# Add: diny message > .git/COMMIT_EDITMSG
# Fork will pick up the message file
```

### Sourcetree
```bash
# Custom Action setup:
# Preferences > Custom Actions
# Script: /usr/local/bin/diny message | pbcopy
# Hotkey: Cmd+Shift+D
# Then paste in commit dialog
```

## Cross-Platform Git GUIs

### GitKraken
```bash
# Use terminal to generate and copy
diny message | pbcopy
# Paste in GitKraken's commit message field

# Or set up a Git Hook (works with GitKraken)
echo '#!/bin/sh\ndiny message > "$1"' > .git/hooks/prepare-commit-msg
chmod +x .git/hooks/prepare-commit-msg
```

### GitHub Desktop
```bash
# Generate and copy to clipboard
diny message | pbcopy
# Paste in GitHub Desktop's commit field

# Or use repository Git hooks
# GitHub Desktop respects prepare-commit-msg hooks
```

### GitExtensions (Windows)
```bash
# Custom Commands in GitExtensions
# Commands > Custom Commands
# Add command: diny message | clip
# Hotkey: Ctrl+Shift+D
# Then paste in commit dialog
```

### SmartGit
```bash
# External Command setup:
# Preferences > Commands
# Add: diny message > .git/COMMIT_EDITMSG
# SmartGit monitors this file
```

## Windows-Specific Tools

### TortoiseGit
```bash
# Hook Scripts setup:
# Settings > Hook Scripts
# Pre-Commit Hook: diny message > %1
# Enables auto-fill in commit dialog
```

### Git GUI (Tcl/Tk)
```bash
# Add button to toolbar
# Edit > Options > Add Tool
# Command: diny message | clip
# Then paste in commit message field
```

## Linux Desktop Integration

### GNOME Integration
```bash
# Custom shortcut in GNOME
# Settings > Keyboard > Custom Shortcuts
# Command: gnome-terminal -e "diny message | xclip -selection clipboard"
# Hotkey: Super+Shift+D
```

### KDE Integration
```bash
# KDE Custom Menu entry
# Right-click desktop > Create New > Link to Application
# Command: konsole -e "diny message | xclip -selection clipboard"
```

## Browser-Based Git Tools

### GitHub Web Interface
```bash
# Use browser extension or bookmarklet
# 1. Generate: diny message | pbcopy
# 2. Paste in GitHub's commit message field
# 3. Could be automated with Tampermonkey script
```

### GitLab Web Interface
```bash
# Same approach as GitHub
# Generate message and paste in web commit interface
```

### Bitbucket Web Interface
```bash
# Browser-based workflow
# Terminal: diny message | pbcopy
# Paste in Bitbucket's commit dialog
```

## Integration Strategies by Platform

### Clipboard-Based (Universal)
```bash
# macOS
diny message | pbcopy

# Windows
diny message | clip

# Linux (X11)
diny message | xclip -selection clipboard

# Linux (Wayland)
diny message | wl-copy
```

### File-Based (.git/COMMIT_EDITMSG)
```bash
# Many GUIs monitor this file
diny message > .git/COMMIT_EDITMSG

# Tools that support this:
# - Fork
# - SmartGit
# - Many JetBrains IDEs
# - Some VS Code extensions
```

### Git Hook-Based (prepare-commit-msg)
```bash
#!/bin/sh
# .git/hooks/prepare-commit-msg
# Auto-fills for GUI tools that respect hooks
if [ -z "$2" ]; then
  diny message > "$1"
fi
chmod +x .git/hooks/prepare-commit-msg

# Tools that support this:
# - GitHub Desktop
# - GitKraken
# - JetBrains IDEs
# - Many others
```

## Mobile & Remote Integration

### iPad/iOS (Working Copy app)
```bash
# Use SSH or shortcuts to run on remote server
# Generate message remotely and copy to Working Copy
```

### VS Code Remote Development
```bash
# Works seamlessly with remote containers/SSH
# diny runs on remote machine, output copied locally
```

### Gitpod/Codespaces
```bash
# Cloud development environments
# Install diny in container: brew install dinoDanic/tap/diny
# Use same recipes as local development
```