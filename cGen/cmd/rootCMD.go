package cmd

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/cGen/utils"
	"github.com/spf13/cobra"
	"os"
)

var (
	name   string
	domain string
)

var rootCmd = &cobra.Command{
	Short: "React Native Component Template Generator",
	Long:  "Generates React Native Folder && File Template For Component",
	Run: func(cmd *cobra.Command, args []string) {
		utils.Component(name, domain)
	},
}

func ExecuteRootCMD() {
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Give Your Component A Name")
	rootCmd.Flags().StringVarP(&domain, "domain", "d", "", "Give Your Component A Domain")

	if err := rootCmd.MarkFlagRequired("name"); err != nil {
		handleStdErr(err)
		os.Exit(1)
	}

	if err := rootCmd.MarkFlagRequired("domain"); err != nil {
		handleStdErr(err)
		return
	}

	if err := rootCmd.Execute(); err != nil {
		handleStdErr(err)
		return
	}

}

func init() {

}

// handleStdErr - Helper Function Which Logs Out Errors On "Warn" Level
func handleStdErr(stdError error) {
	fmt.Fprintln(os.Stderr, stdError)
	return
}
