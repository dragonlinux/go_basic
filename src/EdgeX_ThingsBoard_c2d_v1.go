package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"reflect"
	"time"
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

func getDeviceName(jsonStr []uint8, deviceName string) (retString string, flag bool) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
		if content["name"] == deviceName {
			//fmt.Println(i, content)
			//fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			//fmt.Println(content["name"])
			//fmt.Println(content["commands"])
			//fmt.Println(reflect.TypeOf(content["commands"]))
			//fmt.Println("==============>")

			johnJSON, err := json.Marshal(content)
			if err != nil {
				fmt.Println("error:", err)
			}
			//fmt.Println("再转换成json string+++++++++>>>", string(johnJSON), err)

			return string(johnJSON), true
		}
	}
	return "", false
}

func filterOperator(jsonStr string, filterString string) (url string, param string, flag bool) {
	//fmt.Println("filterOperator ", jsonStr)
	{
		result := gjson.Get(string(jsonStr), "commands")

		//fmt.Println(reflect.TypeOf(result))
		//fmt.Println(result.IsArray())

		if result.IsArray() {
			for _, name := range result.Array() {
				//println(i, name.String())
				{
					result1 := gjson.Get(string(name.String()), "name")
					//fmt.Println("name ..........>", result1)

					if result1.String() == filterString {
						result1 = gjson.Get(string(name.String()), "put")
						//fmt.Println("put  ..........>", result1)

						//fmt.Println(reflect.TypeOf(result1))

						result2 := result1
						result1 = gjson.Get(result1.String(), "url")
						//fmt.Println("url  ..........>", result1)
						url := result1.String()

						result1 = gjson.Get(result2.String(), "parameterNames")

						//fmt.Println("parameterNames length.>", len(result1.Array()))
						if len(result1.Array()) == 1 {
							//fmt.Println("parameterNames.>", result1)
							for _, r := range result1.Array() {
								//fmt.Println(i, r, reflect.TypeOf(r.String()))
								return url, r.String(), true
							}
							//param := result1.Array()[0]
						}
						fmt.Println("")
					}
				}
			}
		}
	}
	return "", "", false
}

func OperatingPlatform(deviceName string, operator string, token string) {

	//fmt.Println("=-=-=-=-:", token)
	//path := "./device.json"
	path := "./device_multi_array.json"

	uint8Result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	//log.Println("uint8Result:", uint8Result)

	//if false {
	//	url := parseJsonArray1(uint8Result)
	//	fmt.Println("final get", url)
	//}
	{
		retJson, flag := getDeviceName(uint8Result, deviceName)
		if flag != true {
			fmt.Println("DeviceName not exist")
			for {
				fmt.Println("after getDeviceName")
				time.Sleep(1000 * time.Millisecond)
			}
			return
		}
		//fmt.Println(retJson)
		url, param, flag := filterOperator(retJson, operator)
		if flag != true {
			fmt.Println("DeviceName not exist")
			for {
				fmt.Println("after filterOperator")
				time.Sleep(1000 * time.Millisecond)
			}
			return
		}
		fmt.Println("go routine get :rul:", url, "\t", param)
	}
	for {
		fmt.Println("end")
		time.Sleep(1000 * time.Millisecond)
	}
}

func parseMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		fmt.Println("\tread from file:", key, ":", value)
		//fmt.Println("++++", reflect.TypeOf(value))

		go OperatingPlatform("Modbus_TCP_test_device", key, reflect.ValueOf(value).String())
	}
	select {}
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

	//OperatingPlatform("Modbus_TCP_test_device", "SwitchB", "")
}
