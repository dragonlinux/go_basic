package main

import (
	"fmt"
	"go_basic/mymqtt"
	"net/url"
	"time"
)

func mqttPublish(content string) {

	var brokerUrl = "192.168.1.80"
	var brokerPort = 1883
	var username = "admin"
	var password = "public"
	var mqttClientId = "ClientID"
	var qos = byte(1)
	var topic = "DataTopic"

	uri := &url.URL{
		Scheme: "tcp",
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	client, err := mymqtt.CreateMqttClientPublish(mqttClientId, uri)
	defer client.Disconnect(5000)
	if err != nil {
		fmt.Println(err)
	}

	//for {
	client.Publish(topic, qos, false, content)

	fmt.Println(fmt.Sprintf("Send response: %v", content))

	//time.Sleep(1000 * time.Millisecond)
	//}
}

func forLoop(i int) {
	for {
		mqttPublish(fmt.Sprintf("%v", i))
		time.Sleep(1000 * time.Millisecond)
	}
}

//以下两种情况皆有问题,so goroutine negative.
func operator() {
	for i := 0; i < 5; i++ {
		go forLoop(i)
	}

	select {}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	operator()
}
