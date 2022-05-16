package config

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
	LatinCatholic      ReligionRoleId
	EasternCatholic    ReligionRoleId
	OrthodoxChristian  ReligionRoleId
	RCIACatechumen     ReligionRoleId
	Protestant         ReligionRoleId
	NonCatholic        ReligionRoleId
	Atheist            ReligionRoleId
}
type Reaction struct {
	Upvote string
	Sin    string
}

type Wording struct {
	AcknowledgementMessageFormat string
	WelcomeMessageEmbedFormat    string
	WelcomeMessageFormat         string
	MissedQuestionOneFormat      string
	WelcomeTitle                 string
}
type ReligionRoleMapping struct {
	LatinCatholic     ReligionRoleId
	EasternCatholic   ReligionRoleId
	OrthodoxChristian ReligionRoleId
	RCIACatechumen    ReligionRoleId
	Protestant        ReligionRoleId
	NonCatholic       ReligionRoleId
	Atheist           ReligionRoleId
}
type GuildConfig struct {
	GuildName                 string
	Channel                   Channel
	Role                      Role
	Reaction                  Reaction
	ReligionRoleMappingStruct ReligionRoleMapping
	ReligionRoleMappingMap    map[ReligionRoleType]ReligionRoleId
	Wording                   Wording
}

