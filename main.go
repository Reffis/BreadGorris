package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	cmd "BreadGorris/src"
)

func main() {
	dg, err := discordgo.New("Bot " + cmd.GetConfig("token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageHandler)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	dg.UpdateStatus(0, "gorris help")
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	result := cmd.Command(s, m, strings.Split(m.Content, " "))
	if result.Result == "None" {
		return
	}
	if result.Err != nil {
		log.Println("[ Error ] [", m.Author.ID, "]", result.Err)
		return
	}
	log.Println("[ Command ] [", m.Author.ID, "]", result.Result)

}
