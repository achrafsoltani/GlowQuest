package world

type TileType int

const (
	TileGrass      TileType = iota // '.'
	TileWall                       // 'W'
	TileWater                      // '~'
	TileTree                       // 'T'
	TileSand                       // 'S'
	TileFloor                      // 'F'
	TileStairs                     // '>'
	TileDoorLocked                 // 'D'
	TileDoorOpen                   // 'O'
)

func (t TileType) IsPassable() bool {
	switch t {
	case TileGrass, TileSand, TileFloor, TileDoorOpen, TileStairs:
		return true
	default:
		return false
	}
}

func TileFromChar(c byte) TileType {
	switch c {
	case '.':
		return TileGrass
	case 'W':
		return TileWall
	case '~':
		return TileWater
	case 'T':
		return TileTree
	case 'S':
		return TileSand
	case 'F':
		return TileFloor
	case '>':
		return TileStairs
	case 'D':
		return TileDoorLocked
	case 'O':
		return TileDoorOpen
	default:
		return TileGrass
	}
}
