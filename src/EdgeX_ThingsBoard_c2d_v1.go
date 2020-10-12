package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("dragonlinux")

	yamlFile, err := ioutil.ReadFile("./src/thingsboard_provide.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	log.Println("yamlFile:", yamlFile)

}
