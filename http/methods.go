package http

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

func SendHttpReq() {
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

func SendGet(url string, strBody string) {

	bodyBuf := bytes.NewBuffer([]byte(strBody))

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

func SendPost(url string, strBody string) {

	//url := "http://xxxxx:8080/v2/repos/wh_flowDataSource1/data"

	payload := strings.NewReader(strBody)

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

func SendPut(url string, strBody string) {

	payload := strings.NewReader(strBody)

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
