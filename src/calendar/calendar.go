package calendar

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/christiansoetanto/servant-of-servus-dei/src/config"
	"github.com/christiansoetanto/servant-of-servus-dei/src/util"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type AllLiturgicalDays struct {
	LiturgicalDaysEn []LiturgicalDay
	LiturgicalDaysLa []LiturgicalDay
}
type LiturgicalDay struct {
	Key                   string        `json:"key"`
	Date                  string        `json:"date"`
	Precedence            string        `json:"precedence"`
	Rank                  string        `json:"rank"`
	IsHolyDayOfObligation bool          `json:"isHolyDayOfObligation"`
	IsOptional            bool          `json:"isOptional"`
	Martyrology           []Martyrology `json:"martyrology"`
	Titles                []string      `json:"titles"`
	Calendar              Calendar      `json:"calendar"`
	Cycles                Cycles        `json:"cycles"`
	Name                  string        `json:"name"`
	RankName              string        `json:"rankName"`
	ColorName             []string      `json:"colorName"`
	SeasonNames           []string      `json:"seasonNames"`
}
type Calendar struct {
	WeekOfSeason          int    `json:"weekOfSeason,omitempty"`
	DayOfSeason           int    `json:"dayOfSeason,omitempty"`
	DayOfWeek             int    `json:"dayOfWeek,omitempty"`
	NthDayOfWeekInMonth   int    `json:"nthDayOfWeekInMonth,omitempty"`
	StartOfSeason         string `json:"startOfSeason,omitempty"`
	EndOfSeason           string `json:"endOfSeason,omitempty"`
	StartOfLiturgicalYear string `json:"startOfLiturgicalYear,omitempty"`
	EndOfLiturgicalYear   string `json:"endOfLiturgicalYear,omitempty"`
}
type Cycles struct {
	ProperCycle  string `json:"properCycle"`
	SundayCycle  string `json:"sundayCycle"`
	WeekdayCycle string `json:"weekdayCycle"`
	PsalterWeek  string `json:"psalterWeek"`
}
type Martyrology struct {
	Key               string   `json:"key"`
	CanonizationLevel string   `json:"canonizationLevel"`
	DateOfDeath       int      `json:"dateOfDeath"`
	Titles            []string `json:"titles,omitempty"`
}

type Messages struct {
	Messages []MessageItem `json:"messages"`
}
type MessageItem struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

const (
	LiturgicalCalendarCron    = "Liturgical Calendar Cron"
	LiturgicalCalendarCommand = "Liturgical Calendar Command"
)

func getText(liturgicalDays []LiturgicalDay) string {
	var text string
	for _, day := range liturgicalDays {

		text += "• "
		//[day, date] //if memorial/feast/solemnity [rank] [name] in [seasonName] season.
		rank, rankName, isHolyDayOfObligation, name, seasonNames := strings.ToLower(day.Rank), day.RankName,
			day.IsHolyDayOfObligation, day.Name, day.SeasonNames
		if rank == "memorial" || rank == "feast" || rank == "solemnity" {
			text += fmt.Sprintf("%s of %s", strings.Title(rankName), name)
			if len(seasonNames) > 0 {
				text += fmt.Sprintf(" in the %s", seasonNames[0])
			}
		} else {
			text += name
		}

		if isHolyDayOfObligation {
			text += ". A Holy Day of Obligation"
		}

		text += ".\n"
	}
	return text

}

func CalendarCronJob(s *discordgo.Session) func() {
	return func() {
		textToSend, msg := GetCalendarText()
		if msg != "" {
			util.ReportError(s, msg)
			return
		}
		_, err := s.ChannelMessageSend(config.ServusDeiConfigLiturgicalCalendarDiscussionsChannelId, textToSend)
		if err != nil {
			msg := fmt.Sprintf("Error: %s : %s", LiturgicalCalendarCron, err.Error())
			util.ReportError(s, msg)
			return
		}
		log.Printf("%s : %s", LiturgicalCalendarCron, textToSend)
	}
}

func GetCalendarText() (string, string) {
	functionsUrl := os.Getenv("ROMCAL_API_FUNCTIONS_URL")
	response, err := http.Get(functionsUrl)
	if err != nil {
		return "", fmt.Sprintf("Error: [Fetching calendar] : %s", err.Error())
	}
	data, _ := ioutil.ReadAll(response.Body)

	var allLiturgicalDays AllLiturgicalDays
	errUnmarshal := json.Unmarshal(data, &allLiturgicalDays)
	if errUnmarshal != nil {
		return "", fmt.Sprintf("Error: [Unmarshall Response] : %s", errUnmarshal.Error())
	}

	currentTime := time.Now()
	greetingText := fmt.Sprintf(
		"Hello! Today is %s, %d %s %d UTC time.\nThe Roman Catholic Church is celebrating:", currentTime.Weekday(),
		currentTime.Day(), currentTime.Month(), currentTime.Year())
	calendarText := getText(allLiturgicalDays.LiturgicalDaysEn)
	return fmt.Sprintf("%s\n%s\n", greetingText, calendarText), ""
}
