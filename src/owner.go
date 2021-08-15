package src

import (
	"errors"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const OWNER_TEXT = "> `status [Text]`"

func Status(s *discordgo.Session, m *discordgo.MessageCreate, args []string) CmdResult {
	if m.Author.ID != owner {
		return CmdResult{"status", errors.New("You are not the owner of this bot.")}
	}
	items := args[2:]
	text := " "
	for _, item := range items {
		text += item + " "
	}
	s.UpdateStatus(0, text)
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintln("Status updated: `", text, "`"))
	return CmdResult{"status", nil}
}
