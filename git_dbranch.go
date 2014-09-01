package main

import (
	"strings"

	"github.com/codegangsta/cli"
)

var gitDbranchCmd = cli.Command{
	Name:      "git-dbranch",
	ShortName: "gd",
	Usage:     "Deletes local and remote branches",
	Description: `Delete local and remote branches. For example,

  $ odt git-dbranch branch1 branch2 ...`,
	Action: gitDbranchAction,
}

func gitDbranchAction(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		cli.ShowCommandHelp(c, "gd")
		return
	}

	for _, arg := range args {
		branch := strings.TrimSpace(arg)
		deleteBranch(branch)
	}
}

func deleteBranch(branch string) {
	execCmd("git branch -D " + branch)
	execCmd("git push origin :" + branch)
}
