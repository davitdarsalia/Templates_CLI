package cmd

import (
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/spf13/cobra"
)

var rnGenCmd = &cobra.Command{
	Use:   constants.RnGen,
	Short: "React Native Template Generator",
	Long:  "Generates React Native Folder && File Template For Component",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(rnGenCmd)
}
