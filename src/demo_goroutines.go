package main

import (
	"fmt"
	"time"
)

var COUNT int

func hello() {
	fmt.Println("Hello world goroutine")
	for {
		count++
		fmt.Println("Hello world goroutine ", COUNT)
		time.Sleep(10 * time.Millisecond)
	}
}

func helloInt(i int) {
	fmt.Println("Hello world goroutine", i)
	for {
		time.Sleep(10 * time.Millisecond)
	}
}

func helloString(str string) {
	fmt.Println("Hello world goroutine", str)
	for {
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {

	z := 1
	if z == 0 {
		i := 0
		go helloInt(i)
		i = 1
		go helloInt(i)
	}

	{
		str := "dragon"
		go helloString(str)
		str = "linux"
		go helloString(str)
	}

	dragon = 1
	fmt.Println("main function")
	select {}

}
