package bot

import (
	"fmt"
	"songwhip_bot/modules/bot"
	config "songwhip_bot/modules/config/discord"
	"songwhip_bot/modules/config/services"
)

func Start() {
	// Load services yaml
	services, err := services.GetStreamingServices()

	if err != nil {
		fmt.Println("error loading services yaml", err)
		panic(err)
	}

	// Load config yaml
	config, err := config.GetConfig()

	if err != nil {
		fmt.Println("error loading config yaml", err)
		panic(err)
	}

	bot.SetupBot(config, services)
}
