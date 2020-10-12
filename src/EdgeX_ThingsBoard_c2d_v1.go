package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

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
}

func parseMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		fmt.Println(key, ":", value)
	}
}
