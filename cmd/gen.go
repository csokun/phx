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
				Prompt:    &survey.Input{Message: "Context name: "},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:      "module",
				Prompt:    &survey.Input{Message: "Module name: ", Help: "e.g User"},
				Validate:  survey.Required,
				Transform: survey.Title,
			},
			{
				Name:     "fields",
				Prompt:   &survey.Input{Message: "Columns: ", Help: "name:string age:integer"},
				Validate: survey.Required,
				Transform: func(ans interface{}) (newAns interface{}) {
					s, ok := ans.(string)
					if !ok {
						return []string{}
					}
					return strings.Split(s, " ")
				},
			},
		}

		answers := struct {
			Name   string
			Module string
			Fields []string
		}{
			Name:   "",
			Module: "",
			Fields: []string{},
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

		runMixCmd(options)
	},
}

func init() {
	RootCmd.AddCommand(genCmd)

	genCmd.AddCommand(contextCmd)
}
