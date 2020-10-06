package main

//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/

import (
	"fmt"
)

func main_defer2() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
}

func main_defer1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	//main1()
	main_defer1()
	main_defer2()
}
