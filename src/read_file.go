package main

import (
	"io/ioutil"
	"log"
)

func main() {

	yamlFile, err := ioutil.ReadFile("/home/dragon/workspace_edgex/edgex-developer_scripts/releases/edinburgh/compose-files/modbus/res/example/modbus.test.device.profile.yml")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	log.Println("yamlFile:", yamlFile)

}
