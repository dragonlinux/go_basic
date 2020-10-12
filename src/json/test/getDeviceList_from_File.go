package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"reflect"
)

func getCommands1(body []uint8) {

	{ //所以这么做，是因为原来的字符串最外层是array，去掉就是object了
		body = body[1 : len(body)-1]
	}
	result := gjson.Get(string(body), "commands")

	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result.IsArray())

	count := 0

	if result.IsArray() {
		for _, name := range result.Array() {
			println(name.String())
			count++
		}
	}

	fmt.Println("array count is :", count)
}

func parseJsonArray(jsonStr []uint8) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}
	//fmt.Println(val)
	//fmt.Println(reflect.TypeOf(val))
	for i, content := range val {
		if i == 0 {
			fmt.Println(i, content)
			fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			fmt.Println(content["name"])
			fmt.Println(content["commands"])
			//for i2, contentCommand := range content["commands"] {
			//	fmt.Println(i2, contentCommand)
			//}
		}
	}
}

func OperatingPlatform1() {

	//path := "./device.json"
	path := "./device_multi_array.json"

	uint8Result, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("uint8Result:", uint8Result)
	//getCommands1(uint8Result)

	parseJsonArray(uint8Result)
}

func main() {
	OperatingPlatform1()
}
