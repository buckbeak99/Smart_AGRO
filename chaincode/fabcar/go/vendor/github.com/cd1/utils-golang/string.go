package utils

import (
	"bytes"
	"math/rand"
)

// StringLength defines the length of the string generated by RandomString.
const StringLength = 10

// AllRunes defines the possible characters which will be used by
// RandomString to generate a string.
const AllRunes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string with a specific size, defined by
// StringLength. The possible characters are defined by AllCharacters
// (A-Z, a-z, 0-9). This function uses the default random source of "math/rand".
func RandomString() string {
	var buffer bytes.Buffer

	for i := 0; i < StringLength; i++ {
		r := AllRunes[rand.Intn(len(AllRunes))]
		buffer.WriteRune(rune(r))
	}

	return buffer.String()
}
