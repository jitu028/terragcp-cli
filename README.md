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
generate terraform code to create a compute instance in gcp

      ┃ resource "google_compute_instance" "webserver" {
      ┃   name         = "webserver-1"
      ┃   machine_type = "n1-standard-1"
      ┃   zone         = "us-central1-a"
      ┃   disk {
      ┃     source_image = "debian-cloud/debian-11"
      ┃     type         = "PERSISTENT"
      ┃     auto_delete  = true
      ┃   }
      ┃   network_interface {
      ┃     name = "global/networks/default"
      ┃   }
      ┃ }

      This code creates a new Compute Engine instance named webserver-1 in the
      us-central1-a zone, with a machine type of n1-standard-1. The instance
      includes a single disk with the Debian 11 image, and is attached to the
      default network.

[2]>>> 
history saved to history-45a256b7-794b-47a5-8d02-593d3756ecf0.txt
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

## Parameter Tuning
Model config params such as `--top-p`, `--top-k`, `--temperature`, `--candiate-count` and 
`--max-output-tokens` can be supplied for fine tuning
