package smconv

func Int32ToBytes(i int32) []byte {
	var b byte = 0
	if i > -1 {
		b = 1
	} else {
		i = -i
	}

	return []byte{
		b,
		byte(i >> 24),
		byte(i >> 16),
		byte(i >> 8),
		byte(i),
	}
}

func Int32FromBytes(b []byte) int32 {
	i := int32(uint32(b[4]) | uint32(b[3])<<8 | uint32(b[2])<<16 | uint32(b[1])<<24)
	if b[0] == 0 {
		i = -i
	}
	return i
}

