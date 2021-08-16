package src

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

const IMAGE_TEXT = "> `fox`"

func decode(link string) (map[string]string, error) {
	resp, err := http.Get(link)
	if err != nil {
		return map[string]string{}, errors.New("Error Getting Image")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	decoder := json.NewDecoder(bytes.NewReader(body))
	var json map[string]string
	if decoder.Decode(&json) != nil {
		return (map[string]string{}), errors.New("Json Decode Error")
	}
	return json, nil
}

func Fox(s *discordgo.Session, m *discordgo.MessageCreate, _ []string) CmdResult {
	json, err := decode("https://randomfox.ca/floof/")
	if err != nil {
		return CmdResult{"fox", errors.New("Json Decode Error")}
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Fox",
		URL:   json["link"],
		Color: RandomColors(),
		Image: &discordgo.MessageEmbedImage{URL: json["image"]},
	})
	return CmdResult{"fox", nil}
}
