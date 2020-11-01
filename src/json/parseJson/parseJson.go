package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func parseL1(input []byte) {
	var resultInterface map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(input), &resultInterface); err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultInterface)

	fmt.Println("=============>")

	fmt.Println(resultInterface["method"])
	fmt.Println(resultInterface["params"])

	fmt.Println("=============>")

	if resultInterface["method"] == "setValue" {
		fmt.Println(resultInterface["params"])
	}
}

func parseL2(input []byte) {
	var resultInterface map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(input), &resultInterface); err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultInterface)

	fmt.Println("=============>")

	fmt.Println(resultInterface["DeviceNames"])
	fmt.Println(resultInterface["params"])

	fmt.Println(resultInterface["params"])

	johnJSON, err := json.Marshal(resultInterface["params"])
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(johnJSON))

	var resultL2 map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(johnJSON), &resultL2); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resultL2)

	fmt.Println("=============>")

	for k, v := range resultL2 {
		fmt.Printf("key:value = %s : %s\n", k, v)
	}

	//for _, content := range val {
	//
	//}

}

func getDeviceNames() string {
	fileJsonL1, err := ioutil.ReadFile("./edgex/device_names.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("fileJsonL1:", fileJsonL1)

	var resultInterface map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(fileJsonL1), &resultInterface); err != nil {
		fmt.Println(err)
	}

	fmt.Println(resultInterface)

	fmt.Println("=============>")
	fmt.Println(resultInterface["DeviceNames"])

	deviceNames := resultInterface["DeviceNames"]

	str := fmt.Sprintf("%v", deviceNames)
	fmt.Println(str)

	return str
}

func main() {
	fileJsonL1, err := ioutil.ReadFile("./src/json/parseJson/L1.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("fileJsonL1:", fileJsonL1)
	parseL1(fileJsonL1)

	fmt.Println("=============================== L1 over ==========================================")
	fmt.Println("++++++>",getDeviceNames())

	fileJsonL2, err := ioutil.ReadFile("./src/json/parseJson/L2.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("fileJsonL2:", fileJsonL2)
	parseL2(fileJsonL2)

}
