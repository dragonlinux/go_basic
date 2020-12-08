package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/tidwall/gjson"
	"go_basic/mymqtt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getHttpRes(url string) []uint8 {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error")
	}
	//fmt.Println("--->", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body error")
	}
	//fmt.Println("--->", reflect.TypeOf(body))
	//fmt.Println("--->", body)
	//fmt.Printf("--->%s", body)
	return body
}

func getEdgeXDeviceList() string {

	voidObject := make(map[string]interface{})
	deviceList := make(map[string]interface{})

	uint8Result := getHttpRes("http://localhost:48082/api/v1/device")

	//fmt.Println(string(uint8Result))

	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(uint8Result), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
		//fmt.Println(content)

		johnJSON, err := json.Marshal(content)
		if err != nil {
			fmt.Println("error:", err)
		}
		//fmt.Println("再转换成json string+++++++++>>>", string(johnJSON), err)

		result := gjson.Get(string(johnJSON), "commands")
		if result.IsArray() {
			for _, name := range result.Array() {
				//println(i, name.String())
				{
					result1 := gjson.Get(string(name.String()), "name")
					//fmt.Println("name ..........>", result1)
					deviceList[result1.String()] = voidObject
				}
			}
		}
	}

	jsonString, _ := json.Marshal(deviceList)
	//fmt.Println("all devices:", string(jsonString))
	return string(jsonString)
}

func thingsBoardInitRunCommandHandler(usernameAsToken string) {

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
	//var topic = "v1/gateway/rpc"

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
	var topic = "v1/gateway/connect"
	client.Publish(topic, qos, false, strConnet)

	{
		deviceListFormat := getEdgeXDeviceList()
		fmt.Println("deviceListFormat:", string(deviceListFormat))
		topic = "v1/gateway/attributes"
		client.Publish(topic, qos, false, deviceListFormat)
	}

	topic = "v1/gateway/rpc"
	token := client.Subscribe(topic, byte(qos), thingsBoardC2D)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {}
}

func thingsBoardC2D(client mqtt.Client, message mqtt.Message) {
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
	thingsBoardInitRunCommandHandler("6vkBgngn2ah7EcHM7mqP")
}

//func mqttPublishAlways(content string) {
//
//	//var brokerUrl = "debug_mqtt_broker.com"
//	var brokerUrl = "192.168.1.71"
//	var brokerPort = 1883
//	var username = "C3Y3tmIpTs1LXHz3UBRP"
//	var password = ""
//	var mqttClientId = ""
//	//var password = "public"
//	//var mqttClientId = "ClientID"
//	var qos = byte(1)
//	var topic = "v1/gateway/rpc"
//
//	uri := &url.URL{
//		Scheme: "tcp",
//		Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
//		User:   url.UserPassword(username, password),
//	}
//
//	client, err := mymqtt.CreateMqttClientPublish(mqttClientId, uri)
//	defer client.Disconnect(5000)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	for {
//		client.Publish(topic, qos, false, content)
//		fmt.Println(fmt.Sprintf("Send response: %v", content))
//		time.Sleep(1000 * time.Millisecond)
//	}
//}

//func main() {
//	mqttPublishAlways(fmt.Sprintf("%v", 2))
//}
