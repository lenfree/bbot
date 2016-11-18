package main

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
