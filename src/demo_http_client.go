package main

import (
	"go_basic/http"
)

func main() {
	//nc -l 12345
	//ref:
	//https://blog.csdn.net/xiaoyida11/article/details/82659017

	http.SendHttpReq()
	http.SendGet("http://192.168.1.66:12345", "dragonlinux")
	http.SendPost("http://192.168.1.66:12345", "dragonlinux")
	http.SendPut("http://192.168.1.66:12345", "dragonlinux")
}
