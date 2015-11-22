package randn

import (
	"bytes"
	"math/rand"
)

const chars = "abcdefghijklmnopqrstuvqxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_"

func Lookup(n int) string {
	res := make([]uint8, n)
	for i := 0; i < n; i++ {
		res[i] = chars[rand.Int31()&63]
	}
	return string(res)
}

func BLookup(n int) string {
	var res bytes.Buffer
	for i := 0; i < n; i++ {
		res.WriteByte(chars[rand.Int31()&63])
	}
	return res.String()
}

func ALookup(n uint8) string {
	res := make([]uint8, n)
	inds := rand.Int63()
	var i uint8
	for i = 0; i < n; i++ {
		res[i] = chars[(inds>>i*6)&63]
	}
	return string(res)
}

func AALookup(n uint8) string {
	res := make([]uint8, 16)
	inds := rand.Int63()
	var i uint8
	for i = 0; i < n; i++ {
		res[i] = chars[(inds>>i*6)&63]
	}
	return string(res[0:n])
}
