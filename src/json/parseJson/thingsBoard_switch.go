package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//{"method":"setValue","params":true}
	jsonStr := "{\"method\":\"setValue\",\"params\":true}"
	var mapValueFromThingsBoard map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapValueFromThingsBoard); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapValueFromThingsBoard)
	fmt.Println(mapValueFromThingsBoard["method"])
	fmt.Println(mapValueFromThingsBoard["params"])

	if mapValueFromThingsBoard["method"] == "setValue" {
		fmt.Println(mapValueFromThingsBoard["params"])
	}

}
