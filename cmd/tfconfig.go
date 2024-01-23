package cmd

import (
	"github.com/jitu028/terragcp-cli/pkg/flags"
	"github.com/jitu028/terragcp-cli/pkg/run"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command for analyzing Terraform configurations
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Analyze Terraform configurations",
	RunE:  run.AnalyzeConfig,
}

func init() {
	analyzeCmd.AddCommand(configCmd)

	// Define flags for the config command
	f := configCmd.Flags()
	f.String(flags.Model, flags.ModelGeminiPro, "Model name")
	f.StringP("file", "f", "", "Path to the Terraform configuration file")
	f.String("format", "hcl", "Configuration file format (default is HCL - HashiCorp Language)")

	// Bind the flags to viper keys
	_ = viper.BindPFlag(flags.Model, f.Lookup(flags.Model))
	_ = viper.BindPFlag("config_file", f.Lookup("file"))
	_ = viper.BindPFlag("config_format", f.Lookup("format"))

	// Register flag completion functions
	_ = configCmd.RegisterFlagCompletionFunc(
		flags.Model,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					flags.ModelGeminiPro,
					flags.ModelGeminiProVision,				
					flags.ModelEmbedding001,
				},
				cobra.ShellCompDirectiveDefault
		},
	)

	_ = configCmd.RegisterFlagCompletionFunc(
		"format",
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{"hcl"},
				cobra.ShellCompDirectiveDefault
		},
	)
}

// Usage example:
// terragcp-cli analyze config \
//   --file path/filename.tf \
//   --format hcl \
//   "could you please list resources created by config"
