package util

import "fmt"

const (
	servusDeiConfigGeneralDiscussionChannelID = "751174152588623912"
	servusDeiConfigReactionRolesChannelID     = "767452241321000970"
	servusDeiConfigServerInformationChannelID = "973586981789499452"

	servusDeiConfigVettingRoleID            = "751145124834312342"
	servusDeiConfigVettingQuestioningRoleID = "914986915030241301"
	servusDeiConfigApprovedUserRoleID       = "751144797938384979"
	servusDeiConfigGuildID                  = "751139261515825162"

	servusDeiConfigLatinCatholic     ReligionRoleID = "751145824532168775"
	servusDeiConfigEasternCatholic   ReligionRoleID = "751148911267414067"
	servusDeiConfigOrthodoxChristian ReligionRoleID = "751148354716565656"
	servusDeiConfigRCIACatechumen    ReligionRoleID = "751196794771472395"
	servusDeiConfigProtestant        ReligionRoleID = "751145951137103872"
	servusDeiConfigNonCatholic       ReligionRoleID = "751146099351224382"
	servusDeiConfigAtheist           ReligionRoleID = "751148904938209351"
)

func ApplyServusDeiServerConfig() {
	fmt.Println("Applying Servus Dei server config...")

	GeneralDiscussionChannelID = servusDeiConfigGeneralDiscussionChannelID
	ReactionRolesChannelID = servusDeiConfigReactionRolesChannelID
	ServerInformationChannelID = servusDeiConfigServerInformationChannelID
	VettingRoleID = servusDeiConfigVettingRoleID
	VettingQuestioningRoleID = servusDeiConfigVettingQuestioningRoleID
	ApprovedUserRoleID = servusDeiConfigApprovedUserRoleID
	GuildID = servusDeiConfigGuildID

	ReligionRoleMapping[LatinCatholic] = servusDeiConfigLatinCatholic
	ReligionRoleMapping[EasternCatholic] = servusDeiConfigEasternCatholic
	ReligionRoleMapping[OrthodoxChristian] = servusDeiConfigOrthodoxChristian
	ReligionRoleMapping[RCIACatechumen] = servusDeiConfigRCIACatechumen
	ReligionRoleMapping[Protestant] = servusDeiConfigProtestant
	ReligionRoleMapping[NonCatholic] = servusDeiConfigNonCatholic
	ReligionRoleMapping[Atheist] = servusDeiConfigAtheist
}
