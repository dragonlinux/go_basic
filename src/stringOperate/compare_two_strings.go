package main

import "fmt"

func compareTwoString() {
	a := "dragon"
	b := "dragon"

	if a == b {
		fmt.Println("=")
	} else {
		fmt.Println("!=")
	}
}

func main() {
	compareTwoString()
}
