package smconv

func StringToNullByte(s string) (b []byte) {
	return append([]byte(s), 0)
}

func StringFromNullByte(b []byte) string {
	i := 0
	for i = 0; i < len(b); i++ {
		if b[i] == 0 {
			break;
		}
	}
	if b[i] == 0 {
		return string(b[:i])
	}
	return string(b)
}

