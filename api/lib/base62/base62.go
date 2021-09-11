package base62

import rand2 "math/rand"

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	length   = int(len(alphabet))
)

func Encode(sizeHash int) string {
	if sizeHash == 0 {
		return string(alphabet[0])
	}

	s := ""
	for i := 0; i <= sizeHash; i++ {
		var rand = rand2.Intn(length)
		s = string(alphabet[rand]) + s
	}
	return s
}