package main

import (
	"fmt"
	"time"
)

var flag = true
var gString = ""

func fucntion1() {

	for true {
		for flag {
			fmt.Println("sub proc running...", gString)
			time.Sleep(1 * time.Second)
		}

	}
	fmt.Println("sub proc exit")
}
func main() {
	//running := true
	go fucntion1()
	go fucntion1()
	go fucntion1()

	time.Sleep(1 * time.Second)
	gString = "dragonlinux 1"

	time.Sleep(2 * time.Second)
	flag = false
	gString = "dragonlinux 2"

	time.Sleep(3 * time.Second)
	fmt.Println("main proc exit")

}
