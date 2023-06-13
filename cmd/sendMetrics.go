package cmd

import (
	"fmt"
	"time"

	"github.com/bradmccoydev/surf-report/pkg/prometheus"
	surf "github.com/bradmccoydev/surf-report/pkg/surf"
	"github.com/spf13/cobra"
)

var (
	hasItems     bool
	body         string
	prNumber     string
	repoFullUrl  string
	fileName     string
	slackWebhook string

	sendMetricsCmd = &cobra.Command{
		Use:   "sendmetrics",
		Short: "Send Metrics to Prometheus",
		Long:  `Send Metrics from Surf API to Prometheus`,
		Run: func(cmd *cobra.Command, args []string) {
			err := sendMetrics(args)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(sendMetricsCmd)
	sendMetricsCmd.PersistentFlags().StringVarP(&apiToken, "apiToken", "s", apiToken, "The API Token")
}

func sendMetrics(args []string) error {
	var closestTide float64
	var waterTemperature float64
	var waveDirection float64
	var waveHeight float64
	var airTemperature float64
	var cloudCover float64
	var swellDirection float64
	var swellHeight float64
	var swellPeriod float64

	now := time.Now().UTC()
	closestTime := time.Time{}
	pushGatewayURL := "https://prometheus-pushgateway.dev.basiq.io/metrics/job/basiqio"

	tides := surf.GetTides(apiToken)
	weather := surf.GetWeather(apiToken)

	for _, tide := range tides.Data {
		if closestTime.IsZero() || absDuration(now.Sub(tide.Time)) < absDuration(now.Sub(closestTime)) {
			closestTime = tide.Time
			closestTide = tide.Sg
		}
	}

	fmt.Println("Closest Time:", closestTime)
	fmt.Println("Value with Closest Time:", closestTide)

	prometheus.PushGaugeMetric(pushGatewayURL, "surf_tide", closestTide)

	for _, weather := range weather.Hours {
		if closestTime.IsZero() || absDuration(now.Sub(weather.Time)) < absDuration(now.Sub(closestTime)) {
			closestTime = weather.Time
			waterTemperature = weather.WaterTemperature.Sg
			waveDirection = weather.WaveDirection.Sg
			waveHeight = weather.WaveHeight.Sg
			airTemperature = weather.AirTemperature.Sg
			cloudCover = weather.CloudCover.Sg
			swellDirection = weather.SwellDirection.Sg
			swellHeight = weather.SwellHeight.Sg
			swellPeriod = weather.SwellPeriod.Sg
		}
	}

	prometheus.PushGaugeMetric(pushGatewayURL, "surf_water_temperature", waterTemperature)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_wave_direction", waveDirection)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_wave_height", waveHeight)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_air_temperature", airTemperature)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_cloud_cover", cloudCover)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_swell_direction", swellDirection)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_swell_height", swellHeight)
	prometheus.PushGaugeMetric(pushGatewayURL, "surf_swell_period", swellPeriod)

	return nil
}

func absDuration(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}
	return d
}
