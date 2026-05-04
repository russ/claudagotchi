// Package config loads claudagotchi's TOML configuration.
//
// Search order when no explicit path is given:
//
//  1. $XDG_CONFIG_HOME/claudagotchi/config.toml
//  2. $HOME/.config/claudagotchi/config.toml
//  3. ./claudagotchi.toml
package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/BurntSushi/toml"
)

// Config is the user-facing configuration.
type Config struct {
	Hosts        []string `toml:"hosts"`
	PollInterval Duration `toml:"poll_interval"`
	MaxSessions  int      `toml:"max_sessions"`
	ActiveWindow Duration `toml:"active_window"`
	Creatures    []string `toml:"creatures"`
}

// Duration is a time.Duration that decodes from a TOML string like "3s".
type Duration struct{ time.Duration }

// UnmarshalText decodes "3s", "500ms", etc.
func (d *Duration) UnmarshalText(text []byte) error {
	dur, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}
	d.Duration = dur
	return nil
}

// Default returns the built-in defaults used when no config is found.
func Default() Config {
	return Config{
		Hosts:        nil,
		PollInterval: Duration{3 * time.Second},
		MaxSessions:  6,
		ActiveWindow: Duration{1 * time.Hour},
	}
}

// SearchPaths returns the candidate config paths in priority order.
// If explicit is non-empty, only that path is returned.
func SearchPaths(explicit string) []string {
	if explicit != "" {
		return []string{explicit}
	}
	var paths []string
	if x := os.Getenv("XDG_CONFIG_HOME"); x != "" {
		paths = append(paths, filepath.Join(x, "claudagotchi", "config.toml"))
	}
	if home, err := os.UserHomeDir(); err == nil && home != "" {
		paths = append(paths, filepath.Join(home, ".config", "claudagotchi", "config.toml"))
	}
	paths = append(paths, "claudagotchi.toml")
	return paths
}

// Load reads and decodes the first existing config file in priority order,
// merging it onto the defaults. The path that was loaded (or "" if no file
// was found) is returned alongside the config. A missing file is not an
// error unless explicit was non-empty.
func Load(explicit string) (Config, string, error) {
	cfg := Default()
	for _, p := range SearchPaths(explicit) {
		data, err := os.ReadFile(p)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				continue
			}
			return cfg, p, fmt.Errorf("read %s: %w", p, err)
		}
		var loaded Config
		if _, err := toml.Decode(string(data), &loaded); err != nil {
			return cfg, p, fmt.Errorf("parse %s: %w", p, err)
		}
		if len(loaded.Hosts) > 0 {
			cfg.Hosts = loaded.Hosts
		}
		if loaded.PollInterval.Duration > 0 {
			cfg.PollInterval = loaded.PollInterval
		}
		if loaded.MaxSessions > 0 {
			cfg.MaxSessions = loaded.MaxSessions
		}
		if loaded.ActiveWindow.Duration > 0 {
			cfg.ActiveWindow = loaded.ActiveWindow
		}
		if len(loaded.Creatures) > 0 {
			cfg.Creatures = loaded.Creatures
		}
		return cfg, p, nil
	}
	if explicit != "" {
		return cfg, explicit, fmt.Errorf("config file not found: %s", explicit)
	}
	return cfg, "", nil
}
