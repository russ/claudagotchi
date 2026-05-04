package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	cactusPalette = map[int]rgb{
		1: {90, 180, 80},   // green body
		2: {60, 130, 50},   // dark green shadow
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // eye white
		5: {255, 150, 200}, // pink flower
		6: {60, 80, 40},    // dark mouth
		7: {165, 200, 255}, // pale blue Z
	}
	cactusAlertPalette = map[int]rgb{
		1: {180, 200, 80},  // parched yellow-green
		2: {130, 150, 50},  // dark olive
		3: {40, 40, 40},
		4: {255, 255, 255},
		5: {255, 150, 200}, // flower stays pink
		6: {60, 80, 40},
		7: {255, 220, 60},  // yellow sweat
	}
)

func init() {
	registerCreature(&Creature{
		Name: "cactus",
		// Saguaro silhouette: central trunk with two short arms branching
		// upward at the shoulders, pink flower on top, cute face on the
		// trunk's lower portion.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — eyes open, cross-eyed cute look.
				{palette: cactusPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 4, 4, 1, 1, 4, 4, 0, 0},
					{0, 0, 4, 3, 1, 1, 3, 4, 0, 0},
					{0, 0, 1, 1, 6, 6, 1, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
				// Frame 2 — blink: closed eye line, no whites.
				{palette: cactusPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 3, 3, 1, 1, 3, 3, 0, 0},
					{0, 0, 1, 1, 6, 6, 1, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
				// Frame 3 — back to frame 1.
				{palette: cactusPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 4, 4, 1, 1, 4, 4, 0, 0},
					{0, 0, 4, 3, 1, 1, 3, 4, 0, 0},
					{0, 0, 1, 1, 6, 6, 1, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — pupils outer (worried), wider mouth, sweat drop,
				// parched-yellow palette.
				{palette: cactusAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 7},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 4, 4, 1, 1, 4, 4, 0, 0},
					{0, 0, 3, 4, 1, 1, 4, 3, 0, 0},
					{0, 0, 1, 6, 6, 6, 6, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
				// Frame 2 — sweat drips down a row.
				{palette: cactusAlertPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 7},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 4, 4, 1, 1, 4, 4, 0, 0},
					{0, 0, 3, 4, 1, 1, 4, 3, 0, 0},
					{0, 0, 1, 6, 6, 6, 6, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Closed-line eyes, calm.
				{palette: cactusPalette, grid: [][]int{
					{0, 0, 0, 0, 5, 5, 0, 0, 0, 0},
					{0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 3, 3, 1, 1, 3, 3, 0, 0},
					{0, 0, 1, 1, 6, 6, 1, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Z drifting up to the upper-left of the flower.
				{palette: cactusPalette, grid: [][]int{
					{0, 0, 7, 0, 5, 5, 0, 0, 0, 0},
					{0, 7, 0, 0, 1, 1, 0, 0, 0, 0},
					{0, 1, 0, 0, 1, 1, 0, 0, 1, 0},
					{0, 1, 1, 0, 1, 1, 0, 1, 1, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 3, 3, 1, 1, 3, 3, 0, 0},
					{0, 0, 1, 1, 6, 6, 1, 1, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
				}},
			},
		},
	})
}
