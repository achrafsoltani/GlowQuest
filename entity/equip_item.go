package entity

// EquipItemID identifies an equippable item for the A/B button system.
type EquipItemID int

const (
	EquipNone          EquipItemID = 0
	EquipSword         EquipItemID = 1
	EquipShield        EquipItemID = 2
	EquipBow           EquipItemID = 3
	EquipBomb          EquipItemID = 4
	EquipRocsFeather   EquipItemID = 5
	EquipPegasusBoots  EquipItemID = 6
	EquipPowerBracelet EquipItemID = 7
	EquipFlippers      EquipItemID = 8
	EquipHookshot      EquipItemID = 9
	EquipMagicRod      EquipItemID = 10
	EquipBoomerang     EquipItemID = 11
	EquipOcarina       EquipItemID = 12
	EquipShovel        EquipItemID = 13
	EquipMagicPowder   EquipItemID = 14
)

// EquipItemName returns a display name for the item.
func EquipItemName(id EquipItemID) string {
	switch id {
	case EquipSword:
		return "Sword"
	case EquipShield:
		return "Shield"
	case EquipBow:
		return "Bow"
	case EquipBomb:
		return "Bombs"
	case EquipRocsFeather:
		return "Feather"
	case EquipPegasusBoots:
		return "Boots"
	case EquipPowerBracelet:
		return "Bracelet"
	case EquipFlippers:
		return "Flippers"
	case EquipHookshot:
		return "Hookshot"
	case EquipMagicRod:
		return "Magic Rod"
	case EquipBoomerang:
		return "Boomerang"
	case EquipOcarina:
		return "Ocarina"
	case EquipShovel:
		return "Shovel"
	case EquipMagicPowder:
		return "Powder"
	default:
		return ""
	}
}
