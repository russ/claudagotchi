package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	turtlePalette = map[int]rgb{
		1: {90, 150, 80},   // shell green
		2: {60, 110, 50},   // shell darker (scute markings)
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // white (unused for 1-px eye, reserved)
		5: {160, 200, 120}, // head / leg light green
		6: {40, 80, 30},    // dark accent
		7: {165, 200, 255}, // pale blue Z / sweat
	}
	turtleAlertPalette = map[int]rgb{
		1: {160, 180, 70},  // alarmed yellow-green shell
		2: {110, 130, 50},  // dark olive
		3: {40, 40, 40},
		4: {255, 255, 255},
		5: {160, 200, 120}, // head still light green (but mostly retracted)
		6: {40, 80, 30},
		7: {255, 220, 60},  // yellow sweat
	}
)

func init() {
	registerCreature(&Creature{
		Name: "turtle",
		// Side-profile turtle facing right: dome shell with scattered scute
		// markings, head poking out at the lower right (with single eye),
		// legs visible at the corners of the body.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — head out, eye visible.
				{palette: turtlePalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 5, 5},
					{0, 5, 0, 0, 0, 0, 0, 0, 5, 3},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
				// Frame 2 — blink: eye pixel disappears.
				{palette: turtlePalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 5, 5},
					{0, 5, 0, 0, 0, 0, 0, 0, 5, 5},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
				// Frame 3 — back to frame 1.
				{palette: turtlePalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 5, 5},
					{0, 5, 0, 0, 0, 0, 0, 0, 5, 3},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — head retracted into shell, sweat drop above,
				// alarmed yellow-green shell palette.
				{palette: turtleAlertPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 7, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 5, 0, 0, 0, 0, 0, 0, 0, 0},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
				// Frame 2 — sweat slides down a row.
				{palette: turtleAlertPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 7},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 5, 0, 0, 0, 0, 0, 0, 0, 0},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Head out, eye closed (no pupil), peaceful.
				{palette: turtlePalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 5, 5},
					{0, 5, 0, 0, 0, 0, 0, 0, 5, 5},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Head still out but eye closed, Z trail upper-left.
				{palette: turtlePalette, grid: [][]int{
					{0, 7, 0, 1, 1, 1, 1, 0, 0, 0},
					{7, 0, 1, 1, 2, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 2, 1, 1, 0},
					{1, 1, 2, 1, 1, 1, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 5},
					{0, 0, 1, 1, 1, 1, 1, 1, 5, 5},
					{0, 5, 0, 0, 0, 0, 0, 0, 5, 5},
					{5, 5, 0, 0, 0, 0, 0, 5, 5, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				}},
			},
		},
	})
}
