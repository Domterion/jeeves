package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/domterion/jeeves/internal/startup"
	"github.com/domterion/jeeves/internal/utils"
	"github.com/domterion/jeeves/pkg/commander"
	"github.com/sarulabs/di/v2"
	"github.com/uptrace/bun"
)

func main() {
	builder, _ := di.NewBuilder()

	builder.Add(di.Def{
		Name: utils.DIConfig,
		Build: func(container di.Container) (interface{}, error) {
			c, err := startup.InitConfig()
			if err != nil {
				log.Fatalf("err initializing config: %v\n", err)
			}
			return c, err
		},
	})

	builder.Add(di.Def{
		Name: utils.DIDatabase,
		Build: func(container di.Container) (interface{}, error) {
			c, err := startup.InitDatabase(container)
			if err != nil {
				log.Fatalf("err initializing database: %v\n", err)
			}
			return c, err
		},
		Close: func(obj interface{}) error {
			db := obj.(*bun.DB)
			db.Close()
			return nil
		},
	})

	builder.Add(di.Def{
		Name: utils.DIDiscord,
		Build: func(container di.Container) (interface{}, error) {
			// This is a hacky workaround so we can pass a session to commander
			c, err := discordgo.New()
			if err != nil {
				log.Fatalf("err initializing discord: %v\n", err)
			}
			return c, err
		},
		Close: func(obj interface{}) error {
			discord := obj.(*discordgo.Session)
			discord.Close()
			return nil
		},
	})

	builder.Add(di.Def{
		Name: utils.DICommander,
		Build: func(container di.Container) (interface{}, error) {
			c, err := startup.InitCommander(container)
			if err != nil {
				log.Fatalf("err initializing commander: %v\n", err)
			}
			return c, err
		},
	})

	container := builder.Build()

	_ = container.Get(utils.DIDiscord).(*discordgo.Session)
	_ = container.Get(utils.DICommander).(*commander.Manager)

	startup.InitDiscord(container)

	occurrences := map[utils.RarityType]int{
		utils.CommonRarity:    0,
		utils.UncommonRarity:  0,
		utils.RareRarity:      0,
		utils.LegendaryRarity: 0,
		utils.MythicRarity:    0,
	}

	const iters = 100000

	for i := 0; i < iters; i++ {
		o := utils.Earth.GetRandomRarity(utils.BootsCategory)
		occurrences[o] += 1
	}

	fmt.Printf(`Iters: %d
Chances:
Common: %d
Uncommon: %d
Rare: %d
Legendary: %d
Mythic: %d

Occurrences:
-
Common: %d
Uncommon: %d
Rare: %d
Legendary: %d
Mythic: %d	
`, iters, utils.CommonRarityChance, utils.UncommonRarityChance, utils.RareRarityChance, utils.LegendaryRarityChance, utils.MythicRarityChance,
		occurrences[utils.CommonRarity], occurrences[utils.UncommonRarity], occurrences[utils.RareRarity], occurrences[utils.LegendaryRarity], occurrences[utils.MythicRarity])

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop

	container.DeleteWithSubContainers()
}
