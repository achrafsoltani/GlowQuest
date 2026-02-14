package entity

type NPC struct {
	X, Y          float64
	Width, Height int
	Dir           Direction
	Name          string
	Dialogue      []string
}

func NewNPC(x, y float64, dir Direction, name string, dialogue []string) *NPC {
	return &NPC{
		X:        x,
		Y:        y,
		Width:    14,
		Height:   14,
		Dir:      dir,
		Name:     name,
		Dialogue: dialogue,
	}
}

func (n *NPC) CenterX() float64 { return n.X + float64(n.Width)/2 }
func (n *NPC) CenterY() float64 { return n.Y + float64(n.Height)/2 }
