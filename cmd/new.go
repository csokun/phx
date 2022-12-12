package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name:   "name",
		Prompt: &survey.Input{Message: "Create new OTP application:"},
		Validate: func(val interface{}) error {
			// since we are validating an Input, the assertion will always succeed
			if str, ok := val.(string); !ok || len(str) > 10 {
				return errors.New("This response cannot be longer than 10 characters.")
			}
			return nil
		},
	},
	{
		Name:     "umbrella",
		Prompt:   &survey.Confirm{Message: "Do you want to generate an umbrella project?"},
		Validate: survey.Required,
	},
	{
		Name: "database",
		Prompt: &survey.Select{
			Message: "Choose database:",
			Options: []string{"none", "postgres", "mysql", "mssql", "sqlite3"},
			Default: "none",
		},
		Validate: survey.Required,
	},
	{
		Name:     "live",
		Prompt:   &survey.Confirm{Message: "Do you want to enable Liveview?"},
		Validate: survey.Required,
	},
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Command to create new Phoenix project",
	RunE: func(cmd *cobra.Command, args []string) error {
		answers := struct {
			Name     string
			Umbrella bool
			Database string
			Live     bool
		}{
			Name:     "",
			Umbrella: false,
			Database: "",
			Live:     false,
		}
		err := survey.Ask(qs, &answers)

		if err != nil {
			return err
		}

		options := []string{"phx.new", answers.Name}

		//options = append(options, "--app "+answers.Name)
		if answers.Umbrella {
			options = append(options, "--umbrella")
		}

		if answers.Database != "none" {
			options = append(options, "--database", answers.Database)
		}

		if !answers.Live {
			options = append(options, "--no-live")
		}

		fmt.Printf("%s\n", strings.Join(options, " "))
		shellCmd := exec.Command("mix", options...)

		shellCmd.Stdin = os.Stdin
		shellCmd.Stdout = os.Stdout
		shellCmd.Stderr = os.Stderr
		return shellCmd.Run()
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
