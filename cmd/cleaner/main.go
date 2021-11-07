package main

import (
	"os"

	"github.com/denisvmedia/gitlab-container-registry-cleaner/pkg/cmd"
	"github.com/jessevdk/go-flags"
)

var (
	AppVersion string
	GitCommit  string
)

func main() {
	if AppVersion == "" {
		AppVersion = "develop"
	}
	if GitCommit == "" {
		GitCommit = "unknown"
	}

	app := &cmd.AppCommand{}
	parser := flags.NewParser(app, flags.Default)
	cmd.RegisterCleanCommand(parser)
	cmd.RegisterVersionCommand(parser, AppVersion, GitCommit)

	_, err := parser.ParseArgs(os.Args[1:])
	if err != nil {
		os.Exit(-1)
	}
}
