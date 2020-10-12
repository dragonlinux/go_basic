package main

import (
	"encoding/json"
	"fmt"
)

const input = `
{
    "outterJSON": {
        "innerJSON1": {
            "value1": 10,
            "value2": 22,
            "InnerInnerArray": [ "test1" , "test2"],
            "InnerInnerJSONArray": [{"fld1" : "val1"} , {"fld2" : "val2"}]
        },
        "InnerJSON2":"NoneValue"
    }
} `

func main() {
	// Creating the maps for JSON
	m := map[string]interface{}{}

	// Parsing/Unmarshalling JSON encoding/json
	err := json.Unmarshal([]byte(input), &m)

	if err != nil {
		panic(err)
	}
	parseMap(m)
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}
