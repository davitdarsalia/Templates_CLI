package configs

import (
	"fmt"
	"github.com/theckman/yacspin"
	"time"
)

var DefaultSpinnerConfig = yacspin.Config{
	Frequency:       150 * time.Millisecond,
	CharSet:         yacspin.CharSets[11],
	SuffixAutoColon: true,
	Message:         "Preparing Files",
	StopCharacter:   "✅  ",
	StopColors:      []string{"fgGreen"},
	StopMessage:     fmt.Sprintf("%s\n", "Template Generated"),
}
