package utils

import (
	"fmt"
	"log"
	"os"
)

const (
	filePermissions         = os.O_APPEND | os.O_CREATE | os.O_TRUNC | os.O_RDWR
	fileUnixPermission      = 0644
	directoryUnixPermission = 755
)

var (
	bytesWritten = 0
)

func handleStdErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func Component(name, domain string) {
	// Generate Directory
	directory(name)
	err := os.Chdir(name)
	handleStdErr(err)

	style(name)
	component(name, domain)
}

func style(name string) {
	content := fmt.Sprintf("import {Dimensions, StyleSheet} from \"react-native\";\n\nconst {width, height} = Dimensions.get(\"screen\")\n\nexport const %sStyle = StyleSheet.create({\n   \n})", name)

	f, err := os.OpenFile(fmt.Sprintf("%s.style.tsx", name), filePermissions, fileUnixPermission)
	defer func(f *os.File) {
		err := f.Close()
		handleStdErr(err)
	}(f)
	handleStdErr(err)

	n, err := f.WriteString(content)
	handleStdErr(err)

	bytesWritten += n
}

func directory(name string) {
	err := os.Mkdir(name, directoryUnixPermission)
	handleStdErr(err)
}

func component(name, domain string) {
	f, err := os.OpenFile(fmt.Sprintf("%s.%s.tsx", name, domain), filePermissions, fileUnixPermission)
	defer func(f *os.File) {
		err := f.Close()
		handleStdErr(err)
	}(f)
	handleStdErr(err)

	content := fmt.Sprintf("import  {FC,useState, useCallback, useRef, useMemo, useEffect} from \"react\";\nimport {Text, View, FlatList, TextInput, Image, Dimensions, SafeAreaView} from \"react-native\";\n\nimport {%sStyle} from './%s.style'\n\ninterface Props {\n    \n}\n\nexport const %sScreen: FC<Props> = ({}) => {\n    return (\n        <View >\n            \n        </View>\n    )\n}\n", name, name, name)

	n, err := f.WriteString(content)
	handleStdErr(err)

	bytesWritten += n
}
