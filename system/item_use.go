package system

import "github.com/AchrafSoltani/GlowQuest/entity"

// ItemUseResult describes the outcome of using an equipped item.
type ItemUseResult struct {
	UsedItem   entity.EquipItemID
	SwordSwing bool // trigger a sword swing
	// Future: Jump, Dash, Lift, Shoot, etc.
}

// UseItem processes the activation of an equipped item.
// Returns what action should be taken by the game.
func UseItem(item entity.EquipItemID, p *entity.Player) ItemUseResult {
	switch item {
	case entity.EquipSword:
		if p.Inventory.SwordLevel > 0 || p.HasSword {
			return ItemUseResult{UsedItem: item, SwordSwing: true}
		}
	case entity.EquipShield:
		// Shield blocking — will be implemented in Phase 7
	case entity.EquipRocsFeather:
		// Jump — will be implemented in Phase 7
	case entity.EquipPegasusBoots:
		// Dash — will be implemented in Phase 12
	case entity.EquipBomb:
		// Place bomb — will be implemented in Phase 12
	case entity.EquipBow:
		// Shoot arrow — will be implemented later
	case entity.EquipPowerBracelet:
		// Lift — will be implemented in Phase 10
	case entity.EquipHookshot:
		// Grapple — will be implemented in Phase 16
	case entity.EquipMagicRod:
		// Fire — will be implemented in Phase 20
	}
	return ItemUseResult{}
}
