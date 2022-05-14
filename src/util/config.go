package util

var (
	GeneralDiscussionChannelID string
	ReactionRolesChannelID     string
	ServerInformationChannelID string
	VettingRoleID              string
	VettingQuestioningRoleID   string
	ApprovedUserRoleID         string

	GuildID string
)
var ReligionRoleMapping = make(map[ReligionRoleType]ReligionRoleID)

const (
	LatinCatholic     ReligionRoleType = "Latin Catholic"
	EasternCatholic   ReligionRoleType = "Eastern Catholic"
	OrthodoxChristian ReligionRoleType = "Orthodox Christian"
	RCIACatechumen    ReligionRoleType = "RCIA / Catechumen"
	Protestant        ReligionRoleType = "Protestant"
	NonCatholic       ReligionRoleType = "Non Catholic"
	Atheist           ReligionRoleType = "Atheist"
)

type ReligionRoleType string
type ReligionRoleID string
