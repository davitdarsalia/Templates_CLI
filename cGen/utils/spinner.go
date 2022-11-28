package utils

import (
	"github.com/theckman/yacspin"
	"log"
	"os"
	"time"
)

func InitSpinner(message, stopMessage, stopIcon string, timing int16, animationOffset uint8) *yacspin.Spinner {
	var animationIndex int8

	if animationOffset < 0 || animationOffset > 90 {
		animationIndex = int8(59)
	}

	defSpinnerConfig := yacspin.Config{
		Frequency:       time.Duration(timing) * time.Millisecond,
		CharSet:         yacspin.CharSets[int(animationIndex)],
		SuffixAutoColon: true,
		Message:         message,
		StopCharacter:   stopIcon,
		StopMessage:     stopMessage,
		StopColors:      []string{"fgGreen"},
	}

	spinner, err := yacspin.New(defSpinnerConfig)

	if err != nil {
		log.Println("Unable To Initialize A Spinner")
		os.Exit(1)
	}

	return spinner
}
