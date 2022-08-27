package cmd

import (
	"errors"
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/davitdarsalia/CLI_Commands/pkg/reactNativeGen"
	"github.com/spf13/cobra"
	"os"
)

var (
	templateVariant string
	componentName   string
	subExtension    string
	help            string
)

// rnGenCmd - Basic Command For Rn Template.
// Requires --template Flag: Ts Or Js Value
var rnGenCmd = &cobra.Command{
	Use:   constants.RnGen,
	Short: "React Native Template Generator",
	Long:  "Generates React Native Folder && File Template For Component",
	Run: func(cmd *cobra.Command, args []string) {
		if len(componentName) < 3 {
			handleStdErr(errors.New("invalid Component Name. Choose Proper One"))
			os.Exit(1)
		}

		switch templateVariant {
		case constants.TypeScript:
			reactNativeGen.GenerateRNTemplate(componentName, constants.TypeScript, subExtension)
		case constants.JavaScript:
			reactNativeGen.GenerateRNTemplate(componentName, constants.JavaScript, subExtension)
		default:
			// TODO - Implement Colored Output
			handleStdErr(errors.New("invalid Template"))
			os.Exit(1)
		}

	},
}

func init() {
	if help == "true" {
		rnGenCmd.Flags().StringVarP(&help, constants.HelpFlag, constants.HelpFlagShorthand, "", constants.HelpDescription)
	} else {
		rnGenCmd.Flags().StringVarP(&templateVariant, constants.TemplateFlag, constants.TemplateShortHand, "", constants.TemplateDescription)
		rnGenCmd.Flags().StringVarP(&componentName, constants.CFlag, constants.CShorthand, "", constants.CDescription)
		rnGenCmd.Flags().StringVarP(&subExtension, constants.SubExtensionFlag, constants.SubExtensionShorthand, "", constants.SubExtensionDescription)
	}

	if err := rnGenCmd.MarkFlagRequired(constants.TemplateFlag); err != nil {
		handleStdErr(err)
		os.Exit(1)
	}

	if err := rnGenCmd.MarkFlagRequired(constants.CFlag); err != nil {
		handleStdErr(err)
		return
	}

	rootCmd.AddCommand(rnGenCmd)
}
