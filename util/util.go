package util

import (
	"math/rand"
)

func RandomString(n int) string {
	var letters = []byte("asdfsadfasdfASDFQWERF")
	result := make([]byte, n)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
