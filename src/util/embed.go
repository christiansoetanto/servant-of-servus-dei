package util

import "github.com/bwmarrin/discordgo"

const (
	WelcomeTitle    = "Welcome to Servus Dei!"
	LogoURL         = "https://cdn.discordapp.com/avatars/767426889294938112/0e100e9fec18866892ed0c875b341926.png"
	Author          = "Servant of Servus Dei"
	WelcomeImageURL = "https://media.discordapp.net/attachments/751174184733900801/974282451096576041/unknown.png"
	FooterText      = "©2022 | soetanto@ServusDei"
)

func EmbedBuilder(title string, description string) *discordgo.MessageEmbed {
	embed := &discordgo.MessageEmbed{
		Type:        discordgo.EmbedTypeRich,
		Title:       title,
		Description: description,
		Footer: &discordgo.MessageEmbedFooter{
			Text: FooterText,
		},
		Image: &discordgo.MessageEmbedImage{
			URL: WelcomeImageURL,
		},
		Author: &discordgo.MessageEmbedAuthor{
			Name:    Author,
			IconURL: LogoURL,
		},
		Fields: nil,
	}
	return embed
}
