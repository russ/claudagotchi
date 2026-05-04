package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	frogPalette = map[int]rgb{
		1: {130, 200, 100}, // green body
		2: {80, 140, 60},   // dark green shadow
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // eye white
		5: {180, 230, 140}, // lighter belly green
		6: {50, 80, 40},    // dark mouth
		7: {165, 200, 255}, // pale blue Z
		8: {255, 180, 200}, // pink tongue (used only in waiting)
	}
	frogAlertPalette = map[int]rgb{
		1: {160, 220, 80},  // brighter alert green
		2: {110, 170, 60},  // dark green
		3: {40, 40, 40},
		4: {255, 255, 255},
		5: {200, 240, 140}, // lighter belly
		6: {60, 90, 40},
		7: {255, 220, 60},  // yellow sweat
		8: {255, 160, 180}, // pink tongue
	}
)

func init() {
	registerCreature(&Creature{
		Name: "frog",
		// Front-view frog: two bulging eyes poking up above the head, wide
		// grin with upturned smile corners, green body with lighter belly,
		// small webbed feet at the bottom.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — default, cross-eyed (pupils on inner side), smiling.
				{palette: frogPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 4, 4, 1, 0, 0, 1, 4, 4, 1},
					{1, 4, 3, 1, 0, 0, 1, 3, 4, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 1, 1, 1, 1, 1, 1, 6, 0},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
				// Frame 2 — blink: eye whites and pupils replaced by a thin
				// horizontal closed-eye line.
				{palette: frogPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 1, 1, 0, 0, 1, 1, 1, 1},
					{1, 3, 3, 1, 0, 0, 1, 3, 3, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 1, 1, 1, 1, 1, 1, 6, 0},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
				// Frame 3 — back to frame 1 (gives a slow blink rhythm).
				{palette: frogPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 4, 4, 1, 0, 0, 1, 4, 4, 1},
					{1, 4, 3, 1, 0, 0, 1, 3, 4, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 1, 1, 1, 1, 1, 1, 6, 0},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — worried, pupils on outer side, mouth open with
				// pink tongue (color 8), sweat drop above the right eye.
				{palette: frogAlertPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 7},
					{1, 4, 4, 1, 0, 0, 1, 4, 4, 1},
					{1, 3, 4, 1, 0, 0, 1, 4, 3, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 6, 6, 6, 6, 1, 1, 1},
					{1, 1, 6, 8, 8, 8, 8, 6, 1, 1},
					{1, 1, 1, 6, 6, 6, 6, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
				// Frame 2 — sweat gone (flicker), mouth slightly wider.
				{palette: frogAlertPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 4, 4, 1, 0, 0, 1, 4, 4, 1},
					{1, 3, 4, 1, 0, 0, 1, 4, 3, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 6, 8, 8, 8, 8, 8, 8, 6, 1},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Eyes squinted (closed-line), peaceful smile.
				{palette: frogPalette, grid: [][]int{
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 1, 1, 0, 0, 1, 1, 1, 1},
					{1, 3, 3, 1, 0, 0, 1, 3, 3, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 6, 1, 1, 1, 1, 1, 1, 6, 0},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Compressed pose, eyes closed, Z drifting up.
				{palette: frogPalette, grid: [][]int{
					{0, 0, 0, 0, 0, 7, 0, 0, 0, 0},
					{0, 0, 0, 0, 7, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 3, 3, 1, 0, 0, 1, 3, 3, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 6, 6, 6, 6, 6, 6, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 5, 5, 5, 5, 5, 5, 1, 0},
					{0, 0, 1, 1, 0, 0, 1, 1, 0, 0},
				}},
			},
		},
	})
}
