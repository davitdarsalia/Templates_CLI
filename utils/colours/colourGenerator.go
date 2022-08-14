package colours

import (
	"fmt"
	"github.com/fatih/color"
)

func useColoredFont(value interface{}, colorCode int) string {
	if colorCode > 29 && colorCode < 38 {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", colorCode, value)
	} else {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", Red, value)
	}
}

func useColorBackground(value interface{}, colorCode int) string {
	if colorCode > 39 && colorCode < 48 {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", colorCode, value)
	} else {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", color.BgRed, value)
	}
}

func useFontStyle(value interface{}, fontStyle int) string {
	if fontStyle >= 0 && fontStyle < 9 {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", fontStyle, value)
	} else {
		return fmt.Sprintf("\x1b[%dm%v\x1b[0m", Bright, value)
	}
}
