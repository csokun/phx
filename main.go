package main

import (
	"fmt"
	"os"

	"github.com/csokun/phx/cmd"
)

func main() {
	err := cmd.Execute()
	if e, ok := err.(cmd.ErrExitCode); ok {
		os.Exit(int(e))
	} else if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
