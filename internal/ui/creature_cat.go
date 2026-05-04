package ui

import "github.com/russ/claudagotchi/internal/poller"

var catPalette = map[int]rgb{
	1: {255, 190, 70},   // bright orange tabby body
	2: {200, 140, 30},   // dark orange shadow
	3: {40, 40, 40},     // pupil black
	4: {255, 255, 255},  // eye white highlight
	5: {245, 175, 195},  // pink (nose, inner ears, tongue)
	6: {130, 70, 30},    // dark stripe (whiskers, mouth, closed-eye line)
	7: {165, 200, 255},  // pale blue (Z, sweat drop)
}

func init() {
	registerCreature(&Creature{
		Name: "cat",
		// Head-only cat face: triangular ears with pink interiors at the top
		// corners, big eyes with white highlights, pink nose, whisker stripes
		// on the cheeks, small mouth. Sleeping uses a loaf-cat pose.
		Frames: map[poller.Phase][]sprite{
			poller.PhaseWorking: {
				// Frame 1 — alert, eyes open, calm small mouth.
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 4, 4, 1, 1, 1, 1, 4, 4, 1},
					{1, 3, 3, 1, 1, 1, 1, 3, 3, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 1, 1, 6, 6, 1, 1, 6, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
				// Frame 2 — happy blink, wider closed-eye smile.
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 6, 6, 1, 1, 1, 1, 6, 6, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 6, 1, 6, 6, 1, 6, 6, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
				// Frame 3 — back to neutral (same as frame 1, makes the blink
				// land once per ~900ms cycle instead of every 600ms).
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 4, 4, 1, 1, 1, 1, 4, 4, 1},
					{1, 3, 3, 1, 1, 1, 1, 3, 3, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 1, 1, 6, 6, 1, 1, 6, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
			},
			poller.PhaseWaiting: {
				// Frame 1 — concerned, sweat drop above the right ear.
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 7},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 7},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 4, 4, 1, 1, 1, 1, 4, 4, 1},
					{1, 3, 3, 1, 1, 1, 1, 3, 3, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 1, 1, 6, 6, 1, 1, 6, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
				// Frame 2 — open mouth meowing, pink tongue showing.
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 7},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 7},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 4, 4, 1, 1, 1, 1, 4, 4, 1},
					{1, 3, 3, 1, 1, 1, 1, 3, 3, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 1, 6, 5, 5, 6, 1, 6, 1},
					{0, 1, 1, 1, 6, 6, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
			},
			poller.PhaseIdle: {
				// Squinted eyes (just the dark line), calm closed mouth.
				{palette: catPalette, grid: [][]int{
					{0, 0, 1, 0, 0, 0, 0, 1, 0, 0},
					{0, 1, 1, 0, 0, 0, 0, 1, 1, 0},
					{1, 1, 5, 1, 0, 0, 1, 5, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{1, 3, 3, 1, 1, 1, 1, 3, 3, 1},
					{1, 6, 1, 1, 5, 5, 1, 1, 6, 1},
					{1, 6, 1, 1, 6, 6, 1, 1, 6, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
				}},
			},
			poller.PhaseSleeping: {
				// Loaf cat, ears tucked away, closed-eye lines, Z drifting up.
				{palette: catPalette, grid: [][]int{
					{0, 0, 0, 0, 7, 0, 0, 0, 0, 0},
					{0, 0, 0, 7, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
					{0, 1, 1, 6, 1, 1, 6, 1, 1, 0},
					{0, 1, 6, 1, 5, 5, 1, 6, 1, 0},
					{0, 1, 1, 1, 6, 6, 1, 1, 1, 0},
					{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
					{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
				}},
			},
		},
	})
}
