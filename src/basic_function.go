package main

import "fmt"

func sum(x int, y int) int {
	return x + y
}

func testSum() {
	result := sum(500, 10000)
	fmt.Println(result)
}

func functionReturnMultiValue(x int, y int) (aInt int, bInt int, cBool bool) {
	return 3, 7, false
}

func testFunctionReturnMultiValue() {
	a, b, c := functionReturnMultiValue(0, 0)
	fmt.Println(a, b, c)
}

func main() {
	testSum()
	testFunctionReturnMultiValue()
}
