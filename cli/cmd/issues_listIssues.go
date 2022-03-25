/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// jira "github.com/andygrunwald/go-jira"
)

// listIssuesCmd represents the list issues command for issue tracking
var listIssuesCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List of issues",
	Long:    `List of issues from issue tracking system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	issuesCmd.AddCommand(listIssuesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
