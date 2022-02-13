package utils

import "math/rand"

type CategoryType string
type SlotType string
type TierType string
type TierChance int

const (
	ShieldCategory     CategoryType = "shield"
	SaberCategory      CategoryType = "saber"
	HelmetCategory     CategoryType = "helmet"
	ChestplateCategory CategoryType = "chestplate"
	GlovesCategory     CategoryType = "gloves"
	LeggingsCategory   CategoryType = "leggings"
	BootsCategory      CategoryType = "boots"

	HeadSlot      SlotType = "head"
	TorsoSlot     SlotType = "torso"
	HandsSlot     SlotType = "hands"
	LegsSlot      SlotType = "legs"
	FeetSlot      SlotType = "feet"
	EquipmentSlot SlotType = "equipment"

	DTier     TierType = "D"
	CTier     TierType = "C"
	BTier     TierType = "B"
	ATier     TierType = "A"
	STier     TierType = "S"
	SPlusTier TierType = "S+"

	DTierChance     TierChance = 51
	CTierChance     TierChance = 24
	BTierChance     TierChance = 13
	ATierChance     TierChance = 7
	STierChance     TierChance = 4
	SPlusTierChance TierChance = 1
)

var (
	TierChances = map[TierType]TierChance{
		DTier:     DTierChance,
		CTier:     CTierChance,
		BTier:     BTierChance,
		ATier:     ATierChance,
		STier:     STierChance,
		SPlusTier: SPlusTierChance,
	}

	ShieldNames = map[TierType][]string{
		DTier: {"Common Shield", "Litter Box", "Spaceship Door", "Wooden Shield",
			"Tree Barks", "Door", "Ducttape Shield", "Scrap Shield", "Junk Shield",
			"Glued Sticks", "Taped Sticks", "Glass Shield", "Plastic Shield", "Big Book",
			"Mirror", "Small Sofa", "Cardboard Box", "Picture Frame", "Television", "Fridge Door",
			"Car Door", "Nailed Planks", "Table", "Teacher Desk", "Chess Board", "Normal Shield",
		},
		CTier:     {},
		BTier:     {},
		ATier:     {},
		STier:     {},
		SPlusTier: {},
	}
	SaberNames = map[TierType][]string{
		DTier: {"Common Saber", "Wooden Saber", "Metal Rod", "Plank with Nails", "Broken Saber", "Lost Saber", "Normal Saber", "Common Saber",
			"Kitchen Knife", "Bread Knife", "Short Saber", "Scrap Saber", "Wooden Stick", "Glass Saber", "Plastic Saber",
			"Rotten Saber", "Saber", "Bad Saber", "Unwanted Saber", "Trash Saber", "Junk Saber", "Useless Saber",
		},
		CTier:     {},
		BTier:     {},
		ATier:     {},
		STier:     {},
		SPlusTier: {"Darth Vader's Lightsaber", "Void Saber"},
	}
	HelmetNames = map[TierType][]string{
		DTier:     {"Common Space Helmet"},
		CTier:     {},
		BTier:     {"Stormtrooper's Helmet"},
		ATier:     {},
		STier:     {},
		SPlusTier: {"Darth Vader's Helmet"},
	}
	ChestplateNames = map[TierType][]string{
		DTier: {"Common Chestplate", "Wooden Chestplate", "Normal Chestplate", "Common Chestplate",
			"Junk Chestplate", "Trash Chestplate", "Damaged Chestplate", "Broken Chestplate", "Baggy Hoodie", "7 T-Shirts",
			"Cardboard Chestplate", "Plastic Chestplate", "Rotten Chestplate", "Chestplate of Scrap", "Homemade Chestplate",
			"5-MinuteCrafts Chestplate", "Mcdonalds Chestplate", "Street Chestplate", "Not so Bulletproof Vest", "Vest",
			"Glass Chestplate", "Half'a Chestplate", "Thin Chestplate", "Very Thin Chestplate", "Winter Jacket", "Jacket",
		},
		CTier:     {},
		BTier:     {},
		ATier:     {},
		STier:     {},
		SPlusTier: {},
	}
	GloveNames = map[TierType][]string{
		SPlusTier: {"Soul Stone Gloves", "Space Stone Gloves", "Power Stone Gloves", "Reality Stone Gloves", "Mind Stone Gloves", "Time Stone Gloves"},
	}
	LeggingsNames = map[TierType][]string{
		DTier:     {"Common Leggings"},
		CTier:     {},
		BTier:     {},
		ATier:     {},
		STier:     {},
		SPlusTier: {},
	}
	BootsNames = map[TierType][]string{
		DTier:     {"Common Boots"},
		CTier:     {},
		BTier:     {"Rocket Boots"},
		ATier:     {},
		STier:     {},
		SPlusTier: {},
	}
)

func GetRandomItemName(category CategoryType, tier TierType) string {
	var names []string

	switch category {
	case ShieldCategory:
		names = ShieldNames[tier]
	case SaberCategory:
		names = SaberNames[tier]
	case HelmetCategory:
		names = HelmetNames[tier]
	case ChestplateCategory:
		names = ChestplateNames[tier]
	case GlovesCategory:
		names = GloveNames[tier]
	case LeggingsCategory:
		names = LeggingsNames[tier]
	case BootsCategory:
		names = BootsNames[tier]
	}

	randomIndex := rand.Intn(len(names))
	name := names[randomIndex]

	return name
}
