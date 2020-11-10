package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
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
