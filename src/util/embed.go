package util

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
)

const (
	WelcomeTitle     = "Welcome to Servus Dei!"
	LogoURL          = "https://cdn.discordapp.com/avatars/767426889294938112/0e100e9fec18866892ed0c875b341926.png"
	Author           = "Servant of Servus Dei"
	WelcomeImageURL  = "https://media.discordapp.net/attachments/751174184733900801/974282451096576041/unknown.png"
	WelcomeImage2URL = "https://media.discordapp.net/attachments/751174152588623912/975368929008558130/Screenshot_2022-05-11_at_11.42.51_PM.png"
	FooterText       = "2022 | Made for Servus Dei by soetanto™"
)

func RandomWelcomeImage() string {
	in := []string{WelcomeImageURL, WelcomeImage2URL}
	return in[rand.Intn(len(in))]
}
func EmbedBuilder(title, description, welcomeImageUrl string) *discordgo.MessageEmbed {

	embed := &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       title,
		Description: description,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    FooterText,
			IconURL: LogoURL,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: welcomeImageUrl,
		},

		Fields: nil,
	}
	return embed
}
