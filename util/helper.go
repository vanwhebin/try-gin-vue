package util

import (
	"math/rand"
	"time"
)

func RandomString(n int) string {
	var letters = []byte("qwertyuiplkjhgfdsazvxbcnmMNHJUYTREFGASDFGHKJ")
	str := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i, _ := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}

	return string(str)
}
