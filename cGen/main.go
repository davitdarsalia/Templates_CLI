package main

import (
	"github.com/davitdarsalia/CLI_Commands/cGen/cmd"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	cmd.ExecuteRootCMD()
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.WarnLevel)
}
