package src

import (
	"github.com/bwmarrin/discordgo"
)

const UTIL_TEXT = "> `help`, `opensource`, `dev`"

func Help(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "도움말",
		Description: "**Prefix**: `gorris`",
		Color:       RandomColors(),
		Footer:      &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Utility", Value: UTIL_TEXT, Inline: true},
			{Name: "Fun", Value: FUN_TEXT, Inline: false},
			{Name: "Owner", Value: OWNER_TEXT, Inline: true},
			{Name: "Image", Value: IMAGE_TEXT, Inline: false},
		},
	})
	return CmdResult{"help", nil}
}

func OpenSource(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "오픈소스",
		Description: "[소스코드](https://github.com/Reffis/BreadGorris)\n\n[MIT LICENSE](https://www.mit.edu/~amini/LICENSE.md)",
		Color:       WHITE,
		Footer:      &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
	})
	return CmdResult{"help", nil}
}

func Dev(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "개발자",
		Description: `		
**빵켓 (! Bread Cat#0002)**

[Fn79](https://github.com/Fn79)
[Reffis](https://github.com/Reffis)
		`,
		Color:     WHITE,
		Footer:    &discordgo.MessageEmbedFooter{Text: m.Author.Username, IconURL: m.Author.AvatarURL("1024")},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: "https://cdn.discordapp.com/avatars/760688241447141395/a_6fc82732de318b2894554c7906ab9b91.gif?size=1024"},
	})
	return CmdResult{"dev", nil}
}

func Test(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	s.ChannelMessageSend(m.ChannelID, "테su투")
	return CmdResult{"Test", nil}
}
