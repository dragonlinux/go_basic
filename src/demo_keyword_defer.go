package main

//https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/

import "fmt"

func main2() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
}

func main1() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	main1()
	main2()
}
