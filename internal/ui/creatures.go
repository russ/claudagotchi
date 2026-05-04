package ui

import (
	"hash/fnv"
	"sort"

	"github.com/russ/claudagotchi/internal/poller"
)

// Creature is a named pet with sprites for each session phase.
// Sleeping/idle/working/waiting are required; missing phases fall back to a
// blank string at render time.
type Creature struct {
	Name   string
	Frames map[poller.Phase][]sprite
}

// RenderSprite emits the sprite for the given session phase at the given
// animation frame. Returns "" if the creature has no art for that phase.
func (c *Creature) RenderSprite(phase poller.Phase, frame int) string {
	frames, ok := c.Frames[phase]
	if !ok || len(frames) == 0 {
		return ""
	}
	return renderGrid(frames[frame%len(frames)])
}

// FrameCount returns the number of animation frames the creature has for
// phase. Zero means the creature has no art for that phase.
func (c *Creature) FrameCount(phase poller.Phase) int {
	return len(c.Frames[phase])
}

// LookupCreature returns a built-in creature by name, or nil.
func LookupCreature(name string) *Creature {
	return creatureRegistry[name]
}

// creatureRegistry holds every built-in creature, keyed by Name. Populated by
// init() in each creature_*.go file.
var creatureRegistry = map[string]*Creature{}

func registerCreature(c *Creature) {
	creatureRegistry[c.Name] = c
}

// AvailableCreatures returns all built-in creature names, sorted, suitable
// for documenting the `creatures` config key.
func AvailableCreatures() []string {
	names := make([]string, 0, len(creatureRegistry))
	for n := range creatureRegistry {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

// PickCreature deterministically selects a creature for a project on a host.
// allowed is the configured roster; an empty roster means "all available".
// Unknown names in allowed are silently skipped. If everything is filtered
// out, the first available creature is returned.
func PickCreature(host, project string, allowed []string) *Creature {
	roster := allowed
	if len(roster) == 0 {
		roster = AvailableCreatures()
	}
	valid := make([]string, 0, len(roster))
	for _, n := range roster {
		if _, ok := creatureRegistry[n]; ok {
			valid = append(valid, n)
		}
	}
	if len(valid) == 0 {
		valid = AvailableCreatures()
	}
	if len(valid) == 0 {
		return nil
	}
	h := fnv.New32a()
	_, _ = h.Write([]byte(host))
	_, _ = h.Write([]byte{'/'})
	_, _ = h.Write([]byte(project))
	return creatureRegistry[valid[h.Sum32()%uint32(len(valid))]]
}
