package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	robotPalette = map[int]rgb{
		1: {200, 200, 220}, // silver body
		2: {100, 100, 130}, // dark metal (shadow, antenna shaft, feet)
		3: {40, 40, 40},    // black
		4: {80, 220, 255},  // cyan LED glow (eyes, antenna ball)
		5: {255, 80, 80},   // red (reserved)
		6: {50, 50, 70},    // dark mouth grille / dark eye socket
		7: {165, 200, 255}, // pale blue Z mark
	}
	robotAlertPalette = map[int]rgb{
		1: {200, 200, 220}, // silver body
		2: {100, 100, 130}, // dark metal
		3: {40, 40, 40},    // black
		4: {255, 80, 80},   // RED replaces cyan for waiting state
		5: {255, 220, 60},  // yellow accent
		6: {50, 50, 70},    // dark mouth
		7: {165, 200, 255}, // unused but consistent
	}
)

func init() {
	registerCreature(&Creature{
		Name: "robot",
		// Boxy retro head (8 cols wide, cols 1-8) with antenna at col 4,
		// LED eyes, side bolts, narrower speaker grille mouth, and short
		// feet at the bottom corners.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — antenna ball lit, eyes glowing cyan.
				{palette: robotPalette, grid: [][]int{
					{0, 0, 0, 0, 4, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
				// Frame 2 — antenna ball off (dark), eyes still on.
				{palette: robotPalette, grid: [][]int{
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
				// Frame 3 — back to lit (lit/dark/lit blink rhythm).
				{palette: robotPalette, grid: [][]int{
					{0, 0, 0, 0, 4, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — alert palette: red antenna ball + red eyes.
				{palette: robotAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 4, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
				// Frame 2 — alarm pulse: antenna and eyes both dark.
				{palette: robotAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 6, 6, 1, 1, 6, 6, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
			},
			poller.PhaseIdle: {
				// Antenna off, eyes still cyan, no animation.
				{palette: robotPalette, grid: [][]int{
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 2, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Powered down: no antenna, eye sockets dark, Z drifting up.
				{palette: robotPalette, grid: [][]int{
					{0, 0, 0, 0, 7, 0, 0, 0, 0, 0},
					{0, 0, 0, 7, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 6, 6, 1, 1, 6, 6, 1, 0},
					{0, 2, 6, 6, 1, 1, 6, 6, 2, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 6, 6, 6, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 2, 2, 0, 0, 0, 0, 2, 2, 0},
				}},
			},
		},
	})
}
