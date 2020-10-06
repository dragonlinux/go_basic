package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println(os.Environ())
	fmt.Println(os.Getpid())
}
