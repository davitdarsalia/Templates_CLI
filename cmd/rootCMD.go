package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	// TODO - Change Command Name To Something Useful
	Use:   "templateGen",
	Short: "File Template Generator",
	Long:  "Generates A Code Template For Various Frameworks",
}

func ExecuteRootCMD() {
	if err := rootCmd.Execute(); err != nil {
		handleStdErr(err)
		return
	}
}

// handleStdErr - Helper Function Which Logs Out Errors On "Warn" Level
func handleStdErr(stdError error) {
	fmt.Fprintln(os.Stderr, stdError)
	return
}
