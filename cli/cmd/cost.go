/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// costCmd represents the cost command
var costCmd = &cobra.Command{
	Use:   "cost",
	Short: "Retrieve cost reports for analysis",
	Long: `Retrieve cost reports for analysis. Useful
for tracking account spend or other details. This purpose 
will likely evolve over time.`,
}

func init() {
	rootCmd.AddCommand(costCmd)
}
