package main

import (
	"github.com/wavesoftware/go-commandline"
	"github.com/wavesoftware/go-commandline/test/internal/cli"
)

func main() {
	commandline.New(new(cli.App)).ExecuteOrDie(cli.Opts...)
}
