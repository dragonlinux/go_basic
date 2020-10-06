package main

import (
	"fmt"
)

func main() {

	//var x, y int
	//var string_a = "dragon"
	//
	//var (
	//	a, b = 11, 22
	//	c, d = 33, "dragon_d"
	//)
	//
	//e, f := 1, 2
	//
	//fmt.Println(x, y, a, b, c, d,
	//	e, f,
	//	string_a)

	//a := []int{1, 2, 3, 4, 5}
	//fmt.Println(a)
	//
	//a = append(a, 13)
	//
	//fmt.Println(a)

	a := make(map[string]int)

	a["dragon"] = 32
	a["aaa"] = 123
	a["b"] = 555

	fmt.Println(a["dragon"])

	//b := map[string]string{}
	b := make(map[string]string)

	b["aaa"] = "aaa_b"
	b["bbb"] = "bbb_b"

	fmt.Println(b["bbb"])

}
