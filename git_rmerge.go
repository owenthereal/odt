package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"

	"github.com/codegangsta/cli"
)

var gitRmergeCmd = cli.Command{
	Name:      "git-rmerge",
	ShortName: "gm",
	Usage:     "Runs Git rebase and Git merge with --no-ff against current branch",
	Description: `Run Git rebase on a branch and then run Git merge with no fast forward
(git merge --no-ff).

As an example, assuming current branch is master, running this command
rebases a list of topic branches on top of master and then merge them
into master with no fast forward.

  $ odt git-rmerge topic1 topic2 ...`,
	Action: gitRmergeAction,
}

func gitRmergeAction(c *cli.Context) {
	args := c.Args()
	if len(args) == 0 {
		cli.ShowCommandHelp(c, "gm")
		return
	}

	baseBranch := getBaseBranch()
	for _, arg := range args {
		topicBranch := strings.TrimSpace(arg)

		execCmd("git fetch")

		execCmd("git checkout " + topicBranch)
		if hasRemoteBranch(topicBranch) {
			execCmd("git pull origin " + topicBranch)
		}
		execCmd("git rebase -i origin/" + baseBranch)
		execCmd("git push origin HEAD -f")

		execCmd("git checkout " + baseBranch)
		if hasRemoteBranch(baseBranch) {
			execCmd("git pull origin " + baseBranch)
		}
		execCmd("git merge " + topicBranch + " --no-ff")
		execCmd("git push origin HEAD")

		deleteBranch(topicBranch)
	}
}

func getBaseBranch() string {
	out, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(out))
}

func hasRemoteBranch(branch string) bool {
	out, err := exec.Command("git", "branch", "-r").Output()
	if err != nil {
		return false
	}
	for _, line := range bytes.Split(out, []byte{'\n'}) {
		if strings.TrimSpace(string(line)) == ("origin/" + branch) {
			return true
		}
	}

	return false
}
