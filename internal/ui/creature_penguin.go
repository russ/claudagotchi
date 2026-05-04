package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	penguinPalette = map[int]rgb{
		1: {45, 55, 75},     // dark slate body
		2: {25, 35, 55},     // shadow
		3: {15, 15, 25},     // pupil black
		4: {255, 250, 245},  // white belly / cheek
		5: {255, 175, 70},   // orange beak / feet
		6: {40, 30, 30},     // dark inside of open beak
		7: {165, 200, 255},  // pale blue Z
	}
	penguinAlertPalette = map[int]rgb{
		1: {55, 55, 80},     // slightly warmer slate
		2: {30, 35, 55},     // shadow
		3: {15, 15, 25},     // pupil
		4: {255, 250, 245},  // white
		5: {255, 175, 70},   // orange
		6: {40, 30, 30},     // dark mouth
		7: {255, 220, 60},   // yellow sweat
	}
)

func init() {
	registerCreature(&Creature{
		Name: "penguin",
		// Stout front-view penguin: round head, white cheek patches, orange
		// beak, dark slate body that bulges at the middle with a white
		// belly, orange feet at the bottom.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — eyes open, cross-eyed cute look, beak closed.
				{palette: penguinPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 1, 1, 1, 1, 4, 1, 0},
					{0, 1, 4, 3, 1, 1, 3, 4, 1, 0},
					{0, 0, 1, 1, 5, 5, 1, 1, 0, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 2 — blink: cheeks stay, eyes close to a thin line.
				{palette: penguinPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 1, 1, 1, 1, 4, 1, 0},
					{0, 1, 3, 3, 1, 1, 3, 3, 1, 0},
					{0, 0, 1, 1, 5, 5, 1, 1, 0, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 3 — back to frame 1 (gives a slow blink rhythm).
				{palette: penguinPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 1, 1, 1, 1, 4, 1, 0},
					{0, 1, 4, 3, 1, 1, 3, 4, 1, 0},
					{0, 0, 1, 1, 5, 5, 1, 1, 0, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — wider eyes, pupils on outer side, beak open
				// (orange edges + dark interior), sweat drop above right side.
				{palette: penguinAlertPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 7},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 1, 3, 4, 1, 1, 4, 3, 1, 0},
					{0, 1, 1, 5, 6, 6, 5, 1, 1, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
				// Frame 2 — beak even wider (squawking), sweat moves down.
				{palette: penguinAlertPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 7},
					{0, 1, 4, 4, 1, 1, 4, 4, 1, 0},
					{0, 1, 3, 4, 1, 1, 4, 3, 1, 0},
					{0, 1, 5, 6, 6, 6, 6, 5, 1, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Closed-line eyes, peaceful.
				{palette: penguinPalette, grid: [][]int{
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 4, 1, 1, 1, 1, 4, 1, 0},
					{0, 1, 3, 3, 1, 1, 3, 3, 1, 0},
					{0, 0, 1, 1, 5, 5, 1, 1, 0, 0},
					{0, 1, 1, 4, 4, 4, 4, 1, 1, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Compressed pose (penguins sleep standing), Z drifting up.
				{palette: penguinPalette, grid: [][]int{
					{0, 0, 0, 0, 7, 0, 0, 0, 0, 0},
					{0, 0, 0, 7, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 4, 1, 1, 1, 1, 4, 1, 0},
					{0, 1, 3, 3, 1, 1, 3, 3, 1, 0},
					{0, 0, 1, 1, 5, 5, 1, 1, 0, 0},
					{1, 1, 4, 4, 4, 4, 4, 4, 1, 1},
					{0, 1, 4, 4, 4, 4, 4, 4, 1, 0},
					{0, 0, 5, 5, 0, 0, 5, 5, 0, 0},
				}},
			},
		},
	})
}
