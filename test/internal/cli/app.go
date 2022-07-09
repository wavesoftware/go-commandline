package cli

import (
	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
)

// Opts is the list of commandline options to pass to the main function.
var Opts []commandline.Option

type App struct{}

func (a App) Command() *cobra.Command {
	return &cobra.Command{
		Use: "example",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println("Hello, world!")
		},
	}
}
