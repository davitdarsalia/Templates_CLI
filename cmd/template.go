package cmd

import (
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/davitdarsalia/CLI_Commands/reactNativeGen"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	templateVariant string
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "Sets File Extension Template",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Template: TypeScript Or JavaScript

		switch templateVariant {
		case constants.TypeScript:
			reactNativeGen.GenerateRNTemplate("CalendarPicker", constants.TypeScript)
		case constants.JavaScript:
			reactNativeGen.GenerateRNTemplate("CalendarPicker", constants.JavaScript)
		default:
			log.Println("Wrong Flag: Choose Ts Or Js As A Flag Value")
		}
	},
}

func init() {

	templateCmd.Flags().StringVarP(&templateVariant, constants.Template, constants.TemplateShortHand, "", constants.TemplateDescription)

	if err := templateCmd.MarkFlagRequired(constants.TemplateFlag); err != nil {
		handleStdErr(err)
		os.Exit(1)
	}

	rnGenCmd.AddCommand(templateCmd)

}
