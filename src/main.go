package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/handler"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// isLocalServer Jadikan false kalau naik prod

func init() {

	initConfig()
	initDiscordGoSession()
	handler.InitCommandHandler()

}

var (
	Token         = flag.String("t", os.Getenv("BOTTOKEN"), "Bot Token")
	isLocalServer = flag.Bool("l", false, "True if local debugging and use local server. Default is false, use Servus Dei config.")
	// GuildID i really cant find a way to put GuildID to config... i did tho
	//GuildID = flag.String("g", "751139261515825162", "Guild ID. If not passed - bot registers commands globally")
	//GuildID = flag.String("g", "751139261515825162", "Guild ID. If not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

func initConfig() {
	flag.Parse()
	if *isLocalServer {
		util.ApplyLocalServerConfig()
	} else {
		util.ApplyServusDeiServerConfig()
	}
}

var dg *discordgo.Session
var err error

func initDiscordGoSession() {
	// Create a new Discord session using the provided bot token.
	dg, err = discordgo.New("Bot " + *Token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}
}

func main() {

	dg.AddHandler(handler.ReadyHandler)
	dg.AddHandler(handler.MessageCreateHandler)
	dg.AddHandler(handler.InteractionCreateHandler)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer dg.Close()

	dg, err = handler.RegisterCommand(dg, util.GuildID)
	if err != nil {
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	if *RemoveCommands {
		err = handler.RemoveCommand(dg, util.GuildID)
		if err != nil {
			return
		}
	}
	log.Println("Gracefully shutting down.")

}
