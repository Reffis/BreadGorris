package src

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string = GetConfig("prefix")
	owner  string = GetConfig("owner")
)

type CmdResult struct {
	Result string
	Err    error
}

func Command(s *discordgo.Session, m *discordgo.MessageCreate, args []string) CmdResult {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR", r)
		}
	}()
	cmd := map[string]func(s *discordgo.Session, m *discordgo.MessageCreate, args []string) CmdResult{
		"help":       Help,
		"avatar":     Avatar,
		"choice":     Choice,
		"status":     Status,
		"opensource": OpenSource,
		"dev":        Dev,
		"test":       Test,
		"fox":        Fox,
	}
	if args[0] == prefix {
		return cmd[args[1]](s, m, args)
	}
	return CmdResult{Result: "None", Err: nil}
}
