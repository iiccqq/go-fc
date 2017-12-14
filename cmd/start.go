package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*StartCmd ...增加命令
 */
func StartCmd() *cobra.Command {
	return startCmd
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the node.",
	Long:  `Starts a node that interacts with the network.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return serve(args)
	},
}

func serve(args []string) error {
	fmt.Println("start server:", args)
	return nil
}
