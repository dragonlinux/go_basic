package main

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/tidwall/gjson"
	http2 "go_basic/http"
	"time"

	//http2 "go_basic/http"
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

func runCommandHandler(i int) {

	var brokerUrl = "192.168.1.190"
	var brokerPort = 1883
	var username = "admin"
	var password = "public"
	//var mqttClientId = "sub"
	var qos = 1
	var topic = "acturatorTopicSwitchA"

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

func runCommandHandlerSwitchB(i int) {

	var brokerUrl = "192.168.1.190"
	var brokerPort = 1883
	var username = "admin"
	var password = "public"
	//var mqttClientId = "sub"
	var qos = 1
	var topic = "acturatorTopicSwitchB"

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

	token := client.Subscribe(topic, byte(qos), onCommandReceivedFromBrokerSwitchB)
	if token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
	}

	select {}
}

func filterOperator(jsonStr string, filterString string) (url string, param string, flag bool) {
	//fmt.Println("filterOperator ", jsonStr)
	{
		result := gjson.Get(string(jsonStr), "commands")

		//fmt.Println(reflect.TypeOf(result))
		//fmt.Println(result.IsArray())

		if result.IsArray() {
			for _, name := range result.Array() {
				//println(i, name.String())
				{
					result1 := gjson.Get(string(name.String()), "name")
					//fmt.Println("name ..........>", result1)

					if result1.String() == filterString {
						result1 = gjson.Get(string(name.String()), "put")
						//fmt.Println("put  ..........>", result1)

						//fmt.Println(reflect.TypeOf(result1))

						result2 := result1
						result1 = gjson.Get(result1.String(), "url")
						//fmt.Println("url  ..........>", result1)
						url := result1.String()

						result1 = gjson.Get(result2.String(), "parameterNames")

						//fmt.Println("parameterNames length.>", len(result1.Array()))
						if len(result1.Array()) == 1 {
							//fmt.Println("parameterNames.>", result1)
							for _, r := range result1.Array() {
								//fmt.Println(i, r, reflect.TypeOf(r.String()))
								return url, r.String(), true
							}
							//param := result1.Array()[0]
						}
						fmt.Println("")
					}
				}
			}
		}
	}
	return "", "", false
}

func getDeviceName(jsonStr []uint8, deviceName string) (retString string, flag bool) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
		//fmt.Println(content["name"])
		if content["name"] == deviceName {
			//fmt.Println(i, content)
			//fmt.Println(reflect.TypeOf(content))
			fmt.Println(content["id"])
			//fmt.Println(content["name"])
			//fmt.Println(content["commands"])
			//fmt.Println(reflect.TypeOf(content["commands"]))
			//fmt.Println("==============>")

			johnJSON, err := json.Marshal(content)
			if err != nil {
				fmt.Println("error:", err)
			}
			//fmt.Println("再转换成json string+++++++++>>>", string(johnJSON), err)

			return string(johnJSON), true
		}
	}
	return "", false
}

func onCommandReceivedFromBroker(client mqtt.Client, message mqtt.Message) {
	//{
	//	optionsReader := client.OptionsReader()
	//	fmt.Println(optionsReader.Username())
	//	fmt.Println(optionsReader.ClientID())
	//}
	{
		fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
	}

	uint8Result := getHttpRes("http://localhost:48082/api/v1/device")

	//fmt.Println(string(uint8Result))
	{
		deviceName := "SerialServer"
		retJson, flag := getDeviceName(uint8Result, deviceName)
		//retJson, flag := getDeviceName(uint8Result, "Modbus_RTU_test_device_ADAM")
		if flag != true {
			fmt.Println("DeviceName not exist:", deviceName)
			for {
				fmt.Print(".")
				time.Sleep(1000 * time.Millisecond)
			}
			return
		}
		//fmt.Println(retJson, flag)

		operator := "SwitchA"
		url, param, flag := filterOperator(retJson, operator)
		//if flag != true {
		//	fmt.Println("DeviceName not exist:", deviceName)
		//	for {
		//		fmt.Println("after filterOperator")
		//		time.Sleep(1000 * time.Millisecond)
		//	}
		//	return
		//}
		fmt.Println("go routine get :rul:", url, "\tparam:", param)
		prepareSendValue := createKeyValueJson(param, string(message.Payload()))
		fmt.Println(prepareSendValue)
		http2.SendPut(url, prepareSendValue)
	}

}

func onCommandReceivedFromBrokerSwitchB(client mqtt.Client, message mqtt.Message) {
	//{
	//	optionsReader := client.OptionsReader()
	//	fmt.Println(optionsReader.Username())
	//	fmt.Println(optionsReader.ClientID())
	//}
	{
		fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
	}

	uint8Result := getHttpRes("http://localhost:48082/api/v1/device")

	//fmt.Println(string(uint8Result))
	{
		deviceName := "SerialServer"
		retJson, flag := getDeviceName(uint8Result, deviceName)
		//retJson, flag := getDeviceName(uint8Result, "Modbus_RTU_test_device_ADAM")
		if flag != true {
			fmt.Println("DeviceName not exist:", deviceName)
			for {
				fmt.Print(".")
				time.Sleep(1000 * time.Millisecond)
			}
			return
		}
		//fmt.Println(retJson, flag)

		operator := "SwitchB"
		url, param, flag := filterOperator(retJson, operator)
		//if flag != true {
		//	fmt.Println("DeviceName not exist:", deviceName)
		//	for {
		//		fmt.Println("after filterOperator")
		//		time.Sleep(1000 * time.Millisecond)
		//	}
		//	return
		//}
		fmt.Println("go routine get :rul:", url, "\tparam:", param)
		prepareSendValue := createKeyValueJson(param, string(message.Payload()))
		fmt.Println(prepareSendValue)
		http2.SendPut(url, prepareSendValue)
	}

}

func createKeyValueJson(keyStr string, in interface{}) string {
	data := make(map[string]interface{})

	//key := "SwitchB"
	//value := false

	//data[key] = value
	data[keyStr] = in
	//fmt.Println(data)

	mapString := make(map[string]string)
	for key, value := range data {
		strKey := fmt.Sprintf("%v", key)
		strValue := fmt.Sprintf("%v", value)
		mapString[strKey] = strValue
	}
	//fmt.Printf("\n%#v\n", mapString)

	jsonString, _ := json.Marshal(mapString)
	//fmt.Println(string(jsonString))
	return string(jsonString)
}

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

func operator() {
	for i := 0; i < 2; i++ {
		fmt.Println("--->", i)
		go runCommandHandler(i)
	}
	select {}
}

func main() {
	//mosquitto_pub -h 192.168.1.190 -t "DataTopic" -m "Hello MQTT1"
	//mosquitto_sub -h 192.168.1.190 -t "DataTopic" -v
	//operator()
	//runCommandHandler(1)

	{
		go runCommandHandler(1)
		go runCommandHandlerSwitchB(2)
		select {}
	}

	//thingsBoardRunCommandHandler("6vkBgngn2ah7EcHM7mqP")
}
