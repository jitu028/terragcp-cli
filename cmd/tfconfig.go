package cmd

import (
	"github.com/jitu028/terragcp-cli/pkg/flags"
	"github.com/jitu028/terragcp-cli/pkg/run"
	"github.com/spf13/cobra"
)

// configCmd represents the config command for analyzing Terraform configurations
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Analyze Terraform configurations",
	RunE:  run.AnalyzeConfig,
}

func init() {
	analyzeCmd.AddCommand(configCmd)
	f := configCmd.Flags()
	f.String(flags.File, "", "Path to the Terraform configuration file")
	f.String(flags.Format, "hcl", "Configuration file format (default is HCL - HashiCorp Language)")

	// Assuming the implementation of the AnalyzeConfig function in the run package
	// This function should contain the logic to analyze the Terraform configuration file
}

// Usage example:
// terragcp-cli analyze config \
//   --file path/filename.tf \
//   --format hcl \
//   "could you please list resources created by config"
