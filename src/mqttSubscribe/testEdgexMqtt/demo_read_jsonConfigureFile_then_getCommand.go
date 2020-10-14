package main

import (
	"encoding/json"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/tidwall/gjson"
	http2 "go_basic/http"
	"go_basic/mymqtt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

func CreateMqttClientPublish(clientID string, uri *url.URL) (mqtt.Client, error) {
	fmt.Println(fmt.Sprintf("Create MQTT client and connection: uri=%v clientID=%v ", uri.String(), clientID))
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s://%s", uri.Scheme, uri.Host))
	opts.SetClientID(clientID)
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(password)

	//opts.SetConnectionLostHandler(func(client mqtt.Client, e error) {
	//	fmt.Println(fmt.Sprintf("Connection lost : %v", e))
	//	token := client.Connect()
	//	if token.Wait() && token.Error() != nil {
	//		fmt.Println(fmt.Sprintf("Reconnection failed : %v", e))
	//	} else {
	//		fmt.Println(fmt.Sprintf("Reconnection sucessful : %v", e))
	//	}
	//})

	client := mqtt.NewClient(opts)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		return client, token.Error()
	}

	return client, nil
}

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

func getDeviceName(jsonStr []uint8, deviceName string) (retString string, flag bool) {
	var val []map[string]interface{} // <---- This must be an array to match input
	if err := json.Unmarshal([]byte(jsonStr), &val); err != nil {
		panic(err)
	}

	for _, content := range val {
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

func thingsBoardRunCommandHandler(usernameAsToken string) {
	{
		// send http request, if not a operator,go to endless loop.

	}
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
	optionsReader := client.OptionsReader()
	{
		fmt.Println("Username:", optionsReader.Username(), "\tClientID:", optionsReader.ClientID())
		fmt.Println(fmt.Sprintf("Send response: %s %s", message.Topic(), message.Payload()))
	}
	{
		topic := "v1/devices/me/rpc/response/" + message.Topic()[26:]

		{
			//var brokerUrl = "demo_thingsboard.com"
			////var brokerUrl = "192.168.1.78"
			//var brokerPort = 1883
			//var username = optionsReader.Username()
			//var password = ""
			//var mqttClientId = ""
			//var qos = byte(1)
			//var topic = topic
			//
			//uri := &url.URL{
			//	Scheme: "tcp",
			//	Host:   fmt.Sprintf("%s:%d", brokerUrl, brokerPort),
			//	User:   url.UserPassword(username, password),
			//}
			//
			//client, err := CreateMqttClientPublish(mqttClientId, uri)
			//defer client.Disconnect(5000)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//
			////for {
			//client.Publish(topic, qos, false, message.Payload())
		}

		var qos = byte(1)
		client.Publish(topic, qos, false, message)
	}

	var mapValueFromThingsBoard map[string]interface{}
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(message.Payload()), &mapValueFromThingsBoard); err != nil {
		fmt.Println(err)
	}

	if mapValueFromThingsBoard["method"] == "setValue" {
		//fmt.Println(mapValueFromThingsBoard["params"])

		operator, _ := getKeyFromValue(optionsReader.Username())
		fmt.Println("operator:", operator) // this is the operator that I want to tell edgex

		uint8Result := getHttpRes("http://localhost:48082/api/v1/device")

		{
			retJson, flag := getDeviceName(uint8Result, "Modbus_TCP_test_device")
			if flag != true {
				fmt.Println("DeviceName not exist")
				for {
					fmt.Println("after getDeviceName")
					time.Sleep(1000 * time.Millisecond)
				}
				return
			}
			//fmt.Println(retJson)
			url, param, flag := filterOperator(retJson, operator)
			if flag != true {
				fmt.Println("DeviceName not exist")
				for {
					fmt.Println("after filterOperator")
					time.Sleep(1000 * time.Millisecond)
				}
				return
			}
			fmt.Println("go routine get :rul:", url, "\tparam:", param)
			prepareSendValue := createKeyValueJson(param, mapValueFromThingsBoard["params"])
			fmt.Println(prepareSendValue)
			http2.SendPut(url, prepareSendValue)
		}
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

func getKeyFromValue(gvalue string) (string, bool) {
	yamlFile, err := ioutil.ReadFile("./src/thingsboard_provide.json")
	if err != nil {
		log.Fatalf("cannot unmarshal data: %v", err)
	}
	//log.Println("yamlFile:", yamlFile)

	m := map[string]interface{}{}
	// Parsing/Unmarshalling JSON encoding/json
	err = json.Unmarshal([]byte(yamlFile), &m)
	if err != nil {
		panic(err)
	}

	for key, value := range m {
		//fmt.Println("\tread from file:", key, ":", value)
		//fmt.Println("++++", reflect.TypeOf(value))

		if gvalue == value {
			return key, true
		}
		//go OperatingPlatform("Modbus_TCP_test_device", key, reflect.ValueOf(value).String())
	}

	return "", false
}

func parseMap(aMap map[string]interface{}) {
	for key, value := range aMap {
		fmt.Println("\tread from file,key:", key, " value:", value)
		fmt.Println("++++", reflect.ValueOf(value))
		go thingsBoardRunCommandHandler(reflect.ValueOf(value).String())

		//go OperatingPlatform("Modbus_TCP_test_device", key, reflect.ValueOf(value).String())
	}
	select {}
}

func operator() {
	//c2d 只能放执行器
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
