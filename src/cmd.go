package src

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var (
	prefix string = "gorris"
)

type cmdresult struct {
	result string
	err    error
}

func Command(s *discordgo.Session, m *discordgo.MessageCreate) cmdresult {
	split := strings.Split(m.Content, " ")
	if split[0] == prefix {
		switch split[1] {
		case "help":
			s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "도움말",
				Description: "Prefix: `gorris`",
				Color:       0x00FF00,
				Footer:      &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
			})
			return cmdresult{"help", nil}
		}
	}
	return cmdresult{result: "None", err: nil}
}
