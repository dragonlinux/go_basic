package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/tidwall/gjson"
	"go_basic/mymqtt"
	"net/url"
	"time"
)

func thingsBoardRunCommandHandler(usernameAsToken string) {

	//modify file of host all by yourself
	//var brokerUrl = "debug_mqtt_broker.com"
	var brokerUrl = "192.168.1.71"
	var brokerPort = 1883
	var username = "C3Y3tmIpTs1LXHz3UBRP"
	var password = ""
	var mqttClientId = ""
	//var password = "public"
	//var mqttClientId = "ClientID"
	var qos = byte(1)
	var topic = "v1/gateway/rpc"

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

	strConnet := "{\n   \"device\" : \"Device A\"\n}\n"
	topic = "v1/gateway/connect"
	client.Publish(topic, qos, false, strConnet)

	topic = "v1/gateway/rpc"

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
			fmt.Println(message.Payload())
			fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
		}

		c2d_device := gjson.Get(string(message.Payload()), "device")
		fmt.Println("device", c2d_device)

		c2d_data := gjson.Get(string(message.Payload()), "data")
		fmt.Println(c2d_data)

		c2d_id := gjson.Get(c2d_data.String(), "id")
		fmt.Println("id", c2d_id)

		c2d_value := gjson.Get(c2d_data.String(), "params")
		fmt.Println("params", c2d_value)

		{
			return_data := make(map[string]bool)
			return_data["success"] = c2d_value.Bool()
			fmt.Println("return_data", return_data)
			jsonString, _ := json.Marshal(return_data)
			fmt.Println(string(jsonString))

			return_frame := make(map[string]interface{})
			return_frame["device"] = c2d_device.String()
			return_frame["id"] = c2d_id.Int()
			return_frame["data"] = return_data
			fmt.Println("return_data", return_frame)

			xxx, _ := json.Marshal(return_frame)
			fmt.Println(string(xxx))
			var qos = byte(1)
			topic := "v1/gateway/rpc"
			client.Publish(topic, qos, false, xxx)
			//return string(jsonString)
		}
		//var mapValueFromThingsBoard map[string]interface{}
		////使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
		//if err := json.Unmarshal([]byte(message.Payload()), &mapValueFromThingsBoard); err != nil {
		//	fmt.Println(err)
		//}
		//fmt.Println(mapValueFromThingsBoard["device"])
		//fmt.Println(mapValueFromThingsBoard["data"])

		//var qos = byte(1)
		//topic := "v1/gateway/rpc"
		//client.Publish(topic, qos, false, strConnet)

	}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	//operator()
	//runCommandHandler(1)
	thingsBoardRunCommandHandler("6vkBgngn2ah7EcHM7mqP")
}

func mqttPublishAlways(content string) {

	//var brokerUrl = "debug_mqtt_broker.com"
	var brokerUrl = "192.168.1.71"
	var brokerPort = 1883
	var username = "C3Y3tmIpTs1LXHz3UBRP"
	var password = ""
	var mqttClientId = ""
	//var password = "public"
	//var mqttClientId = "ClientID"
	var qos = byte(1)
	var topic = "v1/gateway/rpc"

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

	for {
		client.Publish(topic, qos, false, content)
		fmt.Println(fmt.Sprintf("Send response: %v", content))
		time.Sleep(1000 * time.Millisecond)
	}
}

//func main() {
//	mqttPublishAlways(fmt.Sprintf("%v", 2))
//}
