package utils

import "math/rand"

type PlanetType string

const (
	EarthPlanet PlanetType = "earth"
)

type Planet struct {
	Difficulty int
	Loot       map[CategoryType][]TierType // The loot table is a map of categories to a slice of tiers that can be found on that planet
	Enemies    []Enemy
}

type Enemy struct {
	Name   string
	Health int
	Damage int
	Armor  int
}

func (p *Planet) GetRandomRarity(category CategoryType) TierType {
	var tiers []TierType

	for _, tier := range p.Loot[category] {
		chance := int(TierChances[tier])
		for i := 0; i < chance; i++ {
			tiers = append(tiers, tier)
		}
	}

	randomIndex := rand.Intn(len(tiers))
	tier := tiers[randomIndex]

	return tier
}

var (
	Earth = Planet{
		Difficulty: 0,
		Loot: map[CategoryType][]TierType{
			ShieldCategory: {
				DTier,
			},
			SaberCategory: {
				DTier,
			},
			HelmetCategory: {
				DTier,
			},
			ChestplateCategory: {
				DTier,
			},
			GlovesCategory: {
				DTier,
			},
			BootsCategory: {
				DTier, CTier, BTier, ATier, STier, SPlusTier,
			},
		},
		Enemies: []Enemy{},
	}
)

var Planets = map[PlanetType]Planet{EarthPlanet: Earth}
