package src

import "github.com/bwmarrin/discordgo"

const UTIL_TEXT = "> `help`"

func Help(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "도움말",
		Description: "**Prefix**: `gorris`",
		Color:       WHITE,
		Footer:      &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Utility", Value: UTIL_TEXT, Inline: true},
			{Name: "Moderator", Value: MODER_TEXT, Inline: true},
			{Name: "Fun", Value: MODER_TEXT, Inline: false},
		},
	})
	return CmdResult{"help", nil}
}
