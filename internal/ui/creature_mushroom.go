package ui

import "github.com/russ/claudagotchi/internal/poller"

var (
	mushroomPalette = map[int]rgb{
		1: {220, 60, 60},   // red cap
		2: {170, 30, 30},   // dark red shadow
		3: {40, 40, 40},    // pupil black
		4: {255, 255, 255}, // white (cap spots, eye whites)
		5: {240, 220, 180}, // cream stem
		6: {60, 40, 40},    // dark mouth
		7: {165, 200, 255}, // pale blue Z
	}
	mushroomAlertPalette = map[int]rgb{
		1: {240, 90, 50},   // alarmed orange-red cap
		2: {180, 50, 30},   // dark red
		3: {40, 40, 40},
		4: {255, 255, 255},
		5: {240, 220, 180}, // stem stays cream
		6: {60, 40, 40},
		7: {255, 220, 60},  // yellow sweat
	}
)

func init() {
	registerCreature(&Creature{
		Name: "mushroom",
		// Toadstool: rounded red cap (4 rows) with four white spots, cream
		// stem (8 wide, narrower than cap so the cap overhangs at the rim),
		// cute face on the stem.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — eyes open, cross-eyed (pupils on inner side).
				{palette: mushroomPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 0},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 4, 4, 5, 5, 4, 4, 5, 0},
					{0, 5, 4, 3, 5, 5, 3, 4, 5, 0},
					{0, 5, 5, 5, 6, 6, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
				// Frame 2 — blink: closed-line eyes, no whites.
				{palette: mushroomPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 0},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 3, 3, 5, 5, 3, 3, 5, 0},
					{0, 5, 5, 5, 6, 6, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
				// Frame 3 — back to frame 1.
				{palette: mushroomPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 0},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 4, 4, 5, 5, 4, 4, 5, 0},
					{0, 5, 4, 3, 5, 5, 3, 4, 5, 0},
					{0, 5, 5, 5, 6, 6, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — alarmed orange cap, pupils outer, wider mouth,
				// sweat drop in upper-right corner.
				{palette: mushroomAlertPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 7},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 0},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 4, 4, 5, 5, 4, 4, 5, 0},
					{0, 5, 3, 4, 5, 5, 4, 3, 5, 0},
					{0, 5, 5, 6, 6, 6, 6, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
				// Frame 2 — sweat slides down a row.
				{palette: mushroomAlertPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 7},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 4, 4, 5, 5, 4, 4, 5, 0},
					{0, 5, 3, 4, 5, 5, 4, 3, 5, 0},
					{0, 5, 5, 6, 6, 6, 6, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Closed-line eyes, calm.
				{palette: mushroomPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 0},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 3, 3, 5, 5, 3, 3, 5, 0},
					{0, 5, 5, 5, 6, 6, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Z trail to the upper-right (single cap-corner column),
				// closed-line eyes.
				{palette: mushroomPalette, grid: [][]int{
					{0, 0, 0, 1, 1, 1, 1, 0, 0, 7},
					{0, 0, 1, 4, 1, 1, 4, 1, 0, 7},
					{0, 1, 1, 1, 4, 4, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 5, 3, 3, 5, 5, 3, 3, 5, 0},
					{0, 5, 5, 5, 6, 6, 5, 5, 5, 0},
					{0, 5, 5, 5, 5, 5, 5, 5, 5, 0},
					{0, 0, 5, 5, 5, 5, 5, 5, 0, 0},
				}},
			},
		},
	})
}
