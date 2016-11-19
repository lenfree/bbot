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

type Game struct {
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
}

type Games struct {
        Date  string `json:"date"`
        League struct {
                Alias string `json:"alias"`
                ID    string `json:"id"`
                Name  string `json:"name"`
        } `json:"league"`
        Games []Game `json:"games"`
}

type Summary struct {
        Away struct {
                ID      string `json:"id"`
                Market  string `json:"market"`
                Name    string `json:"name"`
                Players []struct {
                        Active          bool   `json:"active"`
                        FirstName       string `json:"first_name"`
                        FullName        string `json:"full_name"`
                        ID              string `json:"id"`
                        JerseyNumber    string `json:"jersey_number"`
                        LastName        string `json:"last_name"`
                        Played          bool   `json:"played"`
                        Position        string `json:"position"`
                        PrimaryPosition string `json:"primary_position"`
                        Starter         bool   `json:"starter"`
                        Statistics      struct {
                                Assists              int    `json:"assists"`
                                AssistsTurnoverRatio int    `json:"assists_turnover_ratio"`
                                BlockedAtt           int    `json:"blocked_att"`
                                Blocks               int    `json:"blocks"`
                                DefensiveRebounds    int    `json:"defensive_rebounds"`
                                FieldGoalsAtt        int    `json:"field_goals_att"`
                                FieldGoalsMade       int    `json:"field_goals_made"`
                                FieldGoalsPct        int    `json:"field_goals_pct"`
                                FlagrantFouls        int    `json:"flagrant_fouls"`
                                FreeThrowsAtt        int    `json:"free_throws_att"`
                                FreeThrowsMade       int    `json:"free_throws_made"`
                                FreeThrowsPct        int    `json:"free_throws_pct"`
                                Minutes              string `json:"minutes"`
                                OffensiveRebounds    int    `json:"offensive_rebounds"`
                                PersonalFouls        int    `json:"personal_fouls"`
                                PlsMin               int    `json:"pls_min"`
                                Points               int    `json:"points"`
                                Rebounds             int    `json:"rebounds"`
                                Steals               int    `json:"steals"`
                                TechFouls            int    `json:"tech_fouls"`
                                ThreePointsAtt       int    `json:"three_points_att"`
                                ThreePointsMade      int    `json:"three_points_made"`
                                ThreePointsPct       int    `json:"three_points_pct"`
                                Turnovers            int    `json:"turnovers"`
                                TwoPointsAtt         int    `json:"two_points_att"`
                                TwoPointsMade        int    `json:"two_points_made"`
                                TwoPointsPct         int    `json:"two_points_pct"`
                        } `json:"statistics"`
                } `json:"players"`
                Points  int `json:"points"`
                Scoring []struct {
                        Number   int    `json:"number"`
                        Points   int    `json:"points"`
                        Sequence int    `json:"sequence"`
                        Type     string `json:"type"`
                } `json:"scoring"`
                Statistics struct {
                        Assists              int     `json:"assists"`
                        AssistsTurnoverRatio float64 `json:"assists_turnover_ratio"`
                        BlockedAtt           int     `json:"blocked_att"`
                        Blocks               int     `json:"blocks"`
                        CoachTechFouls       int     `json:"coach_tech_fouls"`
                        DefensiveRebounds    int     `json:"defensive_rebounds"`
                        Ejections            int     `json:"ejections"`
                        FieldGoalsAtt        int     `json:"field_goals_att"`
                        FieldGoalsMade       int     `json:"field_goals_made"`
                        FieldGoalsPct        int     `json:"field_goals_pct"`
                        FlagrantFouls        int     `json:"flagrant_fouls"`
                        Foulouts             int     `json:"foulouts"`
                        FreeThrowsAtt        int     `json:"free_throws_att"`
                        FreeThrowsMade       int     `json:"free_throws_made"`
                        FreeThrowsPct        float64 `json:"free_throws_pct"`
                        Minutes              string  `json:"minutes"`
                        OffensiveRebounds    int     `json:"offensive_rebounds"`
                        PersonalFouls        int     `json:"personal_fouls"`
                        PlayerTechFouls      int     `json:"player_tech_fouls"`
                        Points               int     `json:"points"`
                        Rebounds             int     `json:"rebounds"`
                        Steals               int     `json:"steals"`
                        TeamRebounds         int     `json:"team_rebounds"`
                        TeamTechFouls        int     `json:"team_tech_fouls"`
                        TeamTurnovers        int     `json:"team_turnovers"`
                        ThreePointsAtt       int     `json:"three_points_att"`
                        ThreePointsMade      int     `json:"three_points_made"`
                        ThreePointsPct       int     `json:"three_points_pct"`
                        Turnovers            int     `json:"turnovers"`
                        TwoPointsAtt         int     `json:"two_points_att"`
                        TwoPointsMade        int     `json:"two_points_made"`
                        TwoPointsPct         float64 `json:"two_points_pct"`
                } `json:"statistics"`
        } `json:"away"`
        Clock    string `json:"clock"`
        Coverage string `json:"coverage"`
        Home     struct {
                ID      string `json:"id"`
                Market  string `json:"market"`
                Name    string `json:"name"`
                Players []struct {
                        Active          bool   `json:"active"`
                        FirstName       string `json:"first_name"`
                        FullName        string `json:"full_name"`
                        ID              string `json:"id"`
                        JerseyNumber    string `json:"jersey_number"`
                        LastName        string `json:"last_name"`
                        Played          bool   `json:"played"`
                        Position        string `json:"position"`
                        PrimaryPosition string `json:"primary_position"`
                        Statistics      struct {
                                Assists              int     `json:"assists"`
                                AssistsTurnoverRatio int     `json:"assists_turnover_ratio"`
                                BlockedAtt           int     `json:"blocked_att"`
                                Blocks               int     `json:"blocks"`
                                DefensiveRebounds    int     `json:"defensive_rebounds"`
                                FieldGoalsAtt        int     `json:"field_goals_att"`
                                FieldGoalsMade       int     `json:"field_goals_made"`
                                FieldGoalsPct        float64 `json:"field_goals_pct"`
                                FlagrantFouls        int     `json:"flagrant_fouls"`
                                FreeThrowsAtt        int     `json:"free_throws_att"`
                                FreeThrowsMade       int     `json:"free_throws_made"`
                                FreeThrowsPct        int     `json:"free_throws_pct"`
                                Minutes              string  `json:"minutes"`
                                OffensiveRebounds    int     `json:"offensive_rebounds"`
                                PersonalFouls        int     `json:"personal_fouls"`
                                PlsMin               int     `json:"pls_min"`
                                Points               int     `json:"points"`
                                Rebounds             int     `json:"rebounds"`
                                Steals               int     `json:"steals"`
                                TechFouls            int     `json:"tech_fouls"`
                                ThreePointsAtt       int     `json:"three_points_att"`
                                ThreePointsMade      int     `json:"three_points_made"`
                                ThreePointsPct       float64 `json:"three_points_pct"`
                                Turnovers            int     `json:"turnovers"`
                                TwoPointsAtt         int     `json:"two_points_att"`
                                TwoPointsMade        int     `json:"two_points_made"`
                                TwoPointsPct         int     `json:"two_points_pct"`
                        } `json:"statistics"`
                } `json:"players"`
                Points  int `json:"points"`
                Scoring []struct {
                        Number   int    `json:"number"`
                        Points   int    `json:"points"`
                        Sequence int    `json:"sequence"`
                        Type     string `json:"type"`
                } `json:"scoring"`
                Statistics struct {
                        Assists              int     `json:"assists"`
                        AssistsTurnoverRatio float64 `json:"assists_turnover_ratio"`
                        BlockedAtt           int     `json:"blocked_att"`
                        Blocks               int     `json:"blocks"`
                        CoachTechFouls       int     `json:"coach_tech_fouls"`
                        DefensiveRebounds    int     `json:"defensive_rebounds"`
                        Ejections            int     `json:"ejections"`
                        FieldGoalsAtt        int     `json:"field_goals_att"`
                        FieldGoalsMade       int     `json:"field_goals_made"`
                        FieldGoalsPct        float64 `json:"field_goals_pct"`
                        FlagrantFouls        int     `json:"flagrant_fouls"`
                        Foulouts             int     `json:"foulouts"`
                        FreeThrowsAtt        int     `json:"free_throws_att"`
                        FreeThrowsMade       int     `json:"free_throws_made"`
                        FreeThrowsPct        float64 `json:"free_throws_pct"`
                        Minutes              string  `json:"minutes"`
                        OffensiveRebounds    int     `json:"offensive_rebounds"`
                        PersonalFouls        int     `json:"personal_fouls"`
                        PlayerTechFouls      int     `json:"player_tech_fouls"`
                        Points               int     `json:"points"`
                        Rebounds             int     `json:"rebounds"`
                        Steals               int     `json:"steals"`
                        TeamRebounds         int     `json:"team_rebounds"`
                        TeamTechFouls        int     `json:"team_tech_fouls"`
                        TeamTurnovers        int     `json:"team_turnovers"`
                        ThreePointsAtt       int     `json:"three_points_att"`
                        ThreePointsMade      int     `json:"three_points_made"`
                        ThreePointsPct       float64 `json:"three_points_pct"`
                        Turnovers            int     `json:"turnovers"`
                        TwoPointsAtt         int     `json:"two_points_att"`
                        TwoPointsMade        int     `json:"two_points_made"`
                        TwoPointsPct         float64 `json:"two_points_pct"`
                } `json:"statistics"`
        } `json:"home"`
        ID          string `json:"id"`
        NeutralSite bool   `json:"neutral_site"`
        Officials   []struct {
                Assignment string `json:"assignment"`
                Experience string `json:"experience"`
                FirstName  string `json:"first_name"`
                FullName   string `json:"full_name"`
                ID         string `json:"id"`
                LastName   string `json:"last_name"`
                Number     string `json:"number"`
        } `json:"officials"`
        Quarter   int    `json:"quarter"`
        Scheduled string `json:"scheduled"`
        Status    string `json:"status"`
        Venue     struct {
                Address  string `json:"address"`
                Capacity int    `json:"capacity"`
                City     string `json:"city"`
                Country  string `json:"country"`
                ID       string `json:"id"`
                Name     string `json:"name"`
                State    string `json:"state"`
                Zip      string `json:"zip"`
        } `json:"venue"`
}

func NbaScoresHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {

        url := "http://api.sportradar.us/nba-t3/games/" + dateToday() + "/schedule.json?api_key=" + espn
        _, body, err := fasthttp.Get(nil, url)
        if err != nil {
                log.Printf("%s", err.Error())
        }
        var games Games
        json.Unmarshal(body, &games)
        for _, g := range games.Games {
                // Sportsradar API terms and agreement one call per second
                delayMilliSecond(300)
                s := summary(g)
                output := strconv.Itoa(s.Quarter) + " Qtr " + s.Away.Name + " " + strconv.Itoa(s.Away.Points) + " - " + s.Home.Name + " " + strconv.Itoa(s.Home.Points)
                bot.Reply(evt, output, slackbot.WithTyping)
        }
}

func summary(g Game) Summary {
        url := "http://api.sportradar.us/nba-t3/games/" + g.ID + "/summary.json?api_key=" + espn
        _, body, err := fasthttp.Get(nil, url)
        if err != nil {
                log.Printf("%s", err.Error())
        }
        var summary Summary
        json.Unmarshal(body, &summary)
        return summary
}

func dateToday() string {
        return time.Now().UTC().AddDate(0, 0, -1).Format("2006/01/02")
}

func delayMilliSecond(n time.Duration) {
        time.Sleep(n * time.Millisecond)
}
