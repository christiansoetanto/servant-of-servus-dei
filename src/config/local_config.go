package config

const (
	LocalServer              = "Local Server"
	LocalServerConfigGuildID = "813302330782253066"

	LocalServerConfigGeneralDiscussionChannelID     = "813302330782253069"
	LocalServerConfigReactionRolesChannelID         = "941213323444244501"
	LocalServerConfigServerInformationChannelID     = "848858055944306698"
	LocalServerConfigReligiousQuestionsChannelID    = ""
	LocalServerConfigReligiousDiscussions1ChannelID = ""
	LocalServerConfigReligiousDiscussions2ChannelID = ""
	LocalServerConfigAnsweredQuestionsChannelID     = ""
	LocalServerConfigFAQChannelID                   = ""
	LocalServerConfigResponsesChannelId             = ""
	LocalServerConfigVettingQuestioningChannelId    = ""
	LocalServerConfigRulesVettingChannelId          = ""

	LocalServerConfigVettingRoleID            = "974632148952809482"
	LocalServerConfigVettingQuestioningRoleID = "974632188823863296"
	LocalServerConfigApprovedUserRoleID       = "974632216304943155"

	LocalServerConfigLatinCatholicReligionRoleId     ReligionRoleId = "974630535395680337"
	LocalServerConfigEasternCatholicReligionRoleId   ReligionRoleId = "974667212587671613"
	LocalServerConfigOrthodoxChristianReligionRoleId ReligionRoleId = "974667248826449950"
	LocalServerConfigRCIACatechumenReligionRoleId    ReligionRoleId = "974667251498225704"
	LocalServerConfigProtestantReligionRoleId        ReligionRoleId = "974667253045919784"
	LocalServerConfigNonCatholicReligionRoleId       ReligionRoleId = "974667254627201084"
	LocalServerConfigAtheistReligionRoleId           ReligionRoleId = "974667257122795570"

	LocalServerConfigUpvoteReactionID = ""
	LocalServerConfigSinReactionID    = ""
)

const (
	LocalServerConfigAcknowledgementMessageFormat = "Verification of user <@%s> with role <@&%s> is successful.\nThank you for using my service. Beep. Boop.\n"
	LocalServerConfigWelcomeMessageEmbedFormat    = "Welcome to Servus Dei, <@%s>! We are happy to have you! Make sure you check out <#%s> to gain access to the various channels we offer and please do visit <#%s> so you can understand our server better and take use of everything we have to offer. God Bless!"
	LocalServerConfigMissedQuestionOneFormat      = "Hey <@%s>! It looks like you missed question 1. Please re-read the <#%s> again, we assure you that the code is in there. Thank you for your understanding.\\nPS: if you are sure you got it right, please ignore this message."
	LocalServerConfigWelcomeTitle                 = "Welcome to Servus Dei!"
	LocalServerConfigWelcomeMessageFormat         = "Hey %s! %s just approved your vetting response. Welcome to the server. Feel free to tag us had you have further questions. Enjoy!"
)