var Config = map[string]GuildConfig{
	ServusDeiConfigGuildID: {
		GuildName: ServusDei,
		Channel: Channel{
			GeneralDiscussion:     ServusDeiConfigGeneralDiscussionChannelID,
			ReactionRoles:         ServusDeiConfigReactionRolesChannelID,
			ServerInformation:     ServusDeiConfigServerInformationChannelID,
			ReligiousQuestions:    ServusDeiConfigReligiousQuestionsChannelID,
			ReligiousDiscussions1: ServusDeiConfigReligiousDiscussions1ChannelID,
			ReligiousDiscussions2: ServusDeiConfigReligiousDiscussions2ChannelID,
			AnsweredQuestions:     ServusDeiConfigAnsweredQuestionsChannelID,
			FAQ:                   ServusDeiConfigFAQChannelID,
			Responses:             ServusDeiConfigResponsesChannelId,
			VettingQuestioning:    ServusDeiConfigVettingQuestioningChannelId,
			RulesVetting:          ServusDeiConfigRulesVettingChannelId,
		},
		Role: Role{
			Vetting:            ServusDeiConfigVettingRoleID,
			VettingQuestioning: ServusDeiConfigVettingQuestioningRoleID,
			ApprovedUser:       ServusDeiConfigApprovedUserRoleID,
			LatinCatholic:      ServusDeiConfigLatinCatholicReligionRoleId,
			EasternCatholic:    ServusDeiConfigEasternCatholicReligionRoleId,
			OrthodoxChristian:  ServusDeiConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:     ServusDeiConfigRCIACatechumenReligionRoleId,
			Protestant:         ServusDeiConfigProtestantReligionRoleId,
			NonCatholic:        ServusDeiConfigNonCatholicReligionRoleId,
			Atheist:            ServusDeiConfigAtheistReligionRoleId,
		},
		Reaction: Reaction{
			Upvote: ServusDeiConfigUpvoteReactionID,
			Sin:    ServusDeiConfigSinReactionID,
		},
		ReligionRoleMappingStruct: ReligionRoleMapping{
			LatinCatholic:     ServusDeiConfigLatinCatholicReligionRoleId,
			EasternCatholic:   ServusDeiConfigEasternCatholicReligionRoleId,
			OrthodoxChristian: ServusDeiConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:    ServusDeiConfigRCIACatechumenReligionRoleId,
			Protestant:        ServusDeiConfigProtestantReligionRoleId,
			NonCatholic:       ServusDeiConfigNonCatholicReligionRoleId,
			Atheist:           ServusDeiConfigAtheistReligionRoleId,
		},
		ReligionRoleMappingMap: map[ReligionRoleType]ReligionRoleId{
			LatinCatholic:     ServusDeiConfigLatinCatholicReligionRoleId,
			EasternCatholic:   ServusDeiConfigEasternCatholicReligionRoleId,
			OrthodoxChristian: ServusDeiConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:    ServusDeiConfigRCIACatechumenReligionRoleId,
			Protestant:        ServusDeiConfigProtestantReligionRoleId,
			NonCatholic:       ServusDeiConfigNonCatholicReligionRoleId,
			Atheist:           ServusDeiConfigAtheistReligionRoleId,
		},
		Wording: Wording{
			AcknowledgementMessageFormat: ServusDeiConfigAcknowledgementMessageFormat,
			WelcomeMessageEmbedFormat:    ServusDeiConfigWelcomeMessageEmbedFormat,
			MissedQuestionOneFormat:      ServusDeiConfigMissedQuestionOneFormat,
			WelcomeTitle:                 ServusDeiConfigWelcomeTitle,
			WelcomeMessageFormat:         ServusDeiConfigWelcomeMessageFormat,
		},
	},
	LocalServerConfigGuildID: {
		GuildName: LocalServer,
		Channel: Channel{
			GeneralDiscussion:     LocalServerConfigGeneralDiscussionChannelID,
			ReactionRoles:         LocalServerConfigReactionRolesChannelID,
			ServerInformation:     LocalServerConfigServerInformationChannelID,
			ReligiousQuestions:    LocalServerConfigReligiousQuestionsChannelID,
			ReligiousDiscussions1: LocalServerConfigReligiousDiscussions1ChannelID,
			ReligiousDiscussions2: LocalServerConfigReligiousDiscussions2ChannelID,
			AnsweredQuestions:     LocalServerConfigAnsweredQuestionsChannelID,
			FAQ:                   LocalServerConfigFAQChannelID,
			Responses:             LocalServerConfigResponsesChannelId,
			VettingQuestioning:    LocalServerConfigVettingQuestioningChannelId,
			RulesVetting:          LocalServerConfigRulesVettingChannelId,
		},
		Role: Role{
			Vetting:            LocalServerConfigVettingRoleID,
			VettingQuestioning: LocalServerConfigVettingQuestioningRoleID,
			ApprovedUser:       LocalServerConfigApprovedUserRoleID,
			LatinCatholic:      LocalServerConfigLatinCatholicReligionRoleId,
			EasternCatholic:    LocalServerConfigEasternCatholicReligionRoleId,
			OrthodoxChristian:  LocalServerConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:     LocalServerConfigRCIACatechumenReligionRoleId,
			Protestant:         LocalServerConfigProtestantReligionRoleId,
			NonCatholic:        LocalServerConfigNonCatholicReligionRoleId,
			Atheist:            LocalServerConfigAtheistReligionRoleId,
		},
		Reaction: Reaction{
			Upvote: LocalServerConfigUpvoteReactionID,
			Sin:    LocalServerConfigSinReactionID,
		},
		ReligionRoleMappingStruct: ReligionRoleMapping{
			LatinCatholic:     LocalServerConfigLatinCatholicReligionRoleId,
			EasternCatholic:   LocalServerConfigEasternCatholicReligionRoleId,
			OrthodoxChristian: LocalServerConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:    LocalServerConfigRCIACatechumenReligionRoleId,
			Protestant:        LocalServerConfigProtestantReligionRoleId,
			NonCatholic:       LocalServerConfigNonCatholicReligionRoleId,
			Atheist:           LocalServerConfigAtheistReligionRoleId,
		},
		ReligionRoleMappingMap: map[ReligionRoleType]ReligionRoleId{
			LatinCatholic:     LocalServerConfigLatinCatholicReligionRoleId,
			EasternCatholic:   LocalServerConfigEasternCatholicReligionRoleId,
			OrthodoxChristian: LocalServerConfigOrthodoxChristianReligionRoleId,
			RCIACatechumen:    LocalServerConfigRCIACatechumenReligionRoleId,
			Protestant:        LocalServerConfigProtestantReligionRoleId,
			NonCatholic:       LocalServerConfigNonCatholicReligionRoleId,
			Atheist:           LocalServerConfigAtheistReligionRoleId,
		},
		Wording: Wording{
			AcknowledgementMessageFormat: LocalServerConfigAcknowledgementMessageFormat,
			WelcomeMessageEmbedFormat:    LocalServerConfigWelcomeMessageEmbedFormat,
			MissedQuestionOneFormat:      LocalServerConfigMissedQuestionOneFormat,
			WelcomeTitle:                 LocalServerConfigWelcomeTitle,
			WelcomeMessageFormat:         LocalServerConfigWelcomeMessageFormat,
		},
	},
}