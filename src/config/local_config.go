package util

import "fmt"

const (
	localConfigGeneralDiscussionChannelID = "813302330782253069"
	localConfigReactionRolesChannelID     = "941213323444244501"
	localConfigServerInformationChannelID = "848858055944306698"

	localConfigVettingRoleID            = "974632148952809482"
	localConfigVettingQuestioningRoleID = "974632188823863296"
	localConfigApprovedUserRoleID       = "974632216304943155"
	localConfigGuildID                  = "813302330782253066"

	localConfigLatinCatholic     ReligionRoleId = "974630535395680337"
	localConfigEasternCatholic   ReligionRoleId = "974667212587671613"
	localConfigOrthodoxChristian ReligionRoleId = "974667248826449950"
	localConfigRCIACatechumen    ReligionRoleId = "974667251498225704"
	localConfigProtestant        ReligionRoleId = "974667253045919784"
	localConfigNonCatholic       ReligionRoleId = "974667254627201084"
	localConfigAtheist           ReligionRoleId = "974667257122795570"

	localConfigUpvoteReactionID = ""
)

func ApplyLocalServerConfig() {
	fmt.Println("Applying local server config...")
	GeneralDiscussionChannelId = localConfigGeneralDiscussionChannelID
	ReactionRolesChannelId = localConfigReactionRolesChannelID
	ServerInformationChannelId = localConfigServerInformationChannelID
	VettingRoleId = localConfigVettingRoleID
	VettingQuestioningRoleId = localConfigVettingQuestioningRoleID
	ApprovedUserRoleId = localConfigApprovedUserRoleID
	GuildID = localConfigGuildID

	ReligionRoleMapping[LatinCatholic] = localConfigLatinCatholic
	ReligionRoleMapping[EasternCatholic] = localConfigEasternCatholic
	ReligionRoleMapping[OrthodoxChristian] = localConfigOrthodoxChristian
	ReligionRoleMapping[RCIACatechumen] = localConfigRCIACatechumen
	ReligionRoleMapping[Protestant] = localConfigProtestant
	ReligionRoleMapping[NonCatholic] = localConfigNonCatholic
	ReligionRoleMapping[Atheist] = localConfigAtheist

	UpvoteReactionId = localConfigUpvoteReactionID

}
