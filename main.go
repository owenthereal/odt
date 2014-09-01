package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "odt"
	app.Usage = "Owen's Development Tool"
	app.EnableBashCompletion = true
	app.Version = Version
	app.Commands = []cli.Command{
		gitRmergeCmd,
		gitDbranchCmd,
	}
	app.Run(os.Args)
}
