package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	cmd "BreadGorris/src"
)

func main() {
	token, _ := ioutil.ReadFile("token")
	dg, err := discordgo.New(string(token))
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
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	result := cmd.Command(s, m)
	if result.Result == "None" {
		return
	}
	if result.Err != nil {
		log.Println("[Error] [", m.Author.ID, "]", result.Err)
		return
	}
	log.Println("[Command] [", m.Author.ID, "]", result.Result)

}
