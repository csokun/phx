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

var contextAndModuleNameQs = []*survey.Question{
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
}

var fieldsAndPrimaryKeyTypeQs = []*survey.Question{
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

func runStandardGenerator(generator string) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		contextQs := append(contextAndModuleNameQs, fieldsAndPrimaryKeyTypeQs...)

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
			generator,
			answers.Name,
			answers.Module,
			strings.ToLower(tableName),
		}
		options = append(options, answers.Fields...)
		if answers.BinaryId {
			options = append(options, "--binary-id")
		}

		runMixCmd(options)
	}
}

var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Generates a context with functions around an Ecto schema.",
	Run:   runStandardGenerator("phx.gen.context"),
}

var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "Generates LiveView, templates, and context for a resource.",
	Run:   runStandardGenerator("phx.gen.live"),
}

var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Generates controller, views, and context for a JSON resource.",
	Run:   runStandardGenerator("phx.gen.json"),
}

var htmlCmd = &cobra.Command{
	Use:   "html",
	Short: "Generates controller, views, and context for an HTML resource.",
	Run:   runStandardGenerator("phx.gen.html"),
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
var channelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Generates a Phoenix channel.",
	Run: func(cmd *cobra.Command, args []string) {
		presenceQs := []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "Module name of the channel (e.g Room): "},
				Validate: survey.Required,
			},
		}

		answers := struct{ Name string }{Name: ""}

		err := survey.Ask(presenceQs, &answers)

		if err != nil {
			log.Fatal(err)
			return
		}

		runMixCmd([]string{"phx.gen.channel", answers.Name})
	},
}

var socketCmd = &cobra.Command{
	Use:   "socket",
	Short: "Generates a Phoenix socket handler.",
	Run: func(cmd *cobra.Command, args []string) {
		presenceQs := []*survey.Question{
			{
				Name:     "name",
				Prompt:   &survey.Input{Message: "Module name of the socket (e.g Room): "},
				Validate: survey.Required,
			},
		}

		answers := struct{ Name string }{Name: ""}

		err := survey.Ask(presenceQs, &answers)

		if err != nil {
			log.Fatal(err)
			return
		}

		runMixCmd([]string{"phx.gen.socket", answers.Name})
	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Generates authentication logic for a resource.",
	Run: func(cmd *cobra.Command, args []string) {
		answers := struct {
			Name   string
			Module string
		}{
			Name:   "",
			Module: "",
		}

		err := survey.Ask(contextAndModuleNameQs, &answers)

		if err != nil {
			log.Fatal(err)
			return
		}

		pl := pluralize.NewClient()
		tableName := pl.Plural(answers.Module)

		options := []string{
			"phx.gen.auth",
			answers.Name,
			answers.Module,
			strings.ToLower(tableName),
		}

		runMixCmd(options)
	},
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.AddCommand(
		contextCmd,
		liveCmd,
		jsonCmd,
		presenceCmd,
		channelCmd,
		socketCmd,
		authCmd,
	)
}
