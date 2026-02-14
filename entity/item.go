package entity

import "math"

type ItemType int

const (
	ItemHeart ItemType = iota
	ItemRupee
	ItemKey
	ItemSword
	ItemHeartContainer
)

type Item struct {
	Type          ItemType
	X, Y          float64
	Width, Height int
	Collected     bool
	BobTimer      float64
}

func NewItem(typ ItemType, x, y float64) *Item {
	return &Item{
		Type:   typ,
		X:      x,
		Y:      y,
		Width:  12,
		Height: 12,
	}
}

func (i *Item) Update(dt float64) {
	i.BobTimer += dt
}

// BobOffset returns the vertical bob offset for rendering.
func (i *Item) BobOffset() float64 {
	return math.Sin(i.BobTimer*4.0) * 1.0
}
