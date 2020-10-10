package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
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

func Get() {
	url := "http://192.168.1.190:12345"

	bodyBuf := bytes.NewBuffer([]byte("dragonlinux"))

	//req, _ := http.NewRequest("GET", url, nil)
	req, _ := http.NewRequest("GET", url, bodyBuf)
	//req, _ := http.NewRequest("GET", url, bodyBuf)
	req.Header.Add("Authorization", "xxxx")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func Post() {

	//url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1/data"
	url := "http://192.168.1.190:12345"

	payload := strings.NewReader("a=111")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Date", "Tue, 11 Sep 2018 10:57:09 GMT")
	req.Header.Add("Authorization", "xxx")
	req.Header.Add("Content-Type", "text/plain")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func Put() {
	url := "http://192.168.1.190:12345"

	payload := strings.NewReader("dragonlinux")

	req, _ := http.NewRequest("PUT", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "xxx")
	req.Header.Add("Date", "Wed, 12 Sep 2018 02:10:09 GMT")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func main() {
	//nc -l 12345

	sendHttpReq()
	Get()
	Post()
	Put()
}
