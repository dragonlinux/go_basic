package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

func sendHttpReq() {
	resp, err := http.Get("http://dragonlinux.cn/myip")

	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("--->", resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println("--->", reflect.TypeOf(body))
	fmt.Println("--->", body)
	fmt.Printf("--->%s", body)
}

func sendHttpReqWithJsonBody() {

}


func main() {
	sendHttpReq()
}
