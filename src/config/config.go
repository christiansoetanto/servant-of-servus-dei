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
	ResponsesChannelId             string
	VettingQuestioningChannelId    string
	RulesVettingChannelId          string
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

type Channel struct {
	GeneralDiscussion     string
	ReactionRoles         string
	ServerInformation     string
	ReligiousQuestions    string
	ReligiousDiscussions1 string
	ReligiousDiscussions2 string
	AnsweredQuestions     string
	FAQ                   string
	Responses             string
	VettingQuestioning    string
	RulesVetting          string
}
type Role struct {
	Vetting            string
	VettingQuestioning string
	ApprovedUser       string
}
type Reaction struct {
	Upvote string
	Sin    string
}

type GuildConfig struct {
}

type GuildId string

type Config map[GuildId]GuildConfig
