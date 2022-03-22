/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

*/
package cmd

import (
	"fmt"
	"strings"

	jira "github.com/andygrunwald/go-jira"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	issueName string
)

// getIssueCmd retrieves an issue by the issue ID. Currently this only supports
// Atlassian Jira. 
var getIssueCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an issue with the issue number passed",
	Long: `Get details about an issue based on the issue 
number passed in. This should be used as a way to get quick 
details about an issue.`,
	Run: func(cmd *cobra.Command, args []string) {
		jiraURL := viper.GetString("jira.url")
		username := viper.GetString("jira.username")
		password := viper.GetString("jira.password")

		if issueName == "" {
			fmt.Println("error: can't use empty string for issue")
			return
		}

		fmt.Printf("looking up issue: %s\n", issueName)

		tp := jira.BasicAuthTransport{
			Username: strings.TrimSpace(username),
			Password: strings.TrimSpace(password),
		}

		jiraClient, err := jira.NewClient(tp.Client(), strings.TrimSpace(jiraURL))
		if err != nil {
			fmt.Printf("\nerror: %v\n", err)
			return
		}

		issue, resp, err := jiraClient.Issue.Get(issueName, nil)
		if err != nil {
			fmt.Printf("\nerror: can't retrieve issue details: %v\n", err)
			return
		}

		if resp.StatusCode != 200 {
			fmt.Printf("\nerror: response error: %+v\n", resp.Status)
			return
		}

		fmt.Printf("\nResponse: %+v\n", resp.Status)
		fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
		fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
		fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
		fmt.Printf("Assignee: %v\n", issue.Fields.Assignee.DisplayName)
		fmt.Printf("Status: %v\n", issue.Fields.Status.Name)

	},
}

func init() {
	issuesCmd.AddCommand(getIssueCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getIssueCmd.PersistentFlags().StringVarP(&issueName, "issue", "i", "", "The issue to get details for")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
