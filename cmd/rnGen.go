package cmd

import (
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/davitdarsalia/CLI_Commands/pkg/reactNativeGen"
	"github.com/spf13/cobra"
	"os"
)

var (
	templateVariant string
	componentName   string
)

// rnGenCmd - Basic Command For Rn Template.
// Requires --template Flag: Ts Or Js Value
var rnGenCmd = &cobra.Command{
	Use:   constants.RnGen,
	Short: "React Native Template Generator",
	Long:  "Generates React Native Folder && File Template For Component",
	Run: func(cmd *cobra.Command, args []string) {
		switch templateVariant {
		case constants.TypeScript:
			reactNativeGen.GenerateRNTemplate("CalendarPicker", constants.TypeScript)
		case constants.JavaScript:
			reactNativeGen.GenerateRNTemplate("CalendarPicker", constants.JavaScript)
		}
	},
}

func init() {
	rnGenCmd.Flags().StringVarP(&templateVariant, constants.Template, constants.TemplateShortHand, "", constants.TemplateDescription)
	rnGenCmd.Flags().StringVarP(&componentName, constants.C, constants.CShorthand, "", constants.CDescription)

	if err := rnGenCmd.MarkFlagRequired(constants.TemplateFlag); err != nil {
		handleStdErr(err)
		os.Exit(1)
	}

	rootCmd.AddCommand(rnGenCmd)
}
