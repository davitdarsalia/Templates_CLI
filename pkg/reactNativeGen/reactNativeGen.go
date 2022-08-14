package reactNativeGen

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/davitdarsalia/CLI_Commands/pkg/spinner"
	"log"
	"os"
	"time"
)

func generateOverallFile(basePath string) (string, error) {
	err := os.Mkdir(basePath, os.ModePerm)

	if err != nil {
		log.Println(constants.CreateDirError)
		return "", err
	}

	return basePath, nil
}

func generateFiles(dir, fileName, template string) error {
	var extension string
	var componentTemplate string

	switch template {
	case constants.TypeScript:
		extension = constants.TypescriptExtension
		componentTemplate = fmt.Sprintf("import React from 'react'\nimport { View , Button, Text, TextInput, FlatList, TouchableOpacity, StyleSheet} from 'react-native'\n\ninterface %sProps {\n\n}\n\nexport const %s: React.FC<%sProps> = ({}) => {\n\treturn(\n\t\t<View>\n\n\t\t</View>\n\t)\n}", fileName, fileName, fileName)
	case constants.JavaScript:
		extension = constants.JavaScriptExtension
		componentTemplate = fmt.Sprintf("import React from 'react'\nimport { View , Button, Text, TextInput, FlatList, TouchableOpacity, StyleSheet} from 'react-native'\n\nexport const %s = ({}) => {\n\treturn(\n\t\t<View>\n\n\t\t</View>\n\t)\n}", fileName)
	default:
		extension = constants.TypescriptExtension
	}

	styleTemplate := fmt.Sprintf("import {Dimensions, StyleSheet} from 'react-native'\n\nconst {width, height} = Dimensions\n\nexport const %sStyles = StyleSheet.create({\n\n})", fileName)

	componentFileNameTemplate := fmt.Sprintf("%s/%s.%s", dir, fileName, extension)
	stylesFileNameTemplate := fmt.Sprintf("%s/style.%s", dir, extension)

	// Component File
	componentFile, err := os.Create(componentFileNameTemplate)
	// Styles File
	stylesFile, err := os.Create(stylesFileNameTemplate)

	if err != nil {
		log.Println(constants.CreateFileError)
		return err
	}

	// Components File Writer
	componentFile.Write([]byte(componentTemplate))
	// Styles File Writer
	stylesFile.Write([]byte(styleTemplate))
	return nil
}

func GenerateRNTemplate(dirName string, template string) error {
	spinner := spinner.InitSpinner("Preparing Files", "Template Generated", "✅", 11, 150)

	spinner.Start()

	dirName, err := generateOverallFile(dirName)
	if err != nil {
		log.Println(err)
		return err
	}

	err = generateFiles(dirName, dirName, template)

	if err != nil {
		log.Println(err)
		return err
	}

	time.Sleep(1 * time.Second)

	defer func() {
		spinner.Stop()
	}()

	return nil
}
