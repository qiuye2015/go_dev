package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"math"
)

func toSha1(data string) string {
	h := sha1.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// characters used for conversion
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Base62Encode converts number to base62
func Base62Encode(number int64) string {
	if number == 0 {
		return string(alphabet[0])
	}
	chars := make([]byte, 0)
	length := (int64)(len(alphabet))
	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

// Base62Decode converts base62 token to int
func Base62Decode(token string) int64 {
	var number int64 = 0
	idx := 0.0
	chars := []byte(alphabet)
	charsLen := float64(len(chars))
	tokenLen := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLen - (idx + 1)
		index := bytes.IndexByte(chars, c)
		number += int64(index) * int64(math.Pow(charsLen, power))
		idx++
	}
	return number
}
