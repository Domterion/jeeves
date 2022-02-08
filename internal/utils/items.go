package utils

import "math/rand"

type CategoryType string

const (
	ShieldCategory      CategoryType = "shield"
	SaberCategory       CategoryType = "saber"
	HelmetCategory      CategoryType = "helmet"
	ChestplateCateegory CategoryType = "chestplate"
	GlovesCategory      CategoryType = "gloves"
	LeggingsCategory    CategoryType = "leggings"
	BootsCategory       CategoryType = "boots"
)

type SlotType string

const (
	HeadSlot      SlotType = "head"
	TorsoSlot     SlotType = "torso"
	HandsSlot     SlotType = "hands"
	LegsSlot      SlotType = "legs"
	FeetSlot      SlotType = "feet"
	EquipmentSlot SlotType = "equipment"
)

type RarityType string

const (
	CommonRarity    RarityType = "common"
	UncommonRarity  RarityType = "uncommon"
	RareRarity      RarityType = "rare"
	LegendaryRarity RarityType = "legendary"
	MythicRarity    RarityType = "mythic"
)

var (
	ShieldNames = map[RarityType][]string{
		CommonRarity:    {},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	SaberNames = map[RarityType][]string{
		CommonRarity:    {"Common Saber"},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {"Darth Vader's Lightsaber", "Void Saber"},
	}
	HelmetNames = map[RarityType][]string{
		CommonRarity:    {"Common Space Helmet"},
		UncommonRarity:  {},
		RareRarity:      {"Stormtrooper's Helmet"},
		LegendaryRarity: {},
		MythicRarity:    {"Darth Vader's Helmet"},
	}
	ChestplateNames = map[RarityType][]string{
		CommonRarity:    {},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	GloveNames = map[RarityType][]string{
		MythicRarity: {"Soul Stone Gloves", "Space Stone Gloves", "Power Stone Gloves", "Reality Stone Gloves", "Mind Stone Gloves", "Time Stone Gloves"},
	}
	LeggingsNames = map[RarityType][]string{
		CommonRarity:    {},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	BootsNames = map[RarityType][]string{
		CommonRarity:    {},
		UncommonRarity:  {},
		RareRarity:      {"Rocket Boots"},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
)

func GetRandomItemName(category CategoryType, rarity RarityType) string {
	var names []string

	switch category {
	case ShieldCategory:
		names = ShieldNames[rarity]
	case SaberCategory:
		names = SaberNames[rarity]
	case HelmetCategory:
		names = HelmetNames[rarity]
	case ChestplateCateegory:
		names = ChestplateNames[rarity]
	case GlovesCategory:
		names = GloveNames[rarity]
	case LeggingsCategory:
		names = LeggingsNames[rarity]
	case BootsCategory:
		names = BootsNames[rarity]
	}

	randomIndex := rand.Intn(len(names))
	name := names[randomIndex]

	return name
}
