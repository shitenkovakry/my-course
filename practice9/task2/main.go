package main

import "fmt"

const (
	maxInt8      int8   = 1<<7 - 1
	maxInt8str          = "int8"
	maxUint8     uint8  = 1<<8 - 1
	maxUint8str         = "uint8"
	maxInt16     int16  = 1<<15 - 1
	maxInt16str         = "int16"
	maxUint16    uint16 = 1<<16 - 1
	maxUint16str        = "uint16"
	maxInt32     int32  = 1<<31 - 1
	maxInt32str         = "int32"
	maxUint32    uint32 = 1<<32 - 1
	maxUint32str        = "uint32"
	maxInt64     int64  = 1<<63 - 1
	maxInt64str         = "int64"
	maxUint64    uint64 = 1<<64 - 1
	maxUint64str        = "uint64"
)

func WhichDataType(number1, number2 int64) string {
	result := number1 * number2

	if result < 0 && -result <= int64(maxInt8)+1 {
		return maxInt8str
	}

	if result >= 0 && result <= int64(maxUint8) {
		return maxUint8str
	}

	if result < 0 && -result <= int64(maxInt16)+1 {
		return maxInt16str
	}

	if result >= 0 && result <= int64(maxUint16) {
		return maxUint16str
	}

	if result < 0 && -result <= int64(maxInt32)+1 {
		return maxInt32str
	}

	if result >= 0 && result <= int64(maxUint32) {
		return maxUint32str
	}

	if result >= 0 {
		return maxUint64str
	}

	return maxInt64str
}

func main() {
	fmt.Println("введите первое число:")
	var number1 int64
	fmt.Scan(&number1)

	fmt.Println("введите вотрое число:")
	var number2 int64
	fmt.Scan(&number2)

	minSizeInteger := WhichDataType(number1, number2)
	fmt.Printf("%d %d результат %s\n", number1, number2, minSizeInteger)

}
