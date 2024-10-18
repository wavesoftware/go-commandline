package commandline

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"github.com/wavesoftware/go-retcode"
)

// ErrNoRootCommand is returned when no root command is provided.
var ErrNoRootCommand = errors.New("no root command provided")

// App represents a command line application.
type App struct {
	CobraProvider
	ErrorHandler
	Exit func(code int)
	root *cobra.Command
}

// ErrorHandler is a function that will be used to handle the errors. The
// function will be called regardless if an error has been received, so proper
// error checking is required. If true is returned, the default error handling
// will not be used.
type ErrorHandler func(err error) bool

// CobraProvider is used to provide a Cobra command.
type CobraProvider interface {
	Command() *cobra.Command
}

// New creates a new App from CobraProvider.
func New(cp CobraProvider) *App {
	return &App{
		CobraProvider: cp,
		Exit:          os.Exit,
	}
}

// ExecuteOrDie will execute the application or perform os.Exit in case of error.
func (a *App) ExecuteOrDie(options ...Option) {
	err := a.Execute(options...)
	if a.ErrorHandler == nil {
		a.defaultErrorHandler(err)
		return
	}
	if !a.ErrorHandler(err) {
		a.defaultErrorHandler(err)
	}
}

// Execute will execute the application with the provided options and return
// error if any.
func (a *App) Execute(options ...Option) error {
	if err := a.init(); err != nil {
		return err
	}
	for _, config := range options {
		config(a)
	}
	// cobra.Command should pass our own errors, no need to wrap them.
	return a.root.Execute() //nolint:wrapcheck
}

func (a *App) init() error {
	if a.Exit == nil {
		a.Exit = os.Exit
	}
	if a.CobraProvider == nil {
		return ErrNoRootCommand
	}
	a.root = a.Command()
	if a.root == nil {
		return ErrNoRootCommand
	}
	return nil
}

func (a *App) defaultErrorHandler(err error) {
	if err != nil {
		a.Exit(retcode.Calc(err))
	}
}
