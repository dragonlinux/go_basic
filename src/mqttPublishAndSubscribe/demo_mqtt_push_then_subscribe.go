package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"go_basic/mymqtt"
	"net/url"
	"time"
)

func d2c(client mqtt.Client) () {
	var qos = byte(1)
	var topic = "DataTopic"

	fmt.Println("d2c")
	//time.Sleep(10000 * time.Millisecond)

	for {
		client.Publish(topic, qos, false, "dragonlinux d2c")
		time.Sleep(1000 * time.Millisecond)
	}
	time.Sleep(100000000 * time.Millisecond)
}

func thingsBoardRunCommandHandler(usernameAsToken string) {

	//modify file of host all by yourself
	var brokerUrl = "192.168.1.190"

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

	client, err := mymqtt.CreateMqttClientSubscribe(mqttClientId, uri) //id必须要不一样才能正常接收
	defer client.Disconnect(5000)
	if err != nil {
		fmt.Println(err)
	}

	client.Publish(topic, qos, false, "dragonlinux connect")

	client.Publish(topic, qos, false, "dragonlinux publish")
	go d2c(client)

	token := client.Subscribe(topic, byte(qos), thingsBoardOnCommandReceivedFromBroker)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {}
}

func thingsBoardOnCommandReceivedFromBroker(client mqtt.Client, message mqtt.Message) {
	{
		{
			optionsReader := client.OptionsReader()
			fmt.Println("Username:", optionsReader.Username(), "\tClientID:", optionsReader.ClientID())
		}
		{
			//fmt.Println(message.Payload())
			fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
		}
	}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -u admin -P public -i dragon -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -u admin -P public -i dragon -t "DataTopic" -v
	//operator()
	//runCommandHandler(1)
	thingsBoardRunCommandHandler("6vkBgngn2ah7EcHM7mqP")
}
