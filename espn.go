package main

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/context"
)

type Schedule struct {
	Date  string `json:"date"`
	Games []struct {
		Away struct {
			Alias string `json:"alias"`
			ID    string `json:"id"`
			Name  string `json:"name"`
		} `json:"away"`
		AwayPoints int `json:"away_points"`
		Broadcast  struct {
			Network   string `json:"network"`
			Satellite string `json:"satellite"`
		} `json:"broadcast"`
		Coverage string `json:"coverage"`
		Home     struct {
			Alias string `json:"alias"`
			ID    string `json:"id"`
			Name  string `json:"name"`
		} `json:"home"`
		HomePoints int    `json:"home_points"`
		ID         string `json:"id"`
		Scheduled  string `json:"scheduled"`
		Status     string `json:"status"`
		Venue      struct {
			Address  string `json:"address"`
			Capacity int    `json:"capacity"`
			City     string `json:"city"`
			Country  string `json:"country"`
			ID       string `json:"id"`
			Name     string `json:"name"`
			State    string `json:"state"`
			Zip      string `json:"zip"`
		} `json:"venue"`
	} `json:"games"`
	League struct {
		Alias string `json:"alias"`
		ID    string `json:"id"`
		Name  string `json:"name"`
	} `json:"league"`
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
		output := g.Away.Name + " " + strconv.Itoa(awayPoints) + " - " + g.Home.Name + " " + strconv.Itoa(homePoints)
		bot.Reply(evt, output, slackbot.WithTyping)
	}
}
