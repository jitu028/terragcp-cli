package cmd

import (
	"github.com/jitu028/terragcp-cli/pkg/flags"
	"github.com/jitu028/terragcp-cli/pkg/run"
	"github.com/spf13/cobra"
)

// chatCmd represents the chat command
var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start chat",
	Long: `
Start an interactive chat with Google Gemini model using this command.

Hit enter twice to send your prompt.

Alternatively, if you have blank lines in your text prompt then
enclose your prompt within {{ and }} as shown below
{{
This is a prompt with blank lines


The prompt ends here
}}
`,
	RunE: run.Chat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
	f := chatCmd.Flags()
	f.String(flags.Model, flags.ModelGeminiPro, "Model name")
	f.Float32(flags.TopP, -1, "Model TopP value (-1 means do not configure)")
	f.Int32(flags.TopK, -1, "Model TopK value (-1 means do not configure)")
	f.Float32(flags.Temperature, -1, "Model temperature (-1 means do not configure)")
	f.Int32(flags.CandidateCount, -1, "Model candidate count (-1 means do not configure)")
	f.Int32(flags.MaxOutputTokens, -1, "Model max output tokens (-1 means do not configure)")
	_ = chatCmd.RegisterFlagCompletionFunc(
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
					flags.ModelEmbedding001,
				},
				cobra.ShellCompDirectiveDefault
		},
	)
}
