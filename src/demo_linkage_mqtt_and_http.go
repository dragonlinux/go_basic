package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	http2 "go_basic/http"
	"go_basic/mymqtt"
	"net/url"
)

func runCommandHandler(i int) {

	var brokerUrl = "debug_mqtt_broker.com"
	var brokerPort = 1883
	var username = "admin"
	var password = "public"
	//var mqttClientId = "sub"
	var qos = 1
	var topic = "DataTopic"

	uri := &url.URL{
		Scheme: "tcp",
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	//client, err := createMqttClient_subscribe(mqttClientId, uri)
	client, err := mymqtt.CreateMqttClientSubscribe(fmt.Sprintf("%v", i), uri) //id必须要不一样才能正常接收
	defer client.Disconnect(5000)
	if err != nil {
		fmt.Println(err)
	}

	token := client.Subscribe(topic, byte(qos), onCommandReceivedFromBroker)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {}
}

func onCommandReceivedFromBroker(client mqtt.Client, message mqtt.Message) {
	{
		optionsReader := client.OptionsReader()
		fmt.Println(optionsReader.Username())
		fmt.Println(optionsReader.ClientID())
	}
	{
		//fmt.Println(message.Payload())
		fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
	}
	http2.SendHttpReq()
	//var request map[string]interface{}
}

//以下两种情况皆有问题,so goroutine negative.
func operator() {
	for i := 0; i < 10; i++ {
		fmt.Println("--->", i)
		go runCommandHandler(i)
	}

	select {}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	operator()
	//runCommandHandler(1)
}
