package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2, 3}
	rt := reflect.TypeOf(s)

	// []int, *reflect.rtype
	fmt.Println("%s, %T", rt, rt)

	// , string
	fmt.Println("%s, %T", rt.Name(), rt.Name())

	// slice, reflect.Kind
	fmt.Println("%s, %T", rt.Kind(), rt.Kind())
}
