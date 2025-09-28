# 🦕 diny CLI recipes

Command-line integration recipes for using `diny message` with terminal tools.

## Git Aliases

### Direct commit
```bash
# Setup: Add this alias to git
git config --global alias.dc '!diny message | git commit -F -'

# Usage: Just run
git dc

# What it does:
# 1. Generates a commit message with diny
# 2. Pipes it directly to git commit
# The '-F -' tells git to read the message from stdin (the pipe)
```

### Edit before commit
```bash
# Setup: Add this alias
git config --global alias.dce '!diny message > .git/COMMIT_EDITMSG && git commit -e -F .git/COMMIT_EDITMSG'

# Usage:
git dce

# What it does:
# 1. Generates message and saves to .git/COMMIT_EDITMSG
# 2. Opens your editor to modify the message
# 3. Commits with the edited message when you save
```

### Save as draft
```bash
# Setup: Add this alias
git config --global alias.draft '!diny message > .git/COMMIT_EDITMSG && echo "Draft saved. Run: git commit -F .git/COMMIT_EDITMSG"'

# Usage:
git draft

# What it does:
# 1. Generates message and saves it as a draft
# 2. You can commit later with: git commit -F .git/COMMIT_EDITMSG
# Perfect for when you want to review before committing
```

## Shell Functions

### Commit with preview
```bash
# Add to ~/.bashrc or ~/.zshrc
dcommit() {
  msg=$(diny message)
  echo "Generated message:"
  echo "$msg"
  echo
  read -p "Commit? (y/n): " -n 1 -r
  echo
  if [[ $REPLY =~ ^[Yy]$ ]]; then
    git commit -m "$msg"  # Simple: use -m with the message
  fi
}

# Usage: Just type
dcommit
```

### Copy to clipboard
```bash
# macOS - copy message to clipboard for pasting anywhere
diny message | pbcopy

# Linux (X11) - same for Linux with X11
diny message | xclip -selection clipboard

# Linux (Wayland) - for modern Linux desktops
diny message | wl-copy

# Then paste (Cmd+V or Ctrl+V) into any git GUI or terminal
```

## Terminal-Based Editors

### Vim/Neovim
```vim
" Generate and insert at cursor
:r !diny message

" Generate and replace buffer
:%!diny message

" Map to keybinding
nnoremap <leader>gm :r !diny message<CR>

" Custom User Command for commit workflow
" Add to your init.vim or init.lua (as vim.cmd)
command! DinyCommit execute '!git add -A && diny message | git commit -F -'
command! DinyMessage r !diny message
command! DinyEdit execute '!diny message > .git/COMMIT_EDITMSG' | edit .git/COMMIT_EDITMSG | setlocal filetype=gitcommit

" Usage:
" :DinyCommit  - Generate and commit immediately
" :DinyMessage - Insert generated message at cursor
" :DinyEdit    - Generate message and open in buffer for editing

" More advanced: Create buffer with commit message
function! DinyCommitBuffer()
  " Create new buffer
  new
  " Set it as a git commit buffer
  setlocal filetype=gitcommit
  setlocal buftype=nofile
  " Get the message and insert it
  execute 'r !diny message'
  " Delete the empty first line
  normal! ggdd
  " Save and commit on :wq
  autocmd BufWritePost <buffer> execute '!git commit -F' expand('%:p') | bdelete!
endfunction
command! DinyBuffer call DinyCommitBuffer()
```

### Emacs
```elisp
(defun diny-insert-message ()
  "Insert diny commit message at point"
  (interactive)
  (insert (shell-command-to-string "diny message")))
```

## Terminal Multiplexers

### tmux
```bash
# Send to new pane
tmux split-window -h "diny message | cat; read"

# Send to current pane
tmux send-keys "$(diny message)"
```

### Zellij
```bash
zellij action new-pane -- sh -c "diny message; read"
```

## Terminal Git GUIs

### lazygit custom command
```yaml
# ~/.config/lazygit/config.yml
customCommands:
  - key: 'd'
    command: "diny message | git commit -F -"
    context: 'files'
    description: 'Commit with diny'
```

### tig
```bash
# ~/.tigrc
bind status D !diny message | git commit -F -
bind status d !sh -c "diny message > .git/COMMIT_EDITMSG && vim .git/COMMIT_EDITMSG"
```

### GitUI
```bash
# Map in key config
diny message | xclip -selection clipboard  # Then Ctrl+V in GitUI
```

## Shell Workflows

### Batch commits
```bash
# Commit changes in multiple folders
for dir in */; do
  cd "$dir"
  git add -A
  msg=$(diny message)
  git commit -m "$msg"
  cd ..
done
```

### Conditional commit
```bash
# Only commit if there are changes
if [ -n "$(git status --porcelain)" ]; then
  git add -A
  msg=$(diny message)
  git commit -m "$msg"
fi
```

### With conventional commits
```bash
# Add a prefix like "feat:" to the message
msg=$(diny message)
git commit -m "feat: $msg"

# Or for a fix:
git commit -m "fix: $(diny message)"
```

## Advanced Shell Tricks

### Add emoji prefix
```bash
# Add sparkles emoji to your commit
msg=$(diny message)
git commit -m "✨ $msg"
```

### Add ticket number
```bash
# Prepend your JIRA/GitHub issue number
TICKET="JIRA-123"
msg=$(diny message)
git commit -m "$TICKET: $msg"
```

### Log and commit
```bash
# Save the message to a log file AND commit
msg=$(diny message)
echo "$msg" >> commit.log
git commit -m "$msg"
```

### Multi-line with details
```bash
# Create a detailed commit with extra info
msg=$(diny message)
git commit -m "$msg

Reviewed-by: @teammate
Closes: #42"
```

## CI/CD Pipelines

### GitHub Actions
```yaml
- name: Auto-commit changes
  run: |
    git add .
    diny message | git commit -F - || echo "No changes"
    git push
```

### Pre-commit hook
```bash
#!/bin/sh
# .git/hooks/prepare-commit-msg
if [ -z "$2" ]; then
  diny message > "$1"
fi
```

## Script Integration

### Python
```python
import subprocess

msg = subprocess.check_output(["diny", "message"]).decode().strip()
subprocess.run(["git", "commit", "-m", msg])
```

### Node.js
```javascript
const { execSync } = require('child_process');

const msg = execSync('diny message').toString().trim();
execSync(`git commit -m "${msg}"`);
```

### Make targets
```makefile
commit:
	@git add -A
	@diny message | git commit -F -

draft:
	@diny message > .git/COMMIT_EDITMSG
	@echo "Draft saved"
```