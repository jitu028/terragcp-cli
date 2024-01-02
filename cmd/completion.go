package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "generate shell completion",
	Long: `To load completions:

Bash:

  $ source <(terragcp-cli completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ terragcp-cli completion bash > /etc/bash_completion.d/terragcp-cli
  # macOS:
  $ terragcp-cli completion bash > /usr/local/etc/bash_completion.d/terragcp-cli

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ terragcp-cli completion zsh > "${fpath[1]}/_terragcp-cli"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ terragcp-cli completion fish | source

  # To load completions for each session, execute once:
  $ terragcp-cli completion fish > ~/.config/fish/completions/terragcp-cli.fish

PowerShell:

  PS> terragcp-cli completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> terragcp-cli completion powershell > terragcp-cli.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			_ = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			_ = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			_ = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			_ = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
