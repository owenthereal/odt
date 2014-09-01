package main

import (
	shellquote "github.com/kballard/go-shellquote"
	"log"
	"os"
	"os/exec"
	"strings"
)

func execCmd(input string) {
	inputs := splitShellInput(input)
	name := inputs[0]
	args := inputs[1:]

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func execCmdOutput(input string) []string {
	inputs := splitShellInput(input)
	name := inputs[0]
	args := inputs[1:]

	out, err := exec.Command(name, args...).Output()
	if err != nil {
		log.Fatal(err)
	}

	result := strings.Split(string(out), "\n")
	output := []string{}
	for _, o := range result {
		o = strings.TrimSpace(o)
		if o != "" {
			output = append(output, o)
		}
	}

	return output
}

func splitShellInput(input string) []string {
	inputs, err := shellquote.Split(input)
	if err != nil {
		inputs = strings.Split(input, " ")
	}

	return inputs
}