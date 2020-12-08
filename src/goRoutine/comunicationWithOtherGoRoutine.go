package main

import (
	"fmt"
	"time"
)

var running = true

func f() {
	for running {
		fmt.Println("sub proc running...")
		time.Sleep(1 * time.Second)
	}
	fmt.Println("sub proc exit")
}
func main() {
	//running := true
	go f()
	go f()
	go f()
	time.Sleep(2 * time.Second)
	running = false
	time.Sleep(3 * time.Second)
	fmt.Println("main proc exit")
}
