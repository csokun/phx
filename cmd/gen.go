/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	pluralize "github.com/gertd/go-pluralize"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "gen [command]",
	Short: "Run mix phx.gen.*",
}

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Generates a context with functions around an Ecto schema.",
	Run: func(cmd *cobra.Command, args []string) {
		contextQs := []*survey.Question{
			{
				Name:      "name",
				Prompt:    &survey.Input{Message: "Context name (Plural noun): ", Help: "e.g Accounts"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:      "module",
				Prompt:    &survey.Input{Message: "Schema module name (Singular noun): ", Help: "e.g User"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name: "fields",
				Prompt: &survey.Input{
					Message: "Columns definition (e.g field_name:field_type): ",
					Help:    "Field type :integer :float :decimal :boolean :map :string :array :references :text :date :time :utc_datetime :uuid :binary :enum",
				},
				Validate: survey.Required,
				Transform: func(ans interface{}) (newAns interface{}) {
					s, ok := ans.(string)
					if !ok {
						return []string{}
					}
					return strings.Split(s, " ")
				},
			},
			{
				Name:     "binaryId",
				Prompt:   &survey.Confirm{Message: "Schema's primary key use binary?"},
				Validate: survey.Required,
			},
		}

		answers := struct {
			Name     string
			Module   string
			Fields   []string
			BinaryId bool
		}{
			Name:     "",
			Module:   "",
			Fields:   []string{},
			BinaryId: true,
		}

		err := survey.Ask(contextQs, &answers)

		if err != nil {
			log.Fatal(err)
			return
		}

		pl := pluralize.NewClient()
		tableName := pl.Plural(answers.Module)

		options := []string{
			"phx.gen.context",
			answers.Name,
			answers.Module,
			strings.ToLower(tableName),
		}
		options = append(options, answers.Fields...)
		if answers.BinaryId {
			options = append(options, "--binary-id")
		}

		runMixCmd(options)
	},
}

var presenceCmd = &cobra.Command{
	Use:   "presence",
	Short: "Generates a Presence tracker.",
	Run: func(cmd *cobra.Command, args []string) {
		presenceQs := []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "Module name of the Presence tracker: "},
				Validate: survey.Required,
			},
		}

		answers := struct{ Name string }{Name: ""}

		err := survey.Ask(presenceQs, &answers)

		if err != nil {
			log.Fatal(err)
			return
		}

		runMixCmd([]string{"phx.gen.presence", answers.Name})
	},
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.AddCommand(contextCmd, presenceCmd)
}
