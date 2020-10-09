package main

import (
	"encoding/json"
	"fmt"
)

func json2Struct() {
	type Student struct {
		Name   string
		Sex    int
		Height int
	}

	str := `{"name": "shanhuhai", "sex": 1,"height": 175}`
	student := Student{}
	json.Unmarshal([]byte(str), &student)
	fmt.Println(student.Name, student.Sex, student.Height)
}

func json2Map() {
	jsonStr := `
    {
        "name":"liangyongxing",
        "age":12
    }
    `
	var mapResult map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(jsonStr), &mapResult); err != nil {
		fmt.Println(err)
	}
	fmt.Println(mapResult)
}

func main() {
	json2Struct()
	fmt.Println("=======>")
	json2Map()
}
