package cmd

import (
	"github.com/jitu028/terragcp-cli/pkg/flags"
	"github.com/jitu028/terragcp-cli/pkg/run"
	"github.com/spf13/cobra"
)

// imageCmd represents the images command
var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Analyze images",
	RunE:  run.AnalyzeImages,
}

func init() {
	analyzeCmd.AddCommand(imageCmd)
	f := imageCmd.Flags()
	f.String(flags.Model, flags.ModelGeminiProVision, "Model name")
	f.StringSlice(flags.File, nil, "Image filenames")
	f.StringSlice(flags.Format, nil, "Image formats (assumes image/jpeg when unspecified)")
	_ = imageCmd.RegisterFlagCompletionFunc(
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

	_ = imageCmd.RegisterFlagCompletionFunc(
		flags.Format,
		func(
			cmd *cobra.Command,
			args []string,
			toComplete string,
		) (
			[]string,
			cobra.ShellCompDirective,
		) {
			return []string{
					flags.FormatJpeg,
					flags.FormatPng,
					flags.FormatHeif,
					flags.FormatHeic,
					flags.FormatWebp,
				},
				cobra.ShellCompDirectiveDefault
		},
	)
}
