package main

import (
	"fmt"
)

func loop_simple() {
	for i := 0; i < 8; i++ {
		fmt.Println(i)
	}
	return
}

func loop_key_value() {
	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println(key, value)
	}
	return
}

func main() {
	loop_simple()

	loop_key_value()

}
