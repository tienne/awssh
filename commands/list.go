package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(lsCmd)
}

var lsCmd = &cobra.Command{
	Use:   "list",
	Short: "list EC2 instance",
	Long:  "Show EC2 instance List and filtering tag or instance id",
	RunE:  command,
}

func command(cmd *cobra.Command, args []string) error {
	if len(args) < 1 {
		return errors.New("no arguments")
	}
	return nil
}
