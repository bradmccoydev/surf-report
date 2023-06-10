package model

import "time"

// tide endpoint
type Tide struct {
	Data []struct {
		Sg   float64   `json:"sg"`
		Time time.Time `json:"time"`
	} `json:"data"`
	Meta struct {
		Cost         int     `json:"cost"`
		DailyQuota   int     `json:"dailyQuota"`
		Datum        string  `json:"datum"`
		End          string  `json:"end"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
		RequestCount int     `json:"requestCount"`
		Start        string  `json:"start"`
		Station      struct {
			Distance int     `json:"distance"`
			Lat      float64 `json:"lat"`
			Lng      float64 `json:"lng"`
			Name     string  `json:"name"`
			Source   string  `json:"source"`
		} `json:"station"`
	} `json:"meta"`
}
