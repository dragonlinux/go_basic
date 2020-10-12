package main

import (
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

func OperatingPlatform1() {

	yamlFile, err := ioutil.ReadFile("./device.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("yamlFile:", yamlFile)

	getCommands1(yamlFile)
}

func main() {
	OperatingPlatform1()
}
