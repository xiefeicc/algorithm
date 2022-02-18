package p7

import (
	"math"
	"strconv"
)

// Reverse 整数反转
func Reverse(x int) int {
	var result int
	for {
		l := x % 10
		result = result*10 + l
		x = x / 10
		if x == 0 {
			break
		}
	}
	if x < 0 {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

func Reverse2(x int) int {
	tmp := x
	if tmp < 0 {
		tmp = -x
	}
	xArr := []byte(strconv.Itoa(tmp))
	for l, r := 0, len(xArr)-1; l < r; l, r = l+1, r-1 {
		xArr[l], xArr[r] = xArr[r], xArr[l]
	}
	result, err := strconv.Atoi(string(xArr))
	if err != nil {
		return 0
	}
	if x < 0 {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}
