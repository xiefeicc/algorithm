package p6

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	convertArr := make(map[int][]byte)
	var row int
	var isReverse bool
	for i := 0; i < len(s); i++ {
		convertArr[row] = append(convertArr[row], s[i])
		if !isReverse {
			row++
		} else {
			row--
		}
		if row == numRows-1 || row == 0 {
			isReverse = !isReverse
		}
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		arr := convertArr[i]
		for _, b := range arr {
			result = append(result, b)
		}
	}
	return string(result)
}
