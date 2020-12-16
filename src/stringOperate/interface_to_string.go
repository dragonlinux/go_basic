package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{} = "abc"
	stringx := fmt.Sprintf("%v", x)
	println(reflect.TypeOf(x))
	println(reflect.TypeOf(stringx))
	println(stringx)
}
