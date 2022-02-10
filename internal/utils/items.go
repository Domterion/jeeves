package utils

import "math/rand"

type CategoryType string

const (
	ShieldCategory     CategoryType = "shield"
	SaberCategory      CategoryType = "saber"
	HelmetCategory     CategoryType = "helmet"
	ChestplateCategory CategoryType = "chestplate"
	GlovesCategory     CategoryType = "gloves"
	LeggingsCategory   CategoryType = "leggings"
	BootsCategory      CategoryType = "boots"
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

type RarityChance int

const (
	CommonRarityChance    RarityChance = 53
	UncommonRarityChance  RarityChance = 26
	RareRarityChance      RarityChance = 13
	LegendaryRarityChance RarityChance = 7
	MythicRarityChance    RarityChance = 1
)

var RarityChances = map[RarityType]RarityChance{
	CommonRarity:    CommonRarityChance,
	UncommonRarity:  UncommonRarityChance,
	RareRarity:      RareRarityChance,
	LegendaryRarity: LegendaryRarityChance,
	MythicRarity:    MythicRarityChance,
}

var (
	ShieldNames = map[RarityType][]string{
		CommonRarity: {"Common Shield", "Litter Box", "Spaceship Door", "Wooden Shield",
			"Tree Barks", "Door", "Ducttape Shield", "Scrap Shield", "Junk Shield",
			"Glued Sticks", "Taped Sticks", "Glass Shield", "Plastic Shield", "Big Book",
			"Mirror", "Small Sofa", "Cardboard Box", "Picture Frame", "Television", "Fridge Door",
			"Car Door", "Nailed Planks", "Table", "Teacher Desk", "Chess Board", "Normal Shield",
		},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	SaberNames = map[RarityType][]string{
		CommonRarity: {"Common Saber", "Wooden Saber", "Metal Rod", "Plank with Nails", "Broken Saber", "Lost Saber", "Normal Saber", "Common Saber",
			"Kitchen Knife", "Bread Knife", "Short Saber", "Scrap Saber", "Wooden Stick", "Glass Saber", "Plastic Saber",
			"Rotten Saber", "Saber", "Bad Saber", "Unwanted Saber", "Trash Saber", "Junk Saber", "Useless Saber",
		},
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
		CommonRarity: {"Common Chestplate", "Wooden Chestplate", "Normal Chestplate", "Common Chestplate",
			"Junk Chestplate", "Trash Chestplate", "Damaged Chestplate", "Broken Chestplate", "Baggy Hoodie", "7 T-Shirts",
			"Cardboard Chestplate", "Plastic Chestplate", "Rotten Chestplate", "Chestplate of Scrap", "Homemade Chestplate",
			"5-MinuteCrafts Chestplate", "Mcdonalds Chestplate", "Street Chestplate", "Not so Bulletproof Vest", "Vest",
			"Glass Chestplate", "Half'a Chestplate", "Thin Chestplate", "Very Thin Chestplate", "Winter Jacket", "Jacket",
		},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	GloveNames = map[RarityType][]string{
		MythicRarity: {"Soul Stone Gloves", "Space Stone Gloves", "Power Stone Gloves", "Reality Stone Gloves", "Mind Stone Gloves", "Time Stone Gloves"},
	}
	LeggingsNames = map[RarityType][]string{
		CommonRarity:    {"Common Leggings"},
		UncommonRarity:  {},
		RareRarity:      {},
		LegendaryRarity: {},
		MythicRarity:    {},
	}
	BootsNames = map[RarityType][]string{
		CommonRarity:    {"Common Boots"},
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
	case ChestplateCategory:
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
