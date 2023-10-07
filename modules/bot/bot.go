package bot

import (
	"os"
	"os/signal"
	"songwhip_bot/models"
	"songwhip_bot/modules/bot/handlers"
	config "songwhip_bot/modules/config/discord"
	"songwhip_bot/modules/logging"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	App     *models.App
	Session *discordgo.Session
}

func (b *Bot) GetApp() *models.App {
	return b.App
}

func NewBot(app *models.App) (*Bot, error) {
	session, err := discordgo.New("Bot " + app.Config.Discord.BotToken)
	if err != nil {
		return nil, err
	}

	return &Bot{
		App:     app,
		Session: session,
	}, nil
}

func (b *Bot) AddHandlers() {
	b.Session.AddHandler(handlers.NewGuildCreateHandler(b))
	b.Session.AddHandler(handlers.NewGuildDeleteHandler(b))
	b.Session.AddHandler(handlers.NewMessageCreateHandler(b))
	b.Session.AddHandler(handlers.NewReadyHandler(b))
}

func (b *Bot) Start(app *models.App) {
	b.AddHandlers()

	// Set the Intents of the session.PrintLog
	b.Session.Identify.Intents = config.GetIntents(app.Config)

	// Open the websocket and begin listening
	err := b.Session.Open()
	if err != nil {
		logging.PrintLog("error opening connection,", err)
		panic(err)
	}

	logging.PrintLog("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	b.Session.Close()
}
