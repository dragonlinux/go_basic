package main

import (
	"fmt"
	"reflect"
	"time"
)

var count int

func hello() {
	fmt.Println("Hello world goroutine")
	for {
		count++
		fmt.Println("Hello world goroutine ", count)
		time.Sleep(1000 * time.Millisecond)
	}
}
func main() {
	start := time.Now()

	go hello()
	go hello()

	t := time.Now()

	fmt.Println("main function", reflect.TypeOf(t))

	elapsed := t.Sub(start)

	fmt.Println("main function", elapsed)
	select {}

}
