package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
	"github.com/valyala/fasthttp"
)

var bot *slackbot.Bot
var espn string

func init() {
        err := godotenv.Load()
         if err != nil {
           log.Fatal("Error loading .env file")
         }
        bot = slackbot.New(os.Getenv("SLACK_TOKEN"))
        espn = os.Getenv("ESPN_TOKEN")
}

func main() {
	toMe := bot.Messages(slackbot.Mention, slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Hear("(?i)how are you(.*)").MessageHandler(HowAreYouHandler)
	//bot.Hear("(?i)(.*)").MessageHandler(HowAreYouHandler)
	bot.Hear("(?)attachment").MessageHandler(AttachmentsHandler)
        bot.Hear("(?i)nba(.*)").MessageHandler(NbaScoresHandler)
	bot.Run()
}

func NbaScoresHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
        current_date := time.Now().UTC().Format("2006/01/02")

        url := "http://api.sportradar.us/nba-t3/games/" + current_date + "/schedule.json?api_key=" + espn
        _, body, err := fasthttp.Get(nil, url)
        if err != nil {
                log.Printf("%s", err.Error())
        }
        var games Schedule
        json.Unmarshal(body, &games)
        for _, g := range games.Games {
                var homePoints int
                var awayPoints int
                if g.AwayPoints == 0 {
                        awayPoints = 0
                }
                if g.HomePoints == 0 {
                        homePoints = 0
                }
                output := g.Away.Name + " "+ strconv.Itoa(awayPoints) + " - " + g.Home.Name + " " + strconv.Itoa(homePoints)
                bot.Reply(evt, output, slackbot.WithTyping)
        }
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
        log.Println(evt)
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "A bit tired, " + evt.User + ". You get it? A bit?", slackbot.WithTyping)
        log.Printf("%+#v", evt)
        log.Printf("%+#v", bot)
        log.Printf("%+#v", evt.User)
}

func AttachmentsHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	txt := "Beep Beep Boop is a ridiculously simple hosting platform for your Slackbots."
	attachment := slack.Attachment{
		Pretext:   "We bring bots to life. :sunglasses: :thumbsup:",
		Title:     "Host, deploy and share your bot in seconds.",
		TitleLink: "https://beepboophq.com/",
		Text:      txt,
		Fallback:  txt,
		ImageURL:  "https://storage.googleapis.com/beepboophq/_assets/bot-1.22f6fb.png",
		Color:     "#7CD197",
	}

	// supports multiple attachments
	attachments := []slack.Attachment{attachment}

	bot.ReplyWithAttachments(evt, attachments, slackbot.WithTyping)
}
