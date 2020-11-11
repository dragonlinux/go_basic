package main

import (
	"bytes"
	"fmt"
)
import bbccdd "github.com/yerden/go-util/bcd"

func main() {
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

	//dec := bbccdd.NewDecoder((*bbccdd.BCD)(Telephony))
	//
	//src := []byte{0x21, 0x43, 0xf5}
	//dst := make([]byte, bbccdd.DecodedLen(len(src)))
	//n, err := dec.Decode(dst, src)
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(dst[:n]))

	enc := bbccdd.NewEncoder((*bbccdd.BCD)(Standard))

	src := []byte("1234")
	dst := make([]byte, bbccdd.EncodedLen(len(src)))
	n, err := enc.Encode(dst, src)
	if err != nil {
		return
	}

	fmt.Println(bytes.Equal(dst[:n], []byte{0x12, 0x34}))
	// Output: true

}
