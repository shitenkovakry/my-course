package main

import "fmt"

const (
	maxUint32 uint32 = 1<<32 - 1
	maxUint16 uint16 = 1<<16 - 1
	maxUint8  uint8  = 1<<8 - 1
)

func main() {
	stopForUint8 := uint32(0)
	stopForUint16 := uint32(0)

	for index := uint32(0); index < maxUint32; index++ {
		if index > uint32(maxUint8) {
			stopForUint8++
		}

		if index > uint32(maxUint16) {
			stopForUint16++
		}
	}

	fmt.Println(stopForUint8)
	fmt.Println(stopForUint16)
}
