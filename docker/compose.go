package docker

import (
	"fmt"
	"os"

	"github.com/samirprakash/boom/utils"
	"github.com/spf13/cobra"
)

// SetupContainerEnv stands up a docker container environement and verifies if the containers are ready for use by doing helthchecks
func SetupContainerEnv(cmd *cobra.Command, args []string) {
	composeFile := args[0]
	healthcheckPorts := args[1]
	if composeFile == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
		return
	} else if healthcheckPorts == "" {
		fmt.Fprintln(os.Stderr, "\nMissing data - please provide the healthcheck ports exposed in the docker compose file. \nRun `boom docker compose -h` for usage guidelines!")
		return
	}
	// clone config source repo if not already present in the build environment
	path := os.Getenv("TC_CONFIG_PATH")
	repo, _ := utils.Exists(path)
	if !repo {
		fmt.Println("cloning into : ", path)
		cloneConfig := "git clone git@github.com:toyota-connected/pg-config-source.git " + path
		utils.Execute(cloneConfig)
	}
	fmt.Println("repository that is being cloned already exists on the build environment")
	setupEnvironment := "docker-compose -f " + composeFile + " up --build --detach --remove-orphans"
	utils.Execute(setupEnvironment)
	// check if the docker containers are healthy or not based on the ports that have been exposed from docker-compose.yaml
	utils.Healthcheck(healthcheckPorts)
}