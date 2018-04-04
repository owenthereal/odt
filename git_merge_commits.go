package main

import (
	"fmt"
	"regexp"

	"github.com/codegangsta/cli"
)

var gitMergeCommitsCmd = cli.Command{
	Name:        "git-merge-commits",
	ShortName:   "gmc",
	Usage:       "Show a list of merge commits since SHA",
	Description: `Show a list of merge commits since a given commit SHA`,
	Action:      gitMergeCommitsAction,
}

func gitMergeCommitsAction(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		cli.ShowCommandHelp(c, "gmc")
		return
	}

	sha := args[0]
	out := execCmdOutput("git log --merges " + sha + "..HEAD")

	reg := regexp.MustCompile(`Merge pull request #(\d+) from .+`)

	var (
		result []string
		prNum  string
	)
	for _, o := range out {
		if prNum != "" {
			result = append(result, o+" "+prNum)
			prNum = ""
			continue
		}

		if reg.MatchString(o) {
			prNum = "(#" + reg.FindAllStringSubmatch(o, -1)[0][1] + ")"
		}
	}

	for _, r := range result {
		fmt.Println(r)
	}
}
