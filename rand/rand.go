package rand

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digitBytes = "0123456789"
const letterDigitBytes = letterBytes + digitBytes

func RandStringBytes(n int, sample string) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, n)
	for i := range b {
		b[i] = sample[rand.Intn(len(sample))]
	}
	return string(b)
}

func RandString(n int) string {
	return RandStringBytes(n, letterDigitBytes)
}

func RandLetter(n int) string {
	return RandStringBytes(n, letterBytes)
}

func RandDigit(n int) string {
	return RandStringBytes(n, digitBytes)
}
