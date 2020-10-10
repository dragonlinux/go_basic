package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	"reflect"
	"strings"
)

func printVal(v interface{}, depth int) {
	typ := reflect.TypeOf(v).Kind()
	if typ == reflect.Int || typ == reflect.String {
		fmt.Printf("%s%v\n", strings.Repeat(" ", depth), v)
	} else if typ == reflect.Slice {
		fmt.Printf("\n")
		printSlice(v.([]interface{}), depth+1)
	} else if typ == reflect.Map {
		fmt.Printf("\n")
		printMap(v.(map[interface{}]interface{}), depth+1)
	}

}

func printMap(m map[interface{}]interface{}, depth int) {
	for k, v := range m {
		fmt.Printf("%sKey:%s", strings.Repeat(" ", depth), k.(string))
		printVal(v, depth+1)
	}
}

func printSlice(slc []interface{}, depth int) {
	for _, v := range slc {
		printVal(v, depth+1)
	}
}

func parseYamlDemo1() {
	var data = `
Data:
    - name: "foo"
      bar1: 0
      k1: val1
      k2:
         val2
         val3
      bar2: 1
      k3: val4
      k4: val5
      k5: val5
      k6: val6
`

	m := make(map[string]interface{})

	err := yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		panic(err)
	}

	for k, v := range m {
		fmt.Printf("Key:%s ", k)
		printVal(v, 1)
	}
}

func parseYamlDemo2() {
	type Config struct {
		Foo string
		Bar []string
	}

	//filename := os.Args[1]
	var config Config
	source, err := ioutil.ReadFile("/home/dragon/workspace_go/go_basic/src/test.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(source, &config)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Value: %#v\n", config.Bar[0])
}

func parseEdgeXYaml() {

	yamlFile, err := ioutil.ReadFile("/home/dragon/workspace_edgex/edgex-developer_scripts/releases/edinburgh/compose-files/modbus/res/example/modbus.test.device.profile.yml")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	//log.Println("yamlFile:", yamlFile)

	m := make(map[string]interface{})

	err = yaml.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		panic(err)
	}
	for k, v := range m {
		fmt.Printf("Key:%s ", k)
		printVal(v, 1)
	}

}

func main() {
	parseEdgeXYaml()
	parseYamlDemo1()
	parseYamlDemo2()
}
