package run

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/google/generative-ai-go/genai"
	"github.com/google/uuid"
	"github.com/jitu028/terragcp-cli/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func AnalyzeConfig(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Bind flags and environment variables
	_ = viper.BindPFlag(flags.ApiKey, cmd.Flag(flags.ApiKey))
	_ = viper.BindPFlag(flags.Model, cmd.Flag(flags.Model))
	_ = viper.BindPFlag(flags.File, cmd.Flag(flags.File))
	_ = viper.BindPFlag(flags.AutoSave, cmd.Flag(flags.AutoSave))
	_ = viper.BindEnv(flags.ApiKey, flags.ApiKeyEnv)

	apiKey := viper.GetString(flags.ApiKey)
	modelName := viper.GetString(flags.Model)
	filePath := viper.GetString(flags.File)
	autoSave := viper.GetBool(flags.AutoSave)

	if len(apiKey) == 0 || len(modelName) == 0 {
		return fmt.Errorf("api-key or model cannot be empty")
	}

	if len(filePath) == 0 {
		return fmt.Errorf("file path cannot be empty")
	}

	// Read the Terraform configuration file
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Create a new client for the Gemini API
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("failed to create new genai client: %w", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelName)

	// Prepare the prompt for the API
	var prompt string
	if len(args) > 0 {
		prompt = strings.Join(args, " ")
	} else {
		prompt = "Analyze the following Terraform configuration and provide insights:"
	}

	// Autosave history setup
	fileName := fmt.Sprintf("history-%s.txt", uuid.New().String())
	var fileWriter *bufio.Writer
	if autoSave {
		f, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create history file: %w", err)
		}
		defer f.Close()

		fileWriter = bufio.NewWriter(f)
		if _, err := fileWriter.WriteString(
			fmt.Sprintf("Command: %s\nTimestamp: %s\n",
				strings.Join(os.Args, " "),
				time.Now().String(),
			),
		); err != nil {
			return fmt.Errorf("failed to write to history file: %w", err)
		}
	}

	// Send the request to the Gemini API
	s := "...sending prompt... please wait"
	_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\r", s)
	res, err := model.GenerateContent(ctx, genai.Text(fileContent), genai.Text(prompt))
	if err != nil {
		return fmt.Errorf("failure at backend: %w", err)
	}
	_, _ = fmt.Fprintf(cmd.OutOrStdout(), "%s\r", strings.Repeat(" ", len(s)+2))

	// Print the response using the printResponse function
	if err := printResponse(res, cmd.OutOrStdout(), flags.RenderFormatPretty, autoSave, fileWriter); err != nil {
		return fmt.Errorf("failed to process response: %w", err)
	}

	// Autosave history
	if autoSave {
		if err := fileWriter.Flush(); err != nil {
			return fmt.Errorf("failed to write to history file: %w", err)
		}
		if _, err := fmt.Fprintln(cmd.OutOrStdout(), fmt.Sprintf("history saved to %s", fileName)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	return nil
}
