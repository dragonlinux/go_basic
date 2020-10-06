package main

//import (
//	"io/ioutil"
//	"log"
//)
//
//func main() {
//
//	yamlFile, err := ioutil.ReadFile("/home/dragon/workspace_edgex/edgex-developer_scripts/releases/edinburgh/compose-files/modbus/res/example/modbus.test.device.profile.yml")
//	if err != nil {
//		log.Fatalf("cannot unmarshal data: %v", err)
//	}
//
//	log.Println("yamlFile:", yamlFile)
//
//}
import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Foo string
	Bar []string
}

func main() {
	//filename := os.Args[1]
	var config Config
	source, err := ioutil.ReadFile("/home/dragon/workspace_go/go_tutorial/src/test.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Bar[0])
}