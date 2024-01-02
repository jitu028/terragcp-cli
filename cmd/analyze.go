package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// analyzeCmd represents the analyze command
var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze images command group",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("please use with sub command")
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)

}
