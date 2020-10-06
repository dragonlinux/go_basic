package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func main() {

	result := sum(500, 10000)

	fmt.Println(result)
}
