package utils

import "math/rand"

type PlanetType string

const (
	EarthPlanet PlanetType = "earth"
)

type Planet struct {
	Difficulty int
	Loot       map[CategoryType][]RarityType // The loot table is a map of categories to a slice of rarities that can be found on that planet
}

func (p *Planet) GetRandomRarity(category CategoryType) RarityType {
	var rarities []RarityType

	for _, rarity := range p.Loot[category] {
		chance := int(RarityChances[rarity])
		for i := 0; i < chance; i++ {
			rarities = append(rarities, rarity)
		}
	}

	randomIndex := rand.Intn(len(rarities))
	rarity := rarities[randomIndex]

	return rarity
}

var (
	Earth = Planet{
		Difficulty: 0,
		Loot: map[CategoryType][]RarityType{
			ShieldCategory: {
				CommonRarity,
			},
			SaberCategory: {
				CommonRarity,
			},
			HelmetCategory: {
				CommonRarity,
			},
			ChestplateCategory: {
				CommonRarity,
			},
			GlovesCategory: {
				CommonRarity,
			},
			BootsCategory: {
				CommonRarity, UncommonRarity, RareRarity, LegendaryRarity, MythicRarity,
			},
		},
	}
)

/*

Each planet should have a few stats:

Difficulty : The overall difficulty of the planet, to include the enemies and the challenge to get there
Loot       : The loot table for the planet, the possible loot and kinds
Enemies    : *wink*

*/
