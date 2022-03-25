/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	watchIssueName string
)

// watchIssueCmd
var watchIssueCmd = &cobra.Command{
	Use:   "watch",
	Short: "Add yourself as a watcher to an issue",
	Long: `Add yourself as a watcher to an issue. 
This will add the user defined in the configuration file
for 6sintune to the watcher's list.`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraURL := viper.GetString("jira.url")
		email := viper.GetString("jira.email")
		username := viper.GetString("jira.username")
		password := viper.GetString("jira.password")

		if watchIssueName == "" {
			fmt.Println("error: can't use empty string for issue")
			return
		}

		tp := jira.BasicAuthTransport{
			Username: strings.TrimSpace(email),
			Password: strings.TrimSpace(password),
		}

		jiraClient, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}

		// blank string == self
		resp, err := jiraClient.Issue.AddWatcher(watchIssueName, "")
		if err != nil {
			fmt.Printf("\nerror: can't add you as a watcher. %v\n", err)
			return
		}

		if resp.StatusCode != 204 {
			fmt.Printf("\nerror: response error: %+v\n", resp.Status)
			return
		}

		fmt.Printf("\nadded %s as a watcher to %s\n", username, watchIssueName)
	},
}

func init() {
	issuesCmd.AddCommand(watchIssueCmd)

	// the issue to watch
	watchIssueCmd.PersistentFlags().StringVarP(&watchIssueName, "issue", "i", "", "The issue to get details for")
}
