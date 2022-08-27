package reactNativeGen

import (
	"fmt"
	"github.com/davitdarsalia/CLI_Commands/constants"
	"github.com/davitdarsalia/CLI_Commands/utils/spinner"
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

	spinner := spinner.InitSpinner("Generating Dir", "Dir Generated Successfully", "📁", 550, 11)

	spinner.Start()

	time.Sleep(1 * time.Second)

	defer func() {
		spinner.Stop()
	}()

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
	stylesFileNameTemplate := fmt.Sprintf("%s/%sStyles.%s", dir, fileName, extension)

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
	spinner := spinner.InitSpinner("Generating Template", `Template Generated Successfully

             k;double sin()
         ,cos();main(){float A=
       0,B=0,i,j,z[1760];char b[
     1760];printf("\x1b[2J");for(;;
  ){memset(b,32,1760);memset(z,0,7040)
  ;for(j=0;6.28>j;j+=0.07)for(i=0;6.28
 >i;i+=0.02){float c=sin(i),d=cos(j),e=
 sin(A),f=sin(j),g=cos(A),h=d+2,D=1/(c*
 h*e+f*g+5),l=cos      (i),m=cos(B),n=s\
in(B),t=c*h*g-f*        e;int x=40+30*D*
(l*h*m-t*n),y=            12+15*D*(l*h*n
+t*m),o=x+80*y,          N=8*((f*e-c*d*g
 )*m-c*d*e-f*g-l        *d*n);if(22>y&&
 y>0&&x>0&&80>x&&D>z[o]){z[o]=D;;;b[o]=
 ".,-~:;=!*#$@"[N>0?N:0];}}/*#****!!-*/
  printf("\x1b[H");for(k=0;1761>k;k++)
   putchar(k%80?b[k]:10);A+=0.04;B+=
     0.02;}}/*****####*******!!=;:~
       ~::==!!!**********!!!==::-
         .,~~;;;========;;;:~-.
             ..,--------,*/
`, "✅ ", 550, 11)

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
