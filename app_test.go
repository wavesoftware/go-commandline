package commandline_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
	"gotest.tools/v3/assert"
)

func TestExecuteOrDie(t *testing.T) {
	var buf bytes.Buffer
	var retcode int
	commandline.New(new(testApp)).ExecuteOrDie(
		commandline.WithCommand(func(cmd *cobra.Command) {
			cmd.SetOut(&buf)
			cmd.SetIn(bytes.NewBufferString("Input"))
			cmd.SetArgs([]string{"arg1", "arg2"})
		}),
		commandline.WithExit(func(code int) {
			retcode = code
		}),
	)
	assert.Equal(t, `example Input: ["arg1" "arg2"]`, buf.String())
	assert.Equal(t, 133, retcode)
}

func TestExit(t *testing.T) {
	app := commandline.App{CobraProvider: nil}
	err := app.Execute()
	assert.ErrorIs(t, err, commandline.ErrNoRootCommand)

	app = commandline.App{CobraProvider: nilApp{}}
	err = app.Execute()
	assert.ErrorIs(t, err, commandline.ErrNoRootCommand)
}

var errExample = errors.New("example error")

type testApp struct{}

func (t testApp) Command() *cobra.Command {
	return &cobra.Command{
		Use:           "example",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			in, err := io.ReadAll(cmd.InOrStdin())
			if err != nil {
				return err
			}
			cmd.Printf("%s %s: %q", cmd.Use, in, args)
			return errExample
		},
	}
}

type nilApp struct{}

func (n nilApp) Command() *cobra.Command {
	return nil
}
