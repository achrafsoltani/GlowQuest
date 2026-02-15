package game

import "github.com/AchrafSoltani/GlowQuest/entity"

type DialogueState struct {
	Active      bool
	NPC         *entity.NPC
	Lines       []string // the active lines being displayed
	CurrentLine int
}

// Start begins a dialogue with the given NPC using its default dialogue.
func (d *DialogueState) Start(npc *entity.NPC) {
	d.Active = true
	d.NPC = npc
	d.Lines = npc.Dialogue
	d.CurrentLine = 0
}

// StartWithLines begins a dialogue with the given NPC using specific lines.
func (d *DialogueState) StartWithLines(npc *entity.NPC, lines []string) {
	d.Active = true
	d.NPC = npc
	d.Lines = lines
	d.CurrentLine = 0
}

// Advance moves to the next set of lines. Returns true if dialogue is finished.
func (d *DialogueState) Advance() bool {
	d.CurrentLine += 3
	if d.CurrentLine >= len(d.Lines) {
		d.Active = false
		d.NPC = nil
		d.Lines = nil
		d.CurrentLine = 0
		return true
	}
	return false
}

// HasMore returns true if there are more lines after the current page.
func (d *DialogueState) HasMore() bool {
	return d.CurrentLine+3 < len(d.Lines)
}
