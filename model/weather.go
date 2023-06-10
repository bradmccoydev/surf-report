package model

import "time"

// weather/point endpoint
type Weather struct {
	Hours []struct {
		AirTemperature struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"airTemperature"`
		CloudCover struct {
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"cloudCover"`
		SwellDirection struct {
			Icon float64 `json:"icon"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellDirection"`
		SwellHeight struct {
			Icon float64 `json:"icon"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellHeight"`
		SwellPeriod struct {
			Icon float64 `json:"icon"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"swellPeriod"`
		Time             time.Time `json:"time"`
		WaterTemperature struct {
			Meto float64 `json:"meto"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waterTemperature"`
		WaveDirection struct {
			Icon float64 `json:"icon"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waveDirection"`
		WaveHeight struct {
			Icon float64 `json:"icon"`
			Noaa float64 `json:"noaa"`
			Sg   float64 `json:"sg"`
		} `json:"waveHeight"`
	} `json:"hours"`
	Meta struct {
		Cost         int      `json:"cost"`
		DailyQuota   int      `json:"dailyQuota"`
		End          string   `json:"end"`
		Lat          float64  `json:"lat"`
		Lng          float64  `json:"lng"`
		Params       []string `json:"params"`
		RequestCount int      `json:"requestCount"`
		Start        string   `json:"start"`
	} `json:"meta"`
}
