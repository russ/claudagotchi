package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	ghostPalette = map[int]rgb{
		1: {240, 240, 255}, // ghost white
		2: {180, 180, 210}, // shadow grey
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // eye highlight (pure white)
		5: {255, 200, 220}, // pink blush
		6: {60, 60, 80},    // dark mouth
		7: {165, 200, 255}, // pale blue Z / sweat
	}
	ghostAlertPalette = map[int]rgb{
		1: {220, 190, 235}, // spooky purple
		2: {170, 130, 200}, // darker purple
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // eye highlight
		5: {255, 200, 220}, // pink blush (unused in waiting)
		6: {80, 40, 100},   // dark mouth
		7: {255, 220, 60},  // yellow sweat drop
	}
)

func init() {
	registerCreature(&Creature{
		Name: "ghost",
		// Floating sheet ghost: tapered dome top, body, and a wavy bottom
		// edge that shifts between frames to suggest hovering.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — eyes open, blush on cheeks, wave A.
				{palette: ghostPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 4, 1, 1, 3, 4, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 5, 1, 1, 1, 1, 1, 1, 5, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
				// Frame 2 — wave B (shifted), same face.
				{palette: ghostPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 4, 1, 1, 3, 4, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 5, 1, 1, 1, 1, 1, 1, 5, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
				// Frame 3 — back to wave A (gives a slow sway cycle).
				{palette: ghostPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 4, 1, 1, 3, 4, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 5, 1, 1, 1, 1, 1, 1, 5, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — shocked O-mouth, no blush, spooky-purple palette.
				{palette: ghostAlertPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 4, 1, 1, 3, 4, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 6, 1, 1, 6, 1, 1, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
				// Frame 2 — sweat drop appears, wave shifts.
				{palette: ghostAlertPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 7},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 7},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 4, 1, 1, 3, 4, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 6, 1, 1, 6, 1, 1, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
			},
			poller.PhaseIdle: {
				// Eyes closed (squinted dashes), blush, peaceful smile.
				{palette: ghostPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 5, 1, 1, 1, 1, 1, 1, 5, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
			},
			poller.PhaseSleeping: {
				// Compressed lower-on-grid pose, Z drifting up to the right.
				{palette: ghostPalette, grid: [][]int{
					{0, 0, 0, 0, 0, 0, 0, 7, 0, 0},
					{0, 0, 0, 0, 0, 0, 7, 0, 0, 0},
					{0, 0, 0, 0, 0, 7, 0, 0, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 3, 3, 1, 1, 3, 3, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 6, 6, 1, 1, 1, 1},
					{1, 1, 0, 1, 1, 0, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				}},
			},
		},
	})
}
