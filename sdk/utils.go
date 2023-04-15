package sdk

import (
	"math/rand"
	"time"
)

const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	
	b := make([]byte, n)
	b[0] = '3'
	b[1] = 'A'
	for i := 2; i < n; i++ {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
