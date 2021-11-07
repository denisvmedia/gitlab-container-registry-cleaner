package cmd

import (
	"fmt"
	"runtime"

	"github.com/jessevdk/go-flags"
)

type VersionCommand struct {
	appVersion string
	gitCommit  string
}

func RegisterVersionCommand(parser *flags.Parser, appVersion, gitCommit string) *VersionCommand {
	cmd := &VersionCommand{
		appVersion: appVersion,
		gitCommit:  gitCommit,
	}
	_, err := parser.AddCommand("version", "application version", "", cmd)
	if err != nil {
		panic(err)
	}

	return cmd
}

func (cmd *VersionCommand) Execute(args []string) error {
	fmt.Printf("gitlab container registry tool version %s@%s (%s/%s)\n", cmd.appVersion, cmd.gitCommit, runtime.GOOS, runtime.GOARCH)

	return nil
}
