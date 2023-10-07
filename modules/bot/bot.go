package bot

import (
	"fmt"
	"os"
	"os/signal"
	"songwhip_bot/models"
	"songwhip_bot/modules/handlers"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func setIntents(dg *discordgo.Session) {
	dg.Identify.Intents = discordgo.IntentsDirectMessages
}

func SetupBot(config *models.Config, services *models.Services) {
	dg, err := discordgo.New("Bot " + config.Discord.BotToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		panic(err)
	}

	// Add handlers to the session
	handlers.AddHandlers(dg)

	// Set the Intents of the session
	setIntents(dg)

	// Open the websocket and begin listening
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		panic(err)
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
