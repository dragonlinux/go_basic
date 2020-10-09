package main

import (
	"fmt"
)

type Interface interface {
	Get() int
	Set(int)
}

type Struct struct {
	Age int
}

func (s Struct) Get() int {
	return s.Age
}

func (s *Struct) Set(age int) {
	s.Age = age
}

func function(input Interface) {
	input.Set(10)
	fmt.Println(input.Get())
}

func main() {
	result := Struct{}
	function(&result)
}
