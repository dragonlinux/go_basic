package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func helloWorld() {
	id := [5]string{"1", "2", "3", "4", "5"}
	name := [5]string{"A", "B", "C", "D", "E"}

	parseData := make([]map[string]interface{}, 0, 0)

	for counter, _ := range id {
		var singleMap = make(map[string]interface{})
		singleMap["id"] = id[counter]
		singleMap["name"] = name[counter]
		parseData = append(parseData, singleMap)
	}
	b, _ := json.Marshal(parseData)
	fmt.Println(string(b))
}

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

	//jsonString, _ := json.Marshal(mapString)
	jsonString, _ := json.MarshalIndent(mapString, "", "  ")

	//fmt.Println(string(jsonString))
	fmt.Println("=========================")
	return string(jsonString)
}

func createTBAttributes() {

	voidObject := make(map[string]interface{})

	deviceList := make(map[string]interface{})
	deviceList["Device A"] = voidObject
	deviceList["Device B"] = voidObject
	fmt.Println("deviceList", deviceList)

	//jsonString, _ := json.Marshal(deviceList)
	jsonString, _ := json.MarshalIndent(deviceList, "", "  ")

	fmt.Println(string(jsonString))
	fmt.Println("=========================")
}

func createTBTelemetry() {

	valuesContent := make(map[string]interface{})
	valuesContent["temperature"] = 123
	valuesContent["humidity"] = 123

	values := make(map[string]interface{})
	values["values"] = valuesContent
	{
		t := time.Now() //It will return time.Time object with current timestamp
		tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
		fmt.Printf("timeUnixMilli: %d\n", tUnixMilli)
		values["ts"] = tUnixMilli
	}

	jsonString, _ := json.Marshal(values)
	fmt.Println(string(jsonString))

	deviceArray := make([]map[string]interface{}, 0, 0)
	deviceArray = append(deviceArray, values)

	//var deviceArray = []map[string]interface{}// <---- This must be an array to match input
	//var deviceArray = make(map[string][]map[string]interface{}) // <---- This must be an array to match input
	//deviceArray["Device A"][0] = values

	deviceObject := make(map[string]interface{})
	deviceObject["Device A"] = deviceArray
	deviceObject["Device B"] = deviceArray

	//jsonString, _ = json.Marshal(deviceObject)
	jsonString, _ = json.MarshalIndent(deviceObject, "", "  ")

	fmt.Println(string(jsonString))
	fmt.Println("=========================")
}

func main() {
	//value := false
	//value := 1.2

	helloWorld()
	value := 1
	ret := createKeyValueJson("SwitchB", value)
	fmt.Println(ret)

	createTBAttributes()
	createTBTelemetry()
}
