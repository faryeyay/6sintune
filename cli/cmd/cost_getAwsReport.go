/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	ceTypes "github.com/aws/aws-sdk-go-v2/service/costexplorer/types"

	"github.com/faryeyay/6sintune/cli/pkg/util"
)

var (
	getAwsReportRegion      string
	getAwsReportGranularity string
	getAwsReportMetrics     string
)

// getAwsReportCmd represents the getAwsReport command
var getAwsReportCmd = &cobra.Command{
	Use:   "get-aws-report",
	Short: "Retrieve the AWS Cost and Usage Report",
	Long:  `Retrieve the AWS Cost and Usage Report`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Retrieving AWS Report")

		// validation
		// granularity flag options
		gfo := []string{"hourly", "daily", "monthly"}

		// make sure the granularity is a valid value
		if !util.SliceContainsString(gfo, getAwsReportGranularity) {
			log.Fatalf("invalid granularity")
		}

		mf := strings.Split(getAwsReportMetrics, ",")

		mo := []string{"AmortizedCost", "BlendedCost", "NetAmortizedCost", "NetUnblendedCost", "NormalizedUsageAmount", "UnblendedCost", "UsageQuantity"}

		// make sure the metrics passed are a subset of the
		// available metrics options
		if !util.StringSubset(mf, mo) {
			log.Fatalf("invalid metric detected")
		}

		// establish the context
		ctx := context.TODO()

		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(getAwsReportRegion))
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		ceClient := costexplorer.NewFromConfig(cfg)

		startDate := time.Now().AddDate(0, -1, 0).Format("2006-01-01")
		endDate := time.Now().Format("2006-01-01")

		tp := &ceTypes.DateInterval{
			Start: aws.String(startDate),
			End:   aws.String(endDate),
		}

		i := &costexplorer.GetCostAndUsageInput{
			Granularity: ceTypes.GranularityDaily,
			Metrics:     mf,
			TimePeriod:  tp,
		}

		// retrieve the cost and usage details
		out, err := ceClient.GetCostAndUsage(ctx, i)

		if err != nil {
			log.Fatalf("unable to retrieve cost and usage: %v", err)
		}

		log.Printf("%v", out)
	},
}

func init() {
	costCmd.AddCommand(getAwsReportCmd)

	// Here you will define your flags and configuration settings.
	getAwsReportCmd.Flags().StringVarP(&getAwsReportRegion, "region", "r", "us-east-1", "The AWS region for the cost report")
	getAwsReportCmd.Flags().StringVarP(&getAwsReportGranularity, "granularity", "g", "daily", "The interval at which to review the data. Valid options are hourly, daily, and monthly.")
	getAwsReportCmd.Flags().StringVarP(&getAwsReportMetrics, "metrics", "m", "NormalizedUsageAmount", "The metrics to retrieve")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getAwsReportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getAwsReportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
