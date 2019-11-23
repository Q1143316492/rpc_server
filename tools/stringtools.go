package tools

func ByteSliceSplit(data []byte, sign byte) [][]byte {
	var ret [][]byte
	var tmp []byte
	for _, v := range data {
		if v != sign {
			tmp = append(tmp, v)
		} else {
			ret = append(ret, tmp)
			tmp = []byte{}
		}
	}
	if len(tmp) != 0 {
		ret = append(ret, tmp)
	}
	return ret
}
