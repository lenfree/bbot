package main

import (
	"log"
	"os"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/joho/godotenv"
	"github.com/nlopes/slack"
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

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	log.Println(evt)
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
}

func HowAreYouHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "A bit tired, "+evt.User+". You get it? A bit?", slackbot.WithTyping)
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
