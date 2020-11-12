package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func test1() {
	buff := make([]byte, 1024)

	buff[0] = 0xff
	buff[1] = 0xee
	buff[2] = 0xdd
	buff[3] = 0xcc
	buff[4] = 0xbb
	buff[5] = 0xaa

	fmt.Println("dragon", hex.EncodeToString(buff[:5]))
	fmt.Println("dragon", hex.EncodeToString(buff[:6]))
}

func test2() {
	byteNumber := []byte("1")
	fmt.Println(len(byteNumber))
	fmt.Println(byteNumber[:1])
	fmt.Println(bytes.Equal(byteNumber[:1], []byte{0x31}))
	fmt.Println(bytes.Equal(byteNumber[:1], []byte{0x31, 0x32}))
}

func main() {
	test1()
	test2()
}
