# terragcp-cli
CLI for analyzing and optimizing Google Terraform configurations using the Google Gemini AI model.

## Disclaimer
> The use of this tool does not guarantee security or usability for any particular purpose. Please review the code and use at your own risk.
>
> Don't trust, verify.

## Installation
This step assumes you have the [Go compiler toolchain](https://go.dev/dl/) installed on your system.

```bash
go install github.com/jitu028/terragcp-cli@latest

## installation
This step assumes you have [Go compiler toolchain](https://go.dev/dl/)
installed on your system.

```bash
go install github.com/jitu028/terragcp-cli@latest
```

Get an API key from [Google AI studio](https://makersuite.google.com/app/apikey)
and setup an env variable `GOOGLE_API_KEY` for it.

Setup shell completion. See more info at:
```bash
terragcp-cli completion -h
```

For instance, setup `bash` completion by adding following line to your `.bashrc`
```text
source <(terragcp-cli completion bash)
```

## usage
```bash
terragcp-cli chat [--auto-save]
```
`--auto-save` flag will save chat history to a randomly generated filename.

## example chat history

```bash
terragcp-cli chat --auto-save
```
```text
please type prompt below
press enter twice to send prompt
just enter to quit
[1]>>> {{
hi, can you explain to me this code below

package main

func main() {
  fmt.Println("hello world")
}

}}
      The code below is a simple "Hello, world!" program written in Go. It
      prints the string "hello world" to the standard output.

      ┃ package main
      ┃ 
      ┃ import "fmt"
      ┃ 
      ┃ func main() {
      ┃     fmt.Println("hello world")
      ┃ }

      Here's a breakdown of the code:
      • package main: This line declares the package name for the program. The
        main package is the entry point for the program.
      • import "fmt": This line imports the fmt package, which provides
        functions for formatted I/O.
      • func main() { ... }: This is the main function of the program. It's the
        entry point for the program, and it's where the program execution
        begins.
      • fmt.Println("hello world"): This line prints the string "hello world" to
        the standard output (usually the console or terminal window where the
        program is running).

      When you run the program, you should see the output:

      ┃ hello world

[2]>>> 
history saved to history-0d9d6887-ce12-4e89-824d-91b87b1a636f.txt
```

The prompt detects a blank line as termination, therefore, in order to send a prompt
that has blank lines in it, start the prompt with double curly braces `{{` and end
with `}}` as shown below.







## Analyze Terraform Configurations
Analyze and optimize your Terraform configurations with AI-powered insights and suggestions.
Below is an example:
```bash
terragcp-cli analyze config \
  --file path/to/your_config.tf \
  --format hcl \
  Please analyze this configuration
```

```bash
terragcp-cli analyze config \
  --file path/to/your_config.tf \
  --format hcl \
  List all resources in this configuration
```

```bash
terragcp-cli analyze config \
  --file path/to/your_config.tf \
  --format hcl \
Suggest fixes for this configuration
```

## Advanced Configuration
Fine-tune the model with parameters such as --top-p, --top-k, --temperature, --candidate-count, and --max-output-tokens.

## List Models
List available models and select one for your analysis using the --model flag.

```bash
terragcp-cli list models
```

## safety
`--allow-harm-probability` flag is set to `negligible` to prevent output from
displaying content that could be harmful. Change it at your own risk, for example,
```bash
terragcp-cli chat --allow-harm-probability=medium --auto-save
```

Model config params such as `--top-p`, `--top-k`, `--temperature`, `--candiate-count` and 
`--max-output-tokens` can be supplied for fine tuning
