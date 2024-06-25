package pkg

import (
	"crypto/sha1"
	"encoding/hex"
	"math/big"
)

// CalculateHashHexBase16String returns the string representation of
// the sha1 value for a given string (hex, base16).
func CalculateHashHexBase16String(key string) string {
	// SHA1 hash
	hash := sha1.New()
	hash.Write([]byte(key))
	hashBytes := hash.Sum(nil)
	// length of hashBytes: 20
	// hexadecimal conversion
	hexSHA1 := hex.EncodeToString(hashBytes)
	// integer base16 conversion
	res, success := new(big.Int).SetString(hexSHA1, 16)
	if !success {
		panic("failed parsing big int from hex")
	}
	return res.String()
}

// CalculateHashBin returns the sha1 binary value for a given string `key`.
func CalculateHashBin(key string) []byte {
	// SHA1 hash, length of result: 20
	hash := sha1.New()
	hash.Write([]byte(key))
	return hash.Sum(nil)
	/*
		data := []byte(key)
		return sha1.Sum(data)

	*/
}
