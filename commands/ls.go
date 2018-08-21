package commands

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "list EC2 instance",
	Long:  "Show EC2 instance List and filtering tag or instance id",
	Run:   command,
}

func command(cmd *cobra.Command, args []string) {

}
