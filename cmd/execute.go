package cmd

import (
	"os"
	"os/exec"
	"strings"
)

func execute(c string) {
	tokens := strings.Fields(c)
	executable := tokens[0]
	tokens = tokens[1:len(tokens)]

	cmd := exec.Command(executable, tokens...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		os.Exit(2)
	}
}
