package main

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"go_basic/mymqtt"
	"net/url"
)

func runCommandHandler(i int) {

	var brokerUrl = "192.168.1.80"
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
	client, err := mymqtt.CreateMqttClient(fmt.Sprintf("%v", i), uri) //id必须要不一样才能正常接收
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
		//fmt.Println(message.Payload())
		fmt.Println(fmt.Sprintf("Send response: %s", message.Payload()))
	}
	//var request map[string]interface{}

	//json.Unmarshal(message.Payload(), &request)
	//uuid, ok := request["uuid"]
	//if ok {
	//	log.Println(fmt.Sprintf("Command response received: topic=%v uuid=%v msg=%v", message.Topic(), uuid, string(message.Payload())))
	//
	//	if request["method"] == "set" {
	//		sendTestData(request)
	//	} else {
	//		switch request["cmd"] {
	//		case "ping":
	//			request["ping"] = "pong"
	//			sendTestData(request)
	//		case "randfloat32":
	//			request["randfloat32"] = rand.Float32()
	//			sendTestData(request)
	//		case "randfloat64":
	//			request["randfloat64"] = rand.Float64()
	//			sendTestData(request)
	//		case "message":
	//			t := time.Now()
	//			request["message"] = "test-message 时间:" + t.String()
	//			sendTestData(request)
	//		}
	//	}
	//} else {
	//	log.Println(fmt.Sprintf("Command response ignored. No UUID found in the message: topic=%v msg=%v", message.Topic(), string(message.Payload())))
	//}
}

//以下两种情况皆有问题,so goroutine negative.
func operator() {
	for i := 0; i < 10; i++ {
		fmt.Println("--->", i)
		go runCommandHandler(i)
		//time.Sleep(1000 * time.Millisecond)

		//select {}
	}

	select {}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	operator()
	//runCommandHandler(1)
}
