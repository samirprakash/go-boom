package maven

import (
	"strconv"

	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

// Package generates an artifact based on your maven based code base
func Package(cmd *cobra.Command, args []string) {
	c := "mvn package"
	skipTests, _ := strconv.ParseBool(args[0])
	if skipTests {
		c += " -DskipTests"
	}
	utils.Execute(c)
}