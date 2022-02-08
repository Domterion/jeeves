package utils

type CategoryType string

const (
	Shield     CategoryType = "shield"
	Saber      CategoryType = "saber"
	Helmet     CategoryType = "helmet"
	Chestplate CategoryType = "chestplate"
	Leggings   CategoryType = "leggings"
	Boots      CategoryType = "boots"
)

type SlotType string

const (
	Head  SlotType = "head"
	Torso SlotType = "torso"
	Legs  SlotType = "legs"
	Feet  SlotType = "feet"
)

type RarityType string

const (
	Common   RarityType = "common"
	Uncommon RarityType = "uncommon"
	Rare     RarityType = "rare"
	Mythic   RarityType = "mythic"
)

var (
	ShieldNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
	SaberNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
	HelmetNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
	ChestplateNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
	LeggingsNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
	BootsNames = map[RarityType][]string{
		Common:   {},
		Uncommon: {},
		Rare:     {},
		Mythic:   {},
	}
)
