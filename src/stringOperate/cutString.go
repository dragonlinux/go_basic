package main

import "fmt"

func main() {

	str := "v1/devices/me/rpc/request/108"
	str = "v1/devices/me/rpc/request/"
	//str = "v1/"

	fmt.Println("length:", len(str))
	fmt.Println(str)
	asciiSubstring := str[26:]
	fmt.Println(asciiSubstring)

}
