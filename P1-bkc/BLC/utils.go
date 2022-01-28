package BLC

import "strconv"

//将int64转成[]byte
func IntTopHex(data int64) []byte {
	x := strconv.FormatInt(data, 10)
	y := []byte(x)
	return y
}
