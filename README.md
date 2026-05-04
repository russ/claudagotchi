# claudagotchi

A cute terminal dashboard that watches your remote [Claude Code](https://claude.com/claude-code) agents over SSH.

Each host gets its own panel; inside, every active Claude session is a little pixel-art creature: a green sparkly blob when it's working, an orange one when it needs your attention, a sleeping blue one when it's idle, an egg when nothing's running.

```
┌──────────────────────────────┐  ┌──────────────────────────────┐
│ host-a    2 sessions         │  │ host-b    1 session          │
│                              │  │                              │
│  ▄▄▄▄    WORKING             │  │  ▄▄▄▄    WAITING             │
│ ▄████▄   repos/api  ×2       │  │ ▄████▄   repos/cli           │
│ █▀█▀█▀█  3s ago              │  │ █▀█▀█▀█  42s ago             │
│  ▀ ▀ ▀   you: refactor auth… │  │  ▀ ▀ ▀   claude: Allow Bash? │
│          claude: 🛠 Edit     │  │                              │
│                              │  │                              │
│  ▄▄▄▄    IDLE                │  │                              │
│ ▄████▄   docs                │  │                              │
│ █▀█▀█▀█  4m ago              │  │                              │
│  ▀ ▀ ▀                       │  │                              │
└──────────────────────────────┘  └──────────────────────────────┘
```

## Features

- **No agent on the remote.** Pure SSH polling — nothing to install on the boxes you're watching.
- **One pet per project.** Concurrent sessions on the same project collapse into a single avatar with a `×N` count, so a host's panel stays scannable even when you've left ten old sessions lying around.
- **mtime-based phase detection.** Reads `~/.claude/projects/**/*.jsonl` and decides if each session is working, waiting, idle, or asleep.
- **Pixel-art pets** rendered with truecolor half-block characters. Inspired by [recon](https://github.com/gavraz/recon)'s Tamagotchi view.
- **Single static binary.** Cross-compiles cleanly for Linux/macOS, ARM/x86 — perfect for parking on a Pi as a side-monitor.

## Install

### From source

```sh
go install github.com/russ/claudagotchi@latest
```

Or clone and build:

```sh
git clone https://github.com/russ/claudagotchi
cd claudagotchi
make build      # produces ./claudagotchi
make install    # installs to $GOBIN
```

### Cross-compile

```sh
make linux-arm64    # → dist/claudagotchi-linux-arm64  (Raspberry Pi 3/4/5)
make linux-amd64    # → dist/claudagotchi-linux-amd64
make darwin-arm64   # → dist/claudagotchi-darwin-arm64 (Apple Silicon)
make darwin-amd64   # → dist/claudagotchi-darwin-amd64
make release        # all four
```

## Configuration

claudagotchi looks for a TOML config in this order:

1. `--config /path/to/config.toml` (explicit)
2. `$XDG_CONFIG_HOME/claudagotchi/config.toml`
3. `~/.config/claudagotchi/config.toml`
4. `./claudagotchi.toml`

A minimal config:

```toml
hosts = ["host-a", "host-b"]

# Optional.
# poll_interval = "3s"   # how often to re-poll each host
# max_sessions  = 6      # max session avatars per host (extras collapsed to "+N more")
# active_window = "1h"   # sessions older than this are hidden ("0s" disables)
# creatures     = ["slime", "cat", ...]  # roster (omit for all built-ins)
```

For hosts that need explicit SSH overrides, use the detailed `[[host]]` form alongside or instead of the simple list:

```toml
hosts = ["host-a", "host-b"]    # plain names use ~/.ssh/config

[[host]]
name          = "weird-box"
hostname      = "192.168.1.100"
user          = "deploy"
port          = 2222
identity_file = "~/.ssh/id_special"
```

Both forms are concatenated in declaration order. The simple list is shorthand for `[[host]]` entries that have only a `name`. See [`config.example.toml`](config.example.toml).

You can also pass hosts as positional arguments, which override the config:

```sh
claudagotchi pi-1 pi-2 box-3
```

### SSH requirements

Each host name must resolve through your `~/.ssh/config` with **passwordless** (key-based) auth — claudagotchi shells out to `ssh -o BatchMode=yes`, so it will never prompt. If you can `ssh <host> exit` without typing anything, you're good.

It uses SSH `ControlMaster=auto` with a 60s persist, so polls reuse one connection per host instead of paying the handshake cost every cycle.

## Usage

```sh
claudagotchi                      # full TUI
claudagotchi --once               # poll each host once, print plain text, exit
claudagotchi --config ./local.toml
claudagotchi --version
claudagotchi host1 host2          # ad-hoc host list
```

### Keys

| Key            | Action  |
|----------------|---------|
| `q` / `Ctrl-C` / `Esc` | Quit |
| `r`            | Refresh now |

## How it works

For each host on every tick:

1. Open (or reuse) an SSH connection.
2. Run a small inline bash probe that lists the top-level session jsonls under `~/.claude/projects/<project>/<uuid>.jsonl`, sorted by mtime descending, and captures the last two lines of each.
3. Each session is independently classified by its mtime age and last-entry role:

| Activity (mtime age) | Last entry role | Phase        |
|----------------------|-----------------|--------------|
| < 5s                 | any             | **Working**  |
| < 60s                | assistant       | **Waiting** (likely a permission prompt or just-finished) |
| < 60s                | other / user    | **Working**  |
| < 5m                 | any             | **Idle**     |
| ≥ 5m                 | any             | **Sleeping** |

Host-level fallbacks:

- SSH failure → **Unreachable** (broken-egg sprite, error shown in panel)
- Reachable but no `~/.claude/projects` dir or no jsonls → **No sessions** (egg sprite)

## Terminal requirements

- 24-bit truecolor (most modern terminals: iTerm2, Ghostty, Kitty, WezTerm, modern Terminal.app, vscode terminal).
- If you run inside `tmux`, enable truecolor passthrough:

  ```tmux
  set -g default-terminal "tmux-256color"
  set -ga terminal-overrides ",*256col*:Tc"
  ```

## Project layout

```
.
├── main.go                       # CLI entry, flag parsing, --once mode
├── internal/
│   ├── config/                   # TOML config loader
│   ├── poller/                   # SSH probe and phase classification
│   └── ui/                       # Bubble Tea model, view, sprite renderer
├── config.example.toml
├── Makefile
├── LICENSE
└── README.md
```

## Credits

Sprite system inspired by [gavraz/recon](https://github.com/gavraz/recon). Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

## License

[MIT](LICENSE)
