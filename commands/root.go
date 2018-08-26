package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tienne/awssh/config"
)

var (
	versionGlobalFlag bool
)

func init() {
	RootCmd.Flags().BoolVarP(&versionGlobalFlag, "version", "v", false, "Print awssh version")
}

var RootCmd = &cobra.Command{
	Use:   config.AppName,
	Short: "aws ec2 ssh",
	Long:  fmt.Sprintf("%s is a powerful CLI to aws ec2 instance ssh connection", config.AppName),
	RunE: func(c *cobra.Command, args []string) error {
		if versionGlobalFlag {
			printVersion(c, args)
			return nil
		}
		return c.Usage()
	},
}
