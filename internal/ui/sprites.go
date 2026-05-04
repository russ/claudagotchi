package ui

import (
	"fmt"
	"strings"

	"github.com/russ/claudagotchi/internal/poller"
)

const spriteWidth = 10

type rgb struct{ r, g, b uint8 }

// sprite is one 10x10 pixel grid plus its palette. 0 = transparent; positive
// values index into palette.
type sprite struct {
	grid    [][]int
	palette map[int]rgb
}

// renderGrid emits the half-block ANSI for a sprite (10 chars wide × 5 lines
// tall for a 10x10 grid).
func renderGrid(s sprite) string {
	rows := len(s.grid)
	if rows == 0 {
		return ""
	}
	cols := len(s.grid[0])

	var b strings.Builder
	for y := 0; y < rows; y += 2 {
		for x := 0; x < cols; x++ {
			top := s.grid[y][x]
			bot := 0
			if y+1 < rows {
				bot = s.grid[y+1][x]
			}
			switch {
			case top == 0 && bot == 0:
				b.WriteByte(' ')
			case top == 0:
				c := s.palette[bot]
				fmt.Fprintf(&b, "\x1b[38;2;%d;%d;%dm▄\x1b[0m", c.r, c.g, c.b)
			case bot == 0:
				c := s.palette[top]
				fmt.Fprintf(&b, "\x1b[38;2;%d;%d;%dm▀\x1b[0m", c.r, c.g, c.b)
			default:
				ct := s.palette[top]
				cb := s.palette[bot]
				fmt.Fprintf(&b, "\x1b[38;2;%d;%d;%dm\x1b[48;2;%d;%d;%dm▀\x1b[0m",
					ct.r, ct.g, ct.b, cb.r, cb.g, cb.b)
			}
		}
		if y+2 < rows {
			b.WriteByte('\n')
		}
	}
	return b.String()
}

// Host-level palettes, used for fallback sprites that are not creature-bound.
var (
	eggPalette = map[int]rgb{
		1: {255, 250, 230},
		2: {220, 200, 170},
		3: {180, 220, 180},
		4: {200, 180, 150},
	}
	brokenPalette = map[int]rgb{
		1: {180, 90, 90},
		2: {130, 60, 60},
		3: {40, 40, 40},
		4: {220, 180, 180},
		5: {120, 60, 60},
	}
)

// hostSprites covers the cases that don't belong to a single session: a host
// is unreachable, or has no Claude sessions at all.
var hostSprites = map[poller.Phase][]sprite{
	poller.PhaseNoSessions: {
		{palette: eggPalette, grid: [][]int{
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 3, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 3, 1, 1, 1, 3, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 1, 2, 1, 2, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}},
	},
	poller.PhaseUnreachable: {
		{palette: brokenPalette, grid: [][]int{
			{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 4, 3, 3, 4, 1, 1, 0},
			{0, 1, 1, 3, 1, 1, 3, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 3, 3, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 5, 0, 0, 5, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}},
	},
	poller.PhaseUnknown: {
		{palette: eggPalette, grid: [][]int{
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}},
	},
}

// RenderHostSprite renders a host-level fallback sprite (egg, broken egg,
// unknown). Returns "" if no sprite is registered for phase.
func RenderHostSprite(phase poller.Phase, frame int) string {
	frames, ok := hostSprites[phase]
	if !ok || len(frames) == 0 {
		return ""
	}
	return renderGrid(frames[frame%len(frames)])
}
