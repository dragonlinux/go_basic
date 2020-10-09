package main

import (
	"encoding/json"
	"fmt"
)

//struct to json
func createJson_Struct2Json() {
	// Profile declares `Profile` structure
	type Data struct {
		Infrared int `json:"infrared"`
	}

	type Frame struct {
		Datatype int  `json:"datatype"`
		Data     Data `json:"datas"`
		Msgid    int  `json:"msgid"`
	}

	var frame Frame

	var value int = 1
	frame = Frame{
		Datatype: 1,
		Data: Data{
			Infrared: value,
		},
		Msgid: 14317,
	}

	// encode `john` as JSON
	johnJSON, _ := json.MarshalIndent(frame, "", "\t")
	fmt.Println(string(johnJSON))
}

//map to json
func createJson_Map2Json() {
	type Foo struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
	}

	datas := make(map[int]Foo)

	for i := 0; i < 10; i++ {
		datas[i] = Foo{Number: 1, Title: "test"}
	}

	jsonString, err := json.Marshal(datas)

	fmt.Println(datas)
	fmt.Println(err)

	//fmt.Println(jsonString)
	fmt.Println(string(jsonString))

}

func main() {
	createJson_Struct2Json()
	createJson_Map2Json()
}
