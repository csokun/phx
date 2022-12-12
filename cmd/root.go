package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:           "phx",
		Short:         "Phoenixframework CLI",
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

// ErrUsage indicates we should print the usage string and exit with code 1
type ErrUsage struct {
	cmd *cobra.Command
}

func (e ErrUsage) Error() string {
	return e.cmd.UsageString()
}

// Indicates we want to exit with a specific error code without printing an error.
type ErrExitCode int

func (e ErrExitCode) Error() string {
	return fmt.Sprintf("exit code: %d", e)
}

func init() {
	// err := flags.Init(RootCmd)
	// if err != nil {
	// 	stdLog.Fatal(err.Error())
	// }
}

func Execute() error {
	return RootCmd.Execute()
}
