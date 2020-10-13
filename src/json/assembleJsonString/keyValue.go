package main

import (
	"encoding/json"
	"fmt"
)

func createKeyValueJson(keyStr string, in interface{}) string {
	data := make(map[string]interface{})

	//key := "SwitchB"
	//value := false

	//data[key] = value
	data[keyStr] = in
	//fmt.Println(data)

	mapString := make(map[string]string)
	for key, value := range data {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		mapString[strKey] = strValue
	}
	//fmt.Printf("\n%#v\n", mapString)

	jsonString, _ := json.Marshal(mapString)
	//fmt.Println(string(jsonString))
	return string(jsonString)
}

func main() {
	//value := false
	//value := 1.2
	value := 1
	ret := createKeyValueJson("SwitchB", value)
	fmt.Println(ret)
}
