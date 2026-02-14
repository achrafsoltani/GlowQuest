package entity

type Direction int

const (
	DirDown Direction = iota
	DirUp
	DirLeft
	DirRight
)

func (d Direction) DX() float64 {
	switch d {
	case DirLeft:
		return -1
	case DirRight:
		return 1
	default:
		return 0
	}
}

func (d Direction) DY() float64 {
	switch d {
	case DirUp:
		return -1
	case DirDown:
		return 1
	default:
		return 0
	}
}
