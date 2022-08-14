package constants

const StyleTemplate = `import {Dimensions, StyleSheet} from "react-native"

const {width, height} = Dimensions

const styles = StyleSheet.create({
	styleOne: {

	}
})`

const ComponentTemplate = `import React from 'react'
import { View , Button, Text, TextInput, FlatList, TouchableOpacity, StyleSheet} from 'react-native'

interface Props {

}

export const style: React.FC<Props> = ({}) => {
  return (
    <View>
      
    </View>
  )
}`
