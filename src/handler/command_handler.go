package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"log"
)

const (
	acknowledgementMessageFormatWithRole    = "Verification of user <@%s> with role <@&%s> is successful.\nThank you for using my service. Beep. Boop.\n"
	acknowledgementMessageFormatWithoutRole = "Verification of user <@%s> is successful.\nThank you for using my service. Beep. Boop.\n"
	welcomeMessageFormat                    = "Welcome to Servus Dei, <@%s>! We are happy to have you! Make sure you check out <#%s> to gain access to the various channels we offer and please do visit <#%s> so you can understand our server better and take use of everything we have to offer. God Bless!"
)

var (
	Ping               = "ping"
	SDVerify           = "sdverify"
	commandHandlers    = map[string]func(dg *discordgo.Session, i *discordgo.InteractionCreate) error{}
	commands           []*discordgo.ApplicationCommand
	registeredCommands []*discordgo.ApplicationCommand
)

func InitCommandHandler() {
	commandHandlers = map[string]func(dg *discordgo.Session, i *discordgo.InteractionCreate) error{
		Ping: func(dg *discordgo.Session, i *discordgo.InteractionCreate) error {
			err := dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong!",
				},
			})
			if err != nil {
				return err
			}
			return nil
		},
		SDVerify: func(dg *discordgo.Session, i *discordgo.InteractionCreate) error {

			err := dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Processing... please wait...",
				},
			})
			if err != nil {
				return err
			}
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options
			guildId := i.GuildID

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			acknowledgementMessageArgs := make([]interface{}, 0, len(options))
			acknowledgementMessageFormat := acknowledgementMessageFormatWithoutRole
			welcomeMessageArgs := make([]interface{}, 0, 1)

			var user *discordgo.User

			userOpt, ok := optionMap["user-option"]
			if ok {
				acknowledgementMessageArgs = append(acknowledgementMessageArgs, userOpt.UserValue(nil).ID)
				user = userOpt.UserValue(nil)
				welcomeMessageArgs = append(welcomeMessageArgs, user.ID)
				welcomeMessageArgs = append(welcomeMessageArgs, config.Config[guildId].Channel.ReactionRoles)
				welcomeMessageArgs = append(welcomeMessageArgs, config.Config[guildId].Channel.ServerInformation)

				//actually i dont need to put this in here, because user is required anyway. but just to be safe haha
				roleOpt, ok := optionMap["role-option"]
				if ok {
					acknowledgementMessageFormat = acknowledgementMessageFormatWithRole
					roleType := roleOpt.StringValue()
					roleId := config.Config[guildId].ReligionRoleMapping[config.ReligionRoleType(roleType)]
					acknowledgementMessageArgs = append(acknowledgementMessageArgs, roleId)
					err := dg.GuildMemberRoleAdd(guildId, user.ID, string(roleId))
					if err != nil {
						fmt.Println(err)
						return err
					}
				}

				err := dg.GuildMemberRoleAdd(guildId, user.ID, config.Config[guildId].Role.ApprovedUser)
				if err != nil {
					return err
				}
				err = dg.GuildMemberRoleRemove(guildId, user.ID, config.Config[guildId].Role.Vetting)
				if err != nil {
					return err
				}
				err = dg.GuildMemberRoleRemove(guildId, user.ID, config.Config[guildId].Role.VettingQuestioning)
				if err != nil {
					return err
				}

			} else {
				_, err := dg.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
					Content: "Please choose user.",
				})
				if err != nil {
					return err
				}
			}

			_, err = dg.ChannelMessageSend(config.Config[guildId].Channel.GeneralDiscussion, user.Mention())
			if err != nil {
				return err
			}

			_, err = dg.ChannelMessageSendEmbed(config.Config[guildId].Channel.GeneralDiscussion, util.EmbedBuilder(util.WelcomeTitle, fmt.Sprintf(welcomeMessageFormat, welcomeMessageArgs...), util.RandomWelcomeImage()))
			if err != nil {
				return err
			}
			//send interaction response later, based on whether the welcome is success or not
			_, err = dg.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
				Content: fmt.Sprintf(
					acknowledgementMessageFormat,
					acknowledgementMessageArgs...,
				),
			})

			if err != nil {
				return err
			}
			return nil
		},
	}
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        Ping,
			Description: "Ping",
		},
		{
			Name:        SDVerify,
			Description: "Command for verifying new peeps and welcoming them",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionUser,
					Name:        "user-option",
					Description: "User to verify",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "role-option",
					Description: "Religion role to give",
					Required:    true,
					Choices:     buildReligionRoleOptionChoices(),
				},
			},
		},
	}
	registeredCommands = make([]*discordgo.ApplicationCommand, len(commands))

}

func InteractionCreateHandler(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		err := h(dg, i)
		if err != nil {
			log.Println(err)
		}
	}
}

func RegisterCommand(dg *discordgo.Session, guildId string) (*discordgo.Session, error) {
	log.Println("Adding commands...")
	for i, v := range commands {
		cmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildId, v)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", v.Name, err)
			return dg, err
		}
		registeredCommands[i] = cmd
	}

	return dg, nil

}

func RemoveCommand(dg *discordgo.Session, guildId string) error {
	log.Println("Removing commands...")
	// // We need to fetch the commands, since deleting requires the command ID.
	// // We are doing this from the returned commands on line 375, because using
	// // this will delete all the commands, which might not be desirable, so we
	// // are deleting only the commands that we added.
	// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)
	// if err != nil {
	// 	log.Fatalf("Could not fetch registered commands: %v", err)
	// }

	for _, v := range registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, guildId, v.ID)
		if err != nil {
			log.Fatalf("Cannot delete '%v' command: %v", v.Name, err)
			return err
		}
	}
	return nil
}

func buildReligionRoleOptionChoices() []*discordgo.ApplicationCommandOptionChoice {

	c := []*discordgo.ApplicationCommandOptionChoice{
		{
			Name:  string(config.LatinCatholic),
			Value: config.LatinCatholic,
		},
		{
			Name:  string(config.EasternCatholic),
			Value: config.EasternCatholic,
		}, {
			Name:  string(config.OrthodoxChristian),
			Value: config.OrthodoxChristian,
		}, {
			Name:  string(config.RCIACatechumen),
			Value: config.RCIACatechumen,
		}, {
			Name:  string(config.Protestant),
			Value: config.Protestant,
		}, {
			Name:  string(config.NonCatholic),
			Value: config.NonCatholic,
		}, {
			Name:  string(config.Atheist),
			Value: config.Atheist,
		},
	}

	return c
}
