package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"net/url"
	"time"
)

func mqttPublish(content string) {

	var brokerUrl = "192.168.1.190"
	var brokerPort = 1883
	var username = "admin"
	var password = "public"
	var mqttClientId = "ClientID"
	var qos = byte(0)
	var topic = "DataTopic"

	uri := &url.URL{
		Scheme: "tcp",
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	client, err := createMqttClient(mqttClientId, uri)
	defer client.Disconnect(5000)
	if err != nil {
		fmt.Println(err)
	}

	for {
		client.Publish(topic, qos, false, content)

		fmt.Println(fmt.Sprintf("Send response: %v", content))

		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	mqttPublish("{\"datatype\":1,\"datas\":{\"mensuo123\":11,\"weidong123\":22,\"hongwai123\":33,\"yanwu123\":44},\"msgid\":14317}")
}

func createMqttClient(clientID string, uri *url.URL) (mqtt.Client, error) {
	fmt.Println(fmt.Sprintf("Create MQTT client and connection: uri=%v clientID=%v ", uri.String(), clientID))
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s", uri.Scheme, uri.Host))
	opts.SetClientID(clientID)
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)
	opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
		fmt.Println(fmt.Sprintf("Connection lost : %v", e))
		token := client.Connect()
		if token.Wait() && token.Error() != nil {
			fmt.Println(fmt.Sprintf("Reconnection failed : %v", e))
		} else {
			fmt.Println(fmt.Sprintf("Reconnection sucessful : %v", e))
		}
	})

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return client, token.Error()
	}

	return client, nil
}
