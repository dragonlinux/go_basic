package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"reflect"
)

func parseJsonArray1(jsonStr []uint8) (url string) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
		DeviceNames := "Modbus_TCP_test_device"
		//DeviceNames := "Random-Integer-Generator01"

		if content["name"] == DeviceNames {
			//fmt.Println(i, content)
			fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			fmt.Println(content["name"])
			//fmt.Println(content["commands"])
			fmt.Println(reflect.TypeOf(content["commands"]))
			fmt.Println("==============>")

			johnJSON, err := json.Marshal(content)
			if err != nil {
				fmt.Println("error:", err)
			}
			//fmt.Println("再转换成json string+++++++++>>>", string(johnJSON), err)
			{
				result := gjson.Get(string(johnJSON), "commands")

				fmt.Println(reflect.TypeOf(result))
				fmt.Println(result.IsArray())

				count := 0

				if result.IsArray() {
					for _, name := range result.Array() {
						//println(i, name.String())
						{
							result1 := gjson.Get(string(name.String()), "name")
							fmt.Println("name ..........>", result1)

							result1 = gjson.Get(string(name.String()), "put")
							//fmt.Println("put  ..........>", result1)

							//fmt.Println(reflect.TypeOf(result1))

							result2 := result1
							result1 = gjson.Get(result1.String(), "url")
							fmt.Println("url  ..........>", result1)
							//url := result1.String()
							_ = result1.String()

							result1 = gjson.Get(result2.String(), "parameterNames")

							fmt.Println("parameterNames length.>", len(result1.Array()))
							if len(result1.Array()) == 1 {
								fmt.Println("parameterNames.>", result1)
								//return url
							}

							fmt.Println("")
						}
						count++
					}
				}
			}
		}
	}
	return ""
}

func OperatingPlatform() {

	//path := "./device.json"
	path := "./device_multi_array.json"

	uint8Result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("uint8Result:", uint8Result)
	//getCommands1(uint8Result)

	url := parseJsonArray1(uint8Result)
	fmt.Println("final get", url)
}

func parseMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		fmt.Println(key, ":", value)
	}
}

func main() {
	fmt.Println("dragonlinux")

	yamlFile, err := ioutil.ReadFile("./src/thingsboard_provide.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	log.Println("yamlFile:", yamlFile)

	{
		m := map[string]interface{}{}
		// Parsing/Unmarshalling JSON encoding/json
		err = json.Unmarshal([]byte(yamlFile), &m)
		if err != nil {
			panic(err)
		}

		parseMap(m)
	}

	OperatingPlatform()
}
