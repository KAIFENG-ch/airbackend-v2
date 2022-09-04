package utils

import (
	"math/rand"
	"time"
)

func Generate() string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < len(str); i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
