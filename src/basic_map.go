package main

import (
	"fmt"
)

func main() {

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
