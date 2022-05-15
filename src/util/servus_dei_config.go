package util

import "fmt"

const (
	servusDeiConfigGeneralDiscussionChannelID = "751174152588623912"
	servusDeiConfigReactionRolesChannelID     = "767452241321000970"
	servusDeiConfigServerInformationChannelID = "973586981789499452"
	servusDeiReligiousQuestionsChannelID      = "751174501307383908"
	servusDeiReligiousDiscussions1ChannelID   = "751174442217898065"
	servusDeiReligiousDiscussions2ChannelID   = "771836244879605811"
	servusDeiAnsweredQuestionsChannelID       = "821657995129126942"
	servusDeiFAQChannelID                     = "806007417321291776"
	servusDeiResponsesChannelId               = "751151421231202363"
	servusDeiVettingQuestioningChannelId      = "914987511481249792"
	servusDeiRulesVettingChannelId            = "775654889934159893"

	servusDeiConfigVettingRoleID            = "751145124834312342"
	servusDeiConfigVettingQuestioningRoleID = "914986915030241301"
	servusDeiConfigApprovedUserRoleID       = "751144797938384979"
	servusDeiConfigGuildID                  = "751139261515825162"

	servusDeiConfigLatinCatholic     ReligionRoleId = "751145824532168775"
	servusDeiConfigEasternCatholic   ReligionRoleId = "751148911267414067"
	servusDeiConfigOrthodoxChristian ReligionRoleId = "751148354716565656"
	servusDeiConfigRCIACatechumen    ReligionRoleId = "751196794771472395"
	servusDeiConfigProtestant        ReligionRoleId = "751145951137103872"
	servusDeiConfigNonCatholic       ReligionRoleId = "751146099351224382"
	servusDeiConfigAtheist           ReligionRoleId = "751148904938209351"

	servusDeiUpvoteReactionID = "762045856592822342"
	servusDeiSinReactionID    = "786282687044124712"
)

func ApplyServusDeiServerConfig() {
	fmt.Println("Applying Servus Dei server config...")

	GeneralDiscussionChannelId = servusDeiConfigGeneralDiscussionChannelID
	ReactionRolesChannelId = servusDeiConfigReactionRolesChannelID
	ServerInformationChannelId = servusDeiConfigServerInformationChannelID
	ReligiousQuestionsChannelId = servusDeiReligiousQuestionsChannelID
	ReligiousDiscussions1ChannelId = servusDeiReligiousDiscussions1ChannelID
	ReligiousDiscussions2ChannelId = servusDeiReligiousDiscussions2ChannelID
	AnsweredQuestionsChannelId = servusDeiAnsweredQuestionsChannelID
	FAQChannelId = servusDeiFAQChannelID
	ResponsesChannelId = servusDeiResponsesChannelId
	VettingQuestioningChannelId = servusDeiVettingQuestioningChannelId
	RulesVettingChannelId = servusDeiRulesVettingChannelId

	VettingRoleId = servusDeiConfigVettingRoleID
	VettingQuestioningRoleId = servusDeiConfigVettingQuestioningRoleID
	ApprovedUserRoleId = servusDeiConfigApprovedUserRoleID
	GuildID = servusDeiConfigGuildID

	ReligionRoleMapping[LatinCatholic] = servusDeiConfigLatinCatholic
	ReligionRoleMapping[EasternCatholic] = servusDeiConfigEasternCatholic
	ReligionRoleMapping[OrthodoxChristian] = servusDeiConfigOrthodoxChristian
	ReligionRoleMapping[RCIACatechumen] = servusDeiConfigRCIACatechumen
	ReligionRoleMapping[Protestant] = servusDeiConfigProtestant
	ReligionRoleMapping[NonCatholic] = servusDeiConfigNonCatholic
	ReligionRoleMapping[Atheist] = servusDeiConfigAtheist

	UpvoteReactionId = servusDeiUpvoteReactionID
	SinReactionId = servusDeiSinReactionID

}
