package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"go_basic/mymqtt"
	"io/ioutil"
	"log"
	"net/url"
	"reflect"
)

func thingsBoardRunCommandHandler(usernameAsToken string) {

	//modify file of host all by yourself
	var brokerUrl = "demo_thingsboard.com"

	var brokerPort = 1883
	var username = usernameAsToken
	var password = ""
	//var mqttClientId = "sub"
	//var mqttClientId = fmt.Sprintf("%v", i)
	var mqttClientId = ""
	var qos = 1
	var topic = "v1/devices/me/rpc/request/+"

	uri := &url.URL{
		Scheme: "tcp",
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	client, err := mymqtt.CreateMqttClientSubscribe(mqttClientId, uri) //id必须要不一样才能正常接收
	defer client.Disconnect(5000)
	if err != nil {
		fmt.Println(err)
	}

	token := client.Subscribe(topic, byte(qos), thingsBoardOnCommandReceivedFromBroker)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {}
}

func thingsBoardOnCommandReceivedFromBroker(client mqtt.Client, message mqtt.Message) {
	{
		optionsReader := client.OptionsReader()
		fmt.Println("Username:", optionsReader.Username(), "\tClientID:", optionsReader.ClientID())
	}
	{
		fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
	}
	{
		topic := "v1/devices/me/rpc/response/" + message.Topic()[26:]
		var qos = byte(1)
		client.Publish(topic, qos, false, message)
	}
}

func parseMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		fmt.Println("\tread from file:", key, ":", value)
		fmt.Println("++++", reflect.TypeOf(value))
		go thingsBoardRunCommandHandler(reflect.ValueOf(value).String())

		//go OperatingPlatform("Modbus_TCP_test_device", key, reflect.ValueOf(value).String())
	}
	select {}
}

func operator() {
	yamlFile, err := ioutil.ReadFile("./src/thingsboard_provide.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}

	log.Println("yamlFile:", yamlFile)

	{
		m := map[string]interface{}{}
		// Parsing/Unmarshalling JSON encoding/json
		err = json.Unmarshal([]byte(yamlFile), &m)
		if err != nil {
			panic(err)
		}

		parseMap(m)
	}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	operator()
}
