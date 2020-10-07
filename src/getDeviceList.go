package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"reflect"
)

func sendHttpRequest(url string) {
	//resp, err := http.Get("http://dragonlinux.cn/myip")
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("--->", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("--->", reflect.TypeOf(body))
	//fmt.Println("--->", body)
	fmt.Printf("--->%s", body)
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

func getCommands(body []uint8) {

	{ //所以这么做，是因为原来的字符串最外层是array，去掉就是object了
		body = body[1 : len(body)-1]
	}
	result := gjson.Get(string(body), "commands")

	fmt.Println(reflect.TypeOf(result))

	if result.IsArray() {
		for _, name := range result.Array() {
			println(name.String())
		}
	}

}

func OperatingPlatform() {
	//sendHttpRequest("http://localhost:48082/api/v1/device")
	res := getHttpRes("http://localhost:48082/api/v1/device")

	getCommands(res)
}

func main() {
	OperatingPlatform()
}
