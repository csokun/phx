package cmd

import (
	"errors"
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name:   "name",
		Prompt: &survey.Input{Message: "Create new OTP application:"},
		Validate: func(val interface{}) error {
			if str, ok := val.(string); !ok || len(str) > 20 {
				return errors.New("Application name is longer than 20 characters.")
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
	Run: func(cmd *cobra.Command, args []string) {
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
			log.Fatal(err)
			return
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

		runMixCmd(options)
	},
}

func init() {
	RootCmd.AddCommand(newCmd)
}
