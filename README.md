# Phoenix Framework CLI

[Phoenix Framework](https://www.phoenixframework.org/) enables developer to build rich, interactive web applications quickly, with less code and fewer moving parts.

The framework built with a bunch of task to generate boilerplate code aka `mix` tasks.

## Why?

What wrong with `mix phx.**`? I keep forgetting which argument comes first.
So, I decided to write this `phx` app to wrap `mix phx*` command with appropriate prompt messages.

## Installation

```bash
go install github.com/csokun/phx@v0.1.0
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

Generate Liveview:

```bash
$ phx gen live
? Context name (Plural noun):  Accounts
? Schema module name (Singular noun):  User
? Columns definition (e.g field_name:field_type):  name:string age:integer
? Schema's primary key use binary? (y/N) y
```
