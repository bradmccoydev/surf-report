package cmd

import (
	"fmt"

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

	sendValidationReportToSlackCmd = &cobra.Command{
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
	rootCmd.AddCommand(sendValidationReportToSlackCmd)
	sendValidationReportToSlackCmd.PersistentFlags().StringVarP(&apiToken, "apiToken", "s", apiToken, "The API Token")
}

func sendMetrics(args []string) error {
	tides := surf.GetTides(apiToken)
	weather := surf.GetWeather(apiToken)

	pushGaugeMetric()

	return nil
}
