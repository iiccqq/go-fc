package main

import (
	"fmt"
	"os"

	"github.com/iiccqq/go-fc/cmd"
	"github.com/spf13/cobra"
)

var versionFlag bool
var mainCmd = &cobra.Command{
	Use: "go-fc",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if versionFlag {
			//        version.Print()
		} else {
			cmd.HelpFunc()(cmd, args)
		}
		fmt.Println(args)
	},
}

func main() {
	mainCmd.AddCommand(cmd.StartCmd())
	if mainCmd.Execute() != nil {
		os.Exit(1)
	}
}
