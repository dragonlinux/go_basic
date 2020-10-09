package main

import (
	"encoding/json"
	"fmt"
)

func createJsonMethod1() {
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

func main() {
	createJsonMethod1()
}
