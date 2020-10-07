package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string
	Sex    int
	Height int
}

func main() {
	str := `{"name": "shanhuhai", "sex": 1,"height": 175}`
	student := Student{}
	json.Unmarshal([]byte(str), &student)
	fmt.Println(student.Name, student.Sex, student.Height)
}
