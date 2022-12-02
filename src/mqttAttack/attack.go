
package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"net/url"
	//"time"
)

func CreateMqttClientSubscribe(clientID string, uri *url.URL) (mqtt.Client, error) {
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

func runCommandHandler(i int) {

	var brokerUrl = "192.168.1.190"
	//var brokerUrl = "52.130.92.191"

	var brokerPort = 1883
	var username = "a"
	var password = "b"
    //var username = "admin"
	//var password = "public"
	//var mqttClientId = "sub"
	var qos = 1
	//var topic = "DataTopic"
	var topic = "t"

	uri := &url.URL{
		Scheme: "tcp",
		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
		User:   url.UserPassword(username, password),
	}

	//client, err := createMqttClient_subscribe(mqttClientId, uri)
	client, err := CreateMqttClientSubscribe(fmt.Sprintf("%v", i), uri) //id必须要不一样才能正常接收
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
	//var request map[string]interface{}

	var topic = "t"
	var qos = byte(1)
	content := "11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111" +
		"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111" +
		"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"22222222222222222222222222222222222222222222222222222222222222222222222222222222222222222222" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"33333333333333333333333333333333333333333333333333333333333333333333333333333333333333333333" +
		"11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
	client.Publish(topic, qos, false, content)

}

//以下两种情况皆有问题,so goroutine negative.
func operator() {
	for i := 0; i < 128*6; i++ {
		fmt.Println("--->", i)
		go runCommandHandler(i)
		//time.Sleep(500 * time.Millisecond)

	}
	select {}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	//mosquitto_pub -h 52.130.92.191 -t "DataTopic" -m "from client"

	operator()
}
