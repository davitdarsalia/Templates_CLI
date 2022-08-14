/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/reactNativeGen"

	"github.com/spf13/cobra"
)

// templateCmd represents the template command
var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Sets File Extension Template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("template called")

		reactNativeGen.GenerateRNTemplate("CalendarPicker")
	},
}

func init() {
	rnGenCmd.AddCommand(templateCmd)

}
