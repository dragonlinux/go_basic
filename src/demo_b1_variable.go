package main

import "fmt"

func testVariable() {
	var x, y int
	var stringA = "dragon"

	var (
		a, b = 11, 22
		c, d = 33, "dragon_d"
	)

	e, f := 1, 2

	fmt.Println(x, y, a, b, c, d,
		e, f,
		stringA)

	array := []int{1, 2, 3, 4, 5}
	fmt.Println(array)

	array = append(array, 13)

	fmt.Println(array)

}

func main() {

	testVariable()


}
