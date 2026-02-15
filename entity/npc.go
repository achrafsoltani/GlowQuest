package entity

// DialogueOption represents a conditional dialogue that an NPC can speak.
// Conditions are checked in order; the first matching condition wins.
type DialogueOption struct {
	Condition string
	Lines     []string
}

type NPC struct {
	ID            string
	X, Y          float64
	Width, Height int
	Dir           Direction
	Name          string
	Dialogue      []string         // default/fallback dialogue
	Dialogues     []DialogueOption // conditional dialogues (checked first)
}

func NewNPC(id string, x, y float64, dir Direction, name string, dialogue []string, dialogues []DialogueOption) *NPC {
	return &NPC{
		ID:        id,
		X:         x,
		Y:         y,
		Width:     14,
		Height:    14,
		Dir:       dir,
		Name:      name,
		Dialogue:  dialogue,
		Dialogues: dialogues,
	}
}

func (n *NPC) CenterX() float64 { return n.X + float64(n.Width)/2 }
func (n *NPC) CenterY() float64 { return n.Y + float64(n.Height)/2 }
