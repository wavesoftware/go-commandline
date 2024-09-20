package main_test

import (
	"bytes"
	"testing"

	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-commandline"
	main "github.com/wavesoftware/go-commandline/test/cmd"
	"github.com/wavesoftware/go-commandline/test/internal/cli"
	"gotest.tools/v3/assert"
)

func TestTheMain(t *testing.T) {
	s := capture(func() {
		main.Main()
	})
	assert.Equal(t, 0, s.exitCode)
	assert.Equal(t, "Hello, world!\n", s.out.String())
}

type state struct {
	exitCode int
	out      bytes.Buffer
}

func (s *state) opts() []commandline.Option {
	return []commandline.Option{
		commandline.WithCommand(func(cmd *cobra.Command) {
			cmd.SetOut(&s.out)
		}),
		commandline.WithExit(func(code int) {
			s.exitCode = code
		}),
	}
}

func capture(fn func()) state {
	var s state
	withOpts(fn, s.opts())
	return s
}

func withOpts(fn func(), opts []commandline.Option) {
	keep := cli.Opts
	defer func() {
		cli.Opts = keep
	}()
	cli.Opts = opts
	fn()
}
