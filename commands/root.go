package commands

import (
	"github.com/spf13/cobra"
)

var (
	versionGlobalFlag bool

	//renderGreenFn    = color.New(color.FgGreen).SprintFunc()
	//renderRedFn      = color.New(color.FgRed).SprintFunc()
	//renderYellowFn   = color.New(color.FgYellow).SprintFunc()
	//renderBlueFn     = color.New(color.FgBlue).SprintFunc()
	//renderCyanBoldFn = color.New(color.FgCyan, color.Bold).SprintFunc()
)

func init() {
	RootCmd.Flags().BoolVar(&versionGlobalFlag, "version", false, "Print awless version")
}

var RootCmd = &cobra.Command{
	Use:   "awssh",
	Short: "aws ec2 ssh",
	Long:  "awssh is a powerful CLI to aws ec2 instance ssh connection",
	RunE: func(c *cobra.Command, args []string) error {
		if versionGlobalFlag {
			printVersion(c, args)
			return nil
		}
		return c.Usage()
	},
}
