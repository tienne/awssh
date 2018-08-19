package commands

import (
	"fmt"
	"github.com/awssh/config"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show awssh version",

	Run: printVersion,
}

func printVersion(*cobra.Command, []string) {
	fmt.Fprint(os.Stderr, config.AwsshAsciiLogo)
	fmt.Println(config.CurrentBuildInfo)
}
