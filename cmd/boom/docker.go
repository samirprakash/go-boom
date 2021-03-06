package main

import (
	"github.com/samirprakash/boom/pkg/docker"
	"github.com/spf13/cobra"
)

var (
	imageTag         string
	appType          string
	uploadPath       string
	composeFile      string
	repoName         string
	healthcheckPorts string
	networkBridge    string
	testCollection   string
	environmentSpec  string
	currentImage     string
	newImage         string

	// dockerCmd is the parent command to execute docker and docker-compose actions
	// execute boom docker -h to check the available options
	dockerCmd = &cobra.Command{
		Use:   "docker",
		Short: "Execute docker commands",
		Long: `
[ boom docker ] provides an option to execute basic docker commands.
It requires at least one sub command from the list of options to be specified.

Prerequisites:
	- Install this binary and add it to your path
	- Install docker

Example usage options:
	- boom docker [ build | compose | run | tag ] -h
	- boom docker build [ --image-tag | -i ] [ --app-type | -t ] -h
	- boom docker compose [ --compose-file | -f ] [ --healthcheck-ports | -p ] [ --repo-name | -n ] -h
	- boom docker test [ --network-bridge | -n ] [ --test-collection | -c ] [ --environment-spec | -e ] -h
	- boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h
	- boom docker clean -h
		`,
		Args: cobra.MinimumNArgs(1),
	}

	// imageCmd is the subcommand to generate docker images
	imageCmd = &cobra.Command{
		Use:     "build",
		Short:   "Build docker images and push to a remote repository",
		Example: "boom docker build [ --image-tag | -i ] [ --app-type | -t ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := docker.Flags{
				UploadPath: uploadPath,
				ImageTag:   imageTag,
				AppType:    appType,
			}
			docker.BuildAndPush(&flags)
		},
	}

	// compose is the subcommand to start a docker compose environment to integration testing
	composeCmd = &cobra.Command{
		Use:     "compose",
		Short:   "Create docker compose environment based on the docker-compose.yaml in the code base",
		Example: "boom docker compose [ --compose-file | -f ] [ --healthcheck-ports | -p ] [ --repo-name | -n ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := docker.Flags{
				ComposeFile:      composeFile,
				RepoName:         repoName,
				HealthCheckPorts: healthcheckPorts,
			}
			docker.SetupContainerEnv(&flags)
		},
	}

	// run is the subcommand to execute tests collection on an existing docker compose environment
	runCmd = &cobra.Command{
		Use:     "test",
		Short:   "run collection of tests using newman command line runner",
		Example: "boom docker test [ --network-bridge | -n ] [ --test-collection | -c ] [ --environment-spec | -e ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := docker.Flags{
				NetworkBridge:   networkBridge,
				TestCollection:  testCollection,
				EnvironmentSpec: environmentSpec,
			}
			docker.ExecuteNewmanTests(&flags)
		},
	}

	// tag is the subcommand to tag and push images created by `go-doom docker compose` command
	tagCmd = &cobra.Command{
		Use:     "tag",
		Short:   "tag and push images to docker registry",
		Example: "boom docker tag [ --current-image | -i ] [ --new-image | -n ] -h",
		Run: func(cmd *cobra.Command, args []string) {
			flags := docker.Flags{
				CurrentImage: currentImage,
				NewImage:     newImage,
			}
			docker.TagAndPush(&flags)
		},
	}

	// clean is the subcommand to clean up containers created by `go-doom docker compose` command
	clearCmd = &cobra.Command{
		Use:     "clean",
		Short:   "clean up docker containers spinned up by docker compose",
		Example: "boom docker clean -h",
		Run: func(cmd *cobra.Command, args []string) {
			docker.Clean()
		},
	}
)

func init() {
	imageCmd.Flags().StringVarP(&uploadPath, "upload-to", "u", "", "specify the url to your docker registry")
	imageCmd.Flags().StringVarP(&imageTag, "image-tag", "i", "", "specify the tag for your image")
	imageCmd.Flags().StringVarP(&appType, "app-type", "t", "", "specify the application type - services/client")
	composeCmd.Flags().StringVarP(&composeFile, "compose-file", "f", "", "specify the compose file to used for setting up the environment")
	composeCmd.Flags().StringVarP(&repoName, "repo-name", "n", "", "specify the path to the config-source repository i.e. username/reponame")
	composeCmd.Flags().StringVarP(&healthcheckPorts, "healthcheck-ports", "p", "", "specify the healthcheck ports exposed in the compose file - use comma seperated format")
	runCmd.Flags().StringVarP(&networkBridge, "network-bridge", "n", "", "specify the network briidge applicable for running these tests")
	runCmd.Flags().StringVarP(&testCollection, "test-collection", "c", "", "specify the test collection file name in your integration-tests folder")
	runCmd.Flags().StringVarP(&environmentSpec, "environment-file", "e", "", "specify the newman environment file name in your integration-tests folder")
	tagCmd.Flags().StringVarP(&currentImage, "current-image", "i", "", "specify the tag of existing docker image")
	tagCmd.Flags().StringVarP(&newImage, "new-image", "n", "", "specify the tag name to tag the existing image with")

	rootCmd.AddCommand(dockerCmd)
	dockerCmd.AddCommand(imageCmd, composeCmd, runCmd, tagCmd, clearCmd)
}
