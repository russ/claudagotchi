package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	duckPalette = map[int]rgb{
		1: {255, 220, 80},  // yellow body
		2: {200, 170, 30},  // dark yellow shadow (unused)
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // eye white
		5: {255, 140, 50},  // orange bill / feet
		6: {180, 80, 30},   // dark inside of open bill
		7: {165, 200, 255}, // pale blue Z / sweat
	}
	duckAlertPalette = map[int]rgb{
		1: {255, 200, 60},  // slightly more saturated yellow
		2: {180, 130, 30},
		3: {40, 40, 40},
		4: {255, 255, 255},
		5: {255, 140, 50},
		6: {180, 80, 30},
		7: {165, 200, 255}, // light blue sweat (visible against yellow)
	}
)

func init() {
	registerCreature(&Creature{
		Name: "duck",
		// Side-profile rubber duck, facing right: round head with single eye
		// on the right, flat orange bill protruding right, body widening
		// underneath, two orange feet at the bottom.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — eye open (white highlight + pupil), bill closed.
				{palette: duckPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 4, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 1, 1, 5, 5, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 5, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 2 — blink: closed-eye line (2-pixel horizontal).
				{palette: duckPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 3, 1, 5, 5, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 5, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 3 — back to frame 1.
				{palette: duckPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 4, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 1, 1, 5, 5, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 5, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — bill opens (dark interior), sweat drop top-right.
				{palette: duckAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 7},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 4, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 1, 1, 5, 6, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 6, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 2 — sweat slides down a row.
				{palette: duckAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 7},
					{0, 0, 0, 1, 4, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 1, 1, 5, 6, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 6, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Closed-line eye, calm.
				{palette: duckPalette, grid: [][]int{
					{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 3, 1, 5, 5, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 5, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Z trail in the upper-left empty area, eye closed.
				{palette: duckPalette, grid: [][]int{
					{0, 7, 0, 0, 1, 1, 1, 1, 0, 0},
					{0, 0, 7, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 0, 1, 3, 3, 1, 5, 5, 5},
					{0, 0, 0, 1, 1, 1, 1, 5, 5, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
		},
	})
}
