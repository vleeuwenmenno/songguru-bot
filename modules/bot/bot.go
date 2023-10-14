package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"songguru_bot/models"
	"songguru_bot/modules/bot/handlers"
	config "songguru_bot/modules/config/app"
	"songguru_bot/modules/logging"
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
	b.Session.AddHandler(handlers.NewMessageCreateHandler(b))
	b.Session.AddHandler(handlers.NewReadyHandler(b))
	b.Session.AddHandler(handlers.NewInteractionCreateHandler(b))

	b.Session.AddHandler(handlers.NewGuildCreateHandler(b))
	b.Session.AddHandler(handlers.NewGuildDeleteHandler(b))

	b.Session.AddHandler(handlers.NewGuildRoleCreateHandler(b))
	b.Session.AddHandler(handlers.NewGuildRoleUpdateHandler(b))
	b.Session.AddHandler(handlers.NewGuildRoleDeleteHandler(b))
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
