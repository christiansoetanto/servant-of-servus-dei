package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/calendar"
	"github.com/christiansoetanto/servant-of-servus-dei/src/handler"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {

	initConfig()
	initDiscordGoSession()
	handler.InitCommandHandler()

}

var (
	Token          = flag.String("t", os.Getenv("BOTTOKEN"), "Bot Token")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

func initConfig() {
	flag.Parse()
}

var s *discordgo.Session
var err error

func initDiscordGoSession() {
	s, err = discordgo.New("Bot " + *Token)
	if err != nil {
		log.Fatalf("error creating Discord session: %v", err)
		return
	}
}

const DailyCron = "@daily"

func cronJob(s *discordgo.Session) {
	c := cron.New()
	_, err := c.AddFunc(DailyCron, calendar.CalendarCronJob(s))
	if err != nil {
		return
	}
	c.Start()

}

func main() {

	s.AddHandler(handler.ReadyHandler)
	s.AddHandler(handler.MessageCreateHandlerQuestionOne)
	s.AddHandler(handler.InteractionCreateHandler)
	s.AddHandler(handler.MessageReactionAddHandler)

	s.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMessageReactions | discordgo.IntentsDirectMessages
	err = s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	defer s.Close()

	//err = handler.RemoveCommand(s)
	//	if err != nil {
	//		return
	//	}

	s, err = handler.RegisterCommand(s)
	if err != nil {
		return
	}

	cronJob(s)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	//syscall.SIGTERM,
	signal.Notify(sc, syscall.SIGINT)
	<-sc

	log.Println("Gracefully shutting down.")

}
