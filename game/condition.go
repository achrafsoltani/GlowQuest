package game

import (
	"strings"

	"github.com/AchrafSoltani/GlowQuest/entity"
)

// CheckCondition evaluates a condition string against the current game state.
// Supported formats:
//   - "flag:key" — true if quest flag is set
//   - "!flag:key" — true if quest flag is NOT set
//   - "item:name" — true if player owns the equip item
//   - "dungeon:N" — true if dungeon N is completed
//   - "" — always true (no condition)
func CheckCondition(cond string, quest *QuestState, inv *entity.Inventory) bool {
	if cond == "" {
		return true
	}

	negate := false
	c := cond
	if strings.HasPrefix(c, "!") {
		negate = true
		c = c[1:]
	}

	result := false

	if strings.HasPrefix(c, "flag:") {
		key := c[5:]
		result = quest.HasFlag(key)
	} else if strings.HasPrefix(c, "item:") {
		name := c[5:]
		result = checkItemOwned(name, inv)
	} else if strings.HasPrefix(c, "dungeon:") {
		numStr := c[8:]
		num := 0
		for _, ch := range numStr {
			if ch >= '0' && ch <= '9' {
				num = num*10 + int(ch-'0')
			}
		}
		result = quest.IsDungeonComplete(num)
	}

	if negate {
		return !result
	}
	return result
}

func checkItemOwned(name string, inv *entity.Inventory) bool {
	switch name {
	case "sword":
		return inv.SwordLevel > 0
	case "shield":
		return inv.ShieldLevel > 0
	case "bow":
		return inv.OwnedItems[entity.EquipBow]
	case "bombs":
		return inv.OwnedItems[entity.EquipBomb]
	case "rocs_feather":
		return inv.OwnedItems[entity.EquipRocsFeather]
	case "pegasus_boots":
		return inv.OwnedItems[entity.EquipPegasusBoots]
	case "power_bracelet":
		return inv.BraceletLevel > 0
	case "flippers":
		return inv.OwnedItems[entity.EquipFlippers]
	case "hookshot":
		return inv.OwnedItems[entity.EquipHookshot]
	case "magic_rod":
		return inv.OwnedItems[entity.EquipMagicRod]
	case "boomerang":
		return inv.OwnedItems[entity.EquipBoomerang]
	case "ocarina":
		return inv.OwnedItems[entity.EquipOcarina]
	default:
		return false
	}
}
