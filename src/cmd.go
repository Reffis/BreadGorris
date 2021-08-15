package src

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string = "gorris"
)

type CmdResult struct {
	Result string
	Err    error
}

func Command(s *discordgo.Session, m *discordgo.MessageCreate) CmdResult {
	cmd := map[string]func(s *discordgo.Session, m *discordgo.MessageCreate) CmdResult{
		"help": Help,
	}
	split := strings.Split(m.Content, " ")
	if split[0] == prefix {
		return cmd[split[1]](s, m)
	}
	return CmdResult{Result: "None", Err: nil}
}
