package reactNativeGen

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/constants"
	"log"
	"os"
)

type fileTemplate string

func generateOverallFile(basePath string) (string, error) {
	err := os.Mkdir(basePath, os.ModePerm)

	if err != nil {
		log.Println(constants.CreateDirError)
		return "", err
	}

	return basePath, nil
}

func GenerateFiles(dir, fileName string) error {
	componentFileTemplate := fmt.Sprintf("%s/%s.tsx", dir, fileName)
	stylesFileTemplate := fmt.Sprintf("%s/style.tsx", dir)

	// Component File
	componentFile, err := os.Create(componentFileTemplate)
	// Styles File
	stylesFile, err := os.Create(stylesFileTemplate)

	if err != nil {
		log.Println(constants.CreateFileError)
		return err
	}

	// Components File Writer
	componentFile.Write([]byte(constants.ComponentTemplate))
	// Styles File Writer
	stylesFile.Write([]byte(constants.StyleTemplate))
	return nil
}

func GenerateRNTemplate(dirName string) {
	dirName, err := generateOverallFile(dirName)
	if err != nil {
		log.Println(err)
		return
	}

	err = GenerateFiles(dirName, dirName)
	if err != nil {
		log.Println(err)
		return
	}

	// Writesr

}
