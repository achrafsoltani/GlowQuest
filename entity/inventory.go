package entity

// Inventory holds all player items, ammo, and equipment state.
type Inventory struct {
	// Currency & ammo
	Rupees    int
	Keys      int
	Bombs     int
	BombsMax  int
	Arrows    int
	ArrowsMax int

	// Equipment
	OwnedItems    map[EquipItemID]bool
	ButtonA       EquipItemID
	ButtonB       EquipItemID
	SwordLevel    int // 0=none, 1=L1, 2=L2
	ShieldLevel   int // 0=none, 1=shield, 2=mirror
	BraceletLevel int // 0=none, 1=L1, 2=L2

	// Collectables
	Instruments     [8]bool
	SecretSeashells int
	TradingItem     int // 0-14
	Songs           [3]bool
}

// NewInventory creates an empty inventory.
func NewInventory() Inventory {
	return Inventory{
		OwnedItems: make(map[EquipItemID]bool),
		BombsMax:   30,
		ArrowsMax:  30,
	}
}

// OwnedItemsList returns a sorted list of owned equippable items.
func (inv *Inventory) OwnedItemsList() []EquipItemID {
	var items []EquipItemID
	// Return in a fixed order matching EquipItemID values
	for id := EquipItemID(1); id <= EquipMagicPowder; id++ {
		if inv.OwnedItems[id] {
			items = append(items, id)
		}
	}
	return items
}
