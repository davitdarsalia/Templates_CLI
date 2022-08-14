package reactNativeGen

import "testing"

func BenchmarkGenerateRNTemplate(b *testing.B) {
	GenerateRNTemplate("CalendarPicker")
}
