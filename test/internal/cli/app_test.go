package cli_test

import (
	"bytes"
	"testing"

	"github.com/wavesoftware/go-commandline/test/internal/cli"
	"gotest.tools/v3/assert"
)

func TestAppCommand(t *testing.T) {
	app := new(cli.App)
	cmd := app.Command()
	var buf bytes.Buffer
	cmd.SetOut(&buf)
	err := cmd.Execute()
	assert.NilError(t, err)
	assert.Equal(t, "Hello, world!\n", buf.String())
}
