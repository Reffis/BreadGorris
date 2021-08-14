package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"internal/commands"
	"internal/events"

	"github.com/bwmarrin/discordgo"
)

func main() {
	token, _ := ioutil.ReadFile("token")
	dg, err := discordgo.New(string(token))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	registerEvents(dg)
	registerCommands(dg)

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

func registerEvents(s *discordgo.Session) {
	joinLeaveHandler := events.NewJoinLeaveHandler()
	s.AddHandler(joinLeaveHandler.HandlerJoin)
	s.AddHandler(joinLeaveHandler.HandlerLeave)

	s.AddHandler(events.NewReadyHandler().Handler)
	s.AddHandler(events.NewMessageHandler().Handler)
}

func registerCommands(s *discordgo.Session) {
	cmdHandler := commands.NewCommandHandler("!")
	cmdHandler.OnError = func(err error, ctx *commands.Context) {
		ctx.Session.ChannelMessageSend(ctx.Message.ChannelID,
			fmt.Sprintf("Command Execution failed: %s", err.Error()))
	}

	cmdHandler.RegisterCommand(&commands.CmdPing{})
	cmdHandler.RegisterMiddleware(&commands.MwPermissions{})

	s.AddHandler(cmdHandler.HandleMessage)
}
