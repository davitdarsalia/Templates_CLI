package reactNativeGen

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/constants"
	"log"
	"os"
)

func generateOverallFile(basePath string) (string, error) {
	err := os.Mkdir(basePath, os.ModePerm)

	if err != nil {
		log.Println(constants.CreateDirError)
		return "", err
	}

	return basePath, nil
}

func generateFiles(dir, fileName string) error {
	componentTemplate := fmt.Sprintf("import React from 'react'\nimport { View , Button, Text, TextInput, FlatList, TouchableOpacity, StyleSheet} from 'react-native'\n\ninterface %sProps {\n\n}\n\nexport const %s: React.FC<%sProps> = ({}) => {\n\treturn(\n\t\t<View>\n\n\t\t</View>\n\t)\n}", fileName, fileName, fileName)
	styleTemplate := fmt.Sprintf("import {Dimensions, StyleSheet} from 'react-native'\n\nconst {width, height} = Dimensions\n\nexport const %sStyles = StyleSheet.create({\n\n})", fileName)

	componentFileNameTemplate := fmt.Sprintf("%s/%s.tsx", dir, fileName)
	stylesFileNameTemplate := fmt.Sprintf("%s/style.tsx", dir)

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

func GenerateRNTemplate(dirName string) error {
	dirName, err := generateOverallFile(dirName)
	if err != nil {
		log.Println(err)
		return err
	}

	err = generateFiles(dirName, dirName)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
