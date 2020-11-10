package main

import (
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("dragon xxx")

	// connect to serverrou
	//conn, _ := net.Dial("tcp", "127.0.0.1:55555")
	conn, _ := net.Dial("tcp", "192.168.1.41:55555")
	for {

		buff := make([]byte, 13)
		buff[0] = 0x08
		buff[1] = 0x00
		buff[2] = 0x00
		buff[3] = 0x06
		buff[4] = 0x05
		buff[5] = 0x40
		buff[6] = 0x04
		buff[7] = 0x10

		buff[8] = 0x00
		buff[9] = 0x00
		buff[10] = 0x00
		buff[11] = 0x00
		buff[12] = 0x00

		//// what to send?
		conn.Write(buff)
		time.Sleep(1000 * time.Millisecond)

		for i := range buff {
			buff[i] = 0
		}

		buffRec := make([]byte, 1024)
		// wait for reply
		n, _ := conn.Read(buffRec)
		//log.Printf("Receive: %s", buff[:n])

		fmt.Print("Message from server: " + hex.EncodeToString(buffRec[:n]) + "\n")
	}
}
