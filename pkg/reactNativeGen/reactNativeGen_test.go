package reactNativeGen

import (
	"github.com/davitdarsalia/CLI_Commands/constants"
	"testing"
)

func BenchmarkGenerateRNTemplate(b *testing.B) {
	GenerateRNTemplate("CalendarPicker", constants.TypeScript)
}
