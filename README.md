# Phoenix Framework CLI

[Phoenix Framework](https://www.phoenixframework.org/) enables developer to build rich, interactive web applications quickly, with less code and fewer moving parts.

The framework built with a bunch of task to generate boilerplate code aka `mix` tasks.

## Why?

What wrong with `mix phx.**`? I keep forgetting which argument come first.
So, I build this utility app which wrapping `mix phx*` command with appropriate prompt messages.

## Installation

```bash
go install github.com/csokun/phx.git
```

## Usage

```bash
Phoenixframework CLI

Usage:
  phx [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         Run mix phx.gen.*
  help        Help about any command
  new         Command to create new Phoenix project

Flags:
  -h, --help   help for phx

Use "phx [command] --help" for more information about a command.
```
