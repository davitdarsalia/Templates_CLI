package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rnGenCmd represents the rnGen command
var rnGenCmd = &cobra.Command{
	Use:   "rnGen",
	Short: "React Native Template Generator",
	Long:  "Generates React Native Folder-File Template For Component",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		fmt.Println("RnGen Called")
	},
}

func init() {
	rootCmd.AddCommand(rnGenCmd)
}
