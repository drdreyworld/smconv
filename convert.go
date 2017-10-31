package smconv

import (
	"strconv"
)

func ValueToBytes(v interface{}, length int) []byte {
	var b1 []byte

	switch v.(type) {
	case int:
		b1 = Int32ToBytes(int32(v.(int)))
	case int32:
		b1 = Int32ToBytes(v.(int32))
	case string:
		b1 = StringToNullByte(v.(string))
	default:
		// @TODO Write value type to error
		panic("Unknown value type")
	}

	if len(b1) > length {
		panic("Length of bytes more than allowed length")
	}

	if length > len(b1) {
		b1 = append(b1, make([]byte, length-len(b1))...)
	}

	return b1
}

func StringValueToBytes(v, vtype string, length int) []byte {
	var val interface{}
	var err error

	switch vtype {
	case "int", "int32":
		val, err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	case "string":
		val = v
	default:
		panic("Unknown value type")
	}

	return ValueToBytes(val, length)
}
