# `polygen`

CLI for generating Polybase Collections API code.

## Setup

To get the [PolyGen](https://github.com/durudex/polygen) CLI, you need to have or install Go version >= [1.18](https://go.dev/dl/). To check your current version of Go, use the `go version` command.

**The command to install cli:**

```bash
go install github.com/durudex/polygen/cmd/polygen@latest
```

## Config

To start using PolyGen, you need to create a configuration file. It defines the rules for generating the code you need.

[**An example of a configuration file:**](https://github.com/durudex/polygen/blob/main/.polygen.example.yml)

```yml
collection:
  - "Collection"
  - "Namespace/Collection"

language:
  # Add configurations for the programming languages you need.
  # An example configuration is available in the GitHub repository.
  ...
```

## Usage

To start generating code using the configuration you created, you need to use the CLI with the specified path to the configuration file.

**The command to start code generation:**

```bash
polygen --config .polygen.yml
```

## License

Copyright © 2022-2023 [Durudex](https://github.com/durudex). Released under the MIT license.
