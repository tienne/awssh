package commands

import (
	"github.com/spf13/cobra"
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
	//fmt.Fprint(os.Stderr, config.AWLESS_ASCII_LOGO)
	//fmt.Println(config.CurrentBuildInfo)
}
