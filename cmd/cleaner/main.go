package main

import (
	"os"

	"github.com/denisvmedia/gitlab-container-registry-cleaner/pkg/cmd"
	"github.com/jessevdk/go-flags"
)

var Version string

func main() {
	if Version == "" {
		Version = "develop"
	}

	app := &cmd.AppCommand{}
	parser := flags.NewParser(app, flags.Default)
	cmd.RegisterCleanCommand(parser)

	_, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		os.Exit(-1)
	}
}
