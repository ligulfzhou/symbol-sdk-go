package utils

func GenByteArray(item, count int) []byte {
	res := make([]byte, count)
	idx := 0
	for idx < count {
		res[idx] = byte(item)
		idx = idx + 1
	}

	return res
}
