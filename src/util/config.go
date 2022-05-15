package util

//Channel group
var (
	GeneralDiscussionChannelId     string
	ReactionRolesChannelId         string
	ServerInformationChannelId     string
	ReligiousQuestionsChannelId    string
	ReligiousDiscussions1ChannelId string
	ReligiousDiscussions2ChannelId string
	AnsweredQuestionsChannelId     string
	FAQChannelId                   string
)

//Role group
var (
	VettingRoleId            string
	VettingQuestioningRoleId string
	ApprovedUserRoleId       string
)

//Reaction group
var (
	UpvoteReactionId string
	SinReactionId    string
)

//Misc
var (
	GuildID string
)
var ReligionRoleMapping = make(map[ReligionRoleType]ReligionRoleId)

const (
	LatinCatholic                                    ReligionRoleType = "Latin Catholic"
	EasternCatholic                                  ReligionRoleType = "Eastern Catholic"
	OrthodoxChristian                                ReligionRoleType = "Orthodox Christian"
	RCIACatechumen                                   ReligionRoleType = "RCIA / Catechumen"
	Protestant                                       ReligionRoleType = "Protestant"
	NonCatholic                                      ReligionRoleType = "Non Catholic"
	Atheist                                          ReligionRoleType = "Atheist"
	ReligiousDiscussions1WhiteCheckMarkEmojiName                      = "✅"
	ReligiousDiscussions2BallotBoxWithCheckEmojiName                  = "☑️"
)

type ReligionRoleType string
type ReligionRoleId string

var Moderator = map[ModeratorUserId]ModeratorUsername{
	"255514888041005057": "soetanto",
}

type ModeratorUserId string
type ModeratorUsername string
