package game

import "github.com/AchrafSoltani/GlowQuest/entity"

type DialogueState struct {
	Active      bool
	NPC         *entity.NPC
	CurrentLine int
}

// Start begins a dialogue with the given NPC.
func (d *DialogueState) Start(npc *entity.NPC) {
	d.Active = true
	d.NPC = npc
	d.CurrentLine = 0
}

// Advance moves to the next set of lines. Returns true if dialogue is finished.
func (d *DialogueState) Advance() bool {
	d.CurrentLine += 3
	if d.CurrentLine >= len(d.NPC.Dialogue) {
		d.Active = false
		d.NPC = nil
		d.CurrentLine = 0
		return true
	}
	return false
}

// HasMore returns true if there are more lines after the current page.
func (d *DialogueState) HasMore() bool {
	return d.CurrentLine+3 < len(d.NPC.Dialogue)
}
