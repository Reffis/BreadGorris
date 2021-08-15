package src

import "github.com/bwmarrin/discordgo"

const UTIL_TEXT = "> `help`"

func Help(s *discordgo.Session, m *discordgo.MessageCreate) CmdResult {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "도움말",
		Description: "**Prefix**: `gorris`",
		Color:       0x00FF00,
		Footer:      &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Utility", Value: UTIL_TEXT, Inline: true},
			{Name: "Moderator", Value: MODER_TEXT, Inline: true},
		},
	})
	return CmdResult{"help", nil}
}
