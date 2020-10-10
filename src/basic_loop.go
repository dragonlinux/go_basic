package main

import (
	"fmt"
)

func loopSimple() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	return
}

func loopKeyValue() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println(key, value)
	}
	return
}

func main() {
	loopSimple()
	fmt.Println("========")
	loopKeyValue()

}
