package main

import (
	"github.com/davitdarsalia/CLI_Commands/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	//reactNativeGen.GenerateRNTemplate("CalendarPicker")
	cmd.ExecuteRootCMD()
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.WarnLevel)
}
