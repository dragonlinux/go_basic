package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	bbccdd "github.com/yerden/go-util/bcd"
	"log"
	"net"
	"reflect"
	"strconv"
	"time"
)

func loopSendAndReceive() {
	fmt.Println("------------loopSendAndReceive")

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

func sendAndReceive() {
	// connect to serverrou
	//conn, _ := net.Dial("tcp", "127.0.0.1:55555")
	conn, _ := net.Dial("tcp", "192.168.1.41:55555")

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

	//for i := range buff {
	//	buff[i] = 0
	//}
	//
	//buffRec := make([]byte, 1024)
	//// wait for reply
	//n, _ := conn.Read(buffRec)
	////log.Printf("Receive: %s", buff[:n])
	//
	//fmt.Print("Message from server: " + hex.EncodeToString(buffRec[:n]) + "\n")

}

func IntToBcd(value int) int {
	return (((value / 10) % 10) << 4) | (value % 10)
}

func BcdToInt(value int) int {
	return (int)((value>>4)*10 + (value & 0x0F))
}

func forEdgeX() {
	var connEdgex net.Conn

	address := "192.168.1.41"
	port := 55555

	urlCan := fmt.Sprintf("%s:%d", address, port)
	fmt.Println("++++++++>>", urlCan)
	//connEdgex, _ = net.Dial("tcp", "192.168.1.41:55555")
	connEdgex, _ = net.Dial("tcp", urlCan)

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
	connEdgex.Write(buff)
	time.Sleep(1000 * time.Millisecond)

	for i := range buff {
		buff[i] = 0
	}

	buffRec := make([]byte, 1024)
	// wait for reply
	n, _ := connEdgex.Read(buffRec)
	//log.Printf("Receive %d bytes: %s", n, buff[:n])
	log.Printf("Receive %d bytes", n)

	fmt.Print("Message from server: " + hex.EncodeToString(buffRec[:n]) + "\n")

	connEdgex.Close()

	//fmt.Print("y: " + hex.EncodeToString(buffRec[3:4]) + "\n") 05
	//fmt.Print("x: " + hex.EncodeToString(buffRec[4:5]) + "\n") 85

	fmt.Print("xSymbol: " + hex.EncodeToString(buffRec[5:6]) + "\n")
	fmt.Print("x: " + hex.EncodeToString(buffRec[6:8]) + "\n")
	//dst := make([]byte, bcd.EncodedLen(len(buffRec[6:8])))
	{
		type BCD struct {
			// Map of symbols to encode and decode routines.
			// Example:
			//    key 'a' -> value 0x9
			Map map[byte]byte

			// If true nibbles (4-bit part of a byte) will
			// be swapped, meaning bits 0123 will encode
			// first digit and bits 4567 will encode the
			// second.
			SwapNibbles bool

			// Filler nibble is used if the input has odd
			// number of bytes. Then the output's final nibble
			// will contain the specified nibble.
			Filler byte
		}
		var (
			// Standard 8-4-2-1 decimal-only encoding.
			Standard = &BCD{
				Map: map[byte]byte{
					'0': 0x0, '1': 0x1, '2': 0x2, '3': 0x3,
					'4': 0x4, '5': 0x5, '6': 0x6, '7': 0x7,
					'8': 0x8, '9': 0x9,
				},
				SwapNibbles: false,
				Filler:      0xf}
		)
		{
			enc := bbccdd.NewEncoder((*bbccdd.BCD)(Standard))

			src := []byte("1234")
			fmt.Println("src:", src)
			dst := make([]byte, bbccdd.EncodedLen(len(src)))
			n, err := enc.Encode(dst, src)
			if err != nil {
				return
			}

			fmt.Println(bytes.Equal(dst[:n], []byte{0x12, 0x34}))
		}

		dec := bbccdd.NewDecoder((*bbccdd.BCD)(Standard))
		srcDec := []byte{0x12, 0x34}
		fmt.Println("srcDec:", srcDec)
		dstDec := make([]byte, bbccdd.DecodedLen(len(srcDec)))
		_, err := dec.Decode(dstDec, srcDec)
		if err != nil {
			return
		}
		fmt.Print("test: " + hex.EncodeToString(dstDec) + "\n")
		fmt.Print("test: ", dstDec, "\n")

		tempStr := hex.EncodeToString(dstDec)
		fmt.Print("tempStr: ", tempStr, "\n")

		fmt.Println("length tempStr:", len(tempStr))

		var xSignArray string

		for i, value := range tempStr {
			fmt.Print(" ", value, " ")
			if i == 1 {
				//xSignArray[0] = string(value)
				xSignArray = xSignArray + string(value)
			}
			if i == 3 {
				//xSignArray[1] = string(value)
				xSignArray = xSignArray + string(value)
			}
			if i == 5 {
				//xSignArray[2] = string(value)
				xSignArray = xSignArray + string(value)
			}
			if i == 7 {
				//xSignArray[3] = string(value)
				xSignArray = xSignArray + string(value)
			}
		}
		fmt.Println("\nxSignArray:", xSignArray)
		fmt.Println("xSignArray:", reflect.TypeOf(xSignArray))
		fmt.Println(len(xSignArray))
		//byteNumber := []byte(xSignArray)
		byteToInt, _ := strconv.Atoi(xSignArray)
		fmt.Println("byteToInt:", byteToInt)

		data := binary.BigEndian.Uint16((dstDec))
		fmt.Println(data)
	}

	x, _ := strconv.Atoi(string(buffRec[6:8]))
	fmt.Println(x)

	data := binary.BigEndian.Uint16((buffRec[6:8]))
	fmt.Println(data)
	ret := BcdToInt(int(data))
	fmt.Println(ret)

	{
		byteNumber := []byte("14")
		byteToInt, _ := strconv.Atoi(string(byteNumber))
		fmt.Println(byteToInt)
	}

	{
		//
		//var Telephony *BCD.BCD
		//enc := BCD.NewEncoder(Telephony)
		//
		//src := []byte("12345")
		//dst := make([]byte, BCD.EncodedLen(len(src)))
		//n, err := enc.Encode(dst, src)
		//if err != nil {
		//	return
		//}
		//
		//fmt.Println(bytes.Equal(dst[:n], []byte{0x21, 0x43, 0xf5}))
	}

	fmt.Print("ySymbol: " + hex.EncodeToString(buffRec[8:9]) + "\n")
	fmt.Print("y: " + hex.EncodeToString(buffRec[9:11]) + "\n")
}

func BCD2Int(inputBytes []byte) int {
	type BCD struct {
		// Map of symbols to encode and decode routines.
		// Example:
		//    key 'a' -> value 0x9
		Map map[byte]byte

		// If true nibbles (4-bit part of a byte) will
		// be swapped, meaning bits 0123 will encode
		// first digit and bits 4567 will encode the
		// second.
		SwapNibbles bool

		// Filler nibble is used if the input has odd
		// number of bytes. Then the output's final nibble
		// will contain the specified nibble.
		Filler byte
	}
	var (
		// Standard 8-4-2-1 decimal-only encoding.
		Standard = &BCD{
			Map: map[byte]byte{
				'0': 0x0, '1': 0x1, '2': 0x2, '3': 0x3,
				'4': 0x4, '5': 0x5, '6': 0x6, '7': 0x7,
				'8': 0x8, '9': 0x9,
			},
			SwapNibbles: false,
			Filler:      0xf}
	)
	dec := bbccdd.NewDecoder((*bbccdd.BCD)(Standard))
	//srcDec := []byte{0x12, 0x34}
	srcDec := inputBytes
	//fmt.Println("srcDec:", srcDec)
	dstDec := make([]byte, bbccdd.DecodedLen(len(srcDec)))
	_, err := dec.Decode(dstDec, srcDec)
	if err != nil {
		return 0
	}
	//fmt.Print("test: " + hex.EncodeToString(dstDec) + "\n")
	//fmt.Print("test: ", dstDec, "\n")

	tempStr := hex.EncodeToString(dstDec)
	//fmt.Print("tempStr: ", tempStr, "\n")
	//
	//fmt.Println("length tempStr:", len(tempStr))

	var xSignArray string

	for i, value := range tempStr {
		//fmt.Print(" ", value, " ")
		if i == 1 {
			//xSignArray[0] = string(value)
			xSignArray = xSignArray + string(value)
		}
		if i == 3 {
			//xSignArray[1] = string(value)
			xSignArray = xSignArray + string(value)
		}
		if i == 5 {
			//xSignArray[2] = string(value)
			xSignArray = xSignArray + string(value)
		}
		if i == 7 {
			//xSignArray[3] = string(value)
			xSignArray = xSignArray + string(value)
		}
	}
	//fmt.Println("\nxSignArray:", xSignArray)
	//fmt.Println("xSignArray:", reflect.TypeOf(xSignArray))
	//fmt.Println(len(xSignArray))
	//byteNumber := []byte(xSignArray)
	byteToInt, _ := strconv.Atoi(xSignArray)
	//fmt.Println("byteToInt:", byteToInt)

	return byteToInt
	//data := binary.BigEndian.Uint16((dstDec))
	//fmt.Println(data)
}

func forEdgeXSimply() {
	var connEdgex net.Conn

	address := "192.168.1.41"
	port := 55555

	urlCan := fmt.Sprintf("%s:%d", address, port)
	fmt.Println("++++++++>>", urlCan)
	//connEdgex, _ = net.Dial("tcp", "192.168.1.41:55555")
	connEdgex, _ = net.Dial("tcp", urlCan)

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
	connEdgex.Write(buff)
	time.Sleep(1000 * time.Millisecond)

	for i := range buff {
		buff[i] = 0
	}

	buffRec := make([]byte, 1024)
	// wait for reply
	n, _ := connEdgex.Read(buffRec)
	//log.Printf("Receive %d bytes: %s", n, buff[:n])
	log.Printf("Receive %d bytes", n)

	fmt.Print("Message from server: " + hex.EncodeToString(buffRec[:n]) + "\n")

	connEdgex.Close()

	{
		fmt.Print("xSymbol: " + hex.EncodeToString(buffRec[5:6]) + "\n")
		fmt.Print("x: " + hex.EncodeToString(buffRec[6:8]) + "\n")

		x := BCD2Int(buffRec[6:8])
		fmt.Println("BCD2Int x :", x)

		fmt.Print("ySymbol: " + hex.EncodeToString(buffRec[8:9]) + "\n")
		fmt.Print("y: " + hex.EncodeToString(buffRec[9:11]) + "\n")

		y := BCD2Int(buffRec[9:11])
		fmt.Println("BCD2Int y :", y)
	}
}

func main() {
	//loopSendAndReceive()
	//sendAndReceive()
	//forEdgeX()
	forEdgeXSimply()
}
