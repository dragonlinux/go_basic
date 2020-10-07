package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	str := `{"name": "shanhuhai", "sex": 1,"height": 175, "classmate": ["王小五","赵小六","白小七"]}`

	name := gjson.Get(str, "name")
	classmate := gjson.Get(str, "classmate")
	fmt.Println(name)

	if classmate.IsArray() {
		fmt.Println(classmate.Array()[0])
		fmt.Println(classmate.Array()[1])
		fmt.Println(classmate.Array()[2])
	}
}
