package pkg

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"math/big"
	"slices"
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

func BinarySearchBytes(arr [][]byte, target []byte, left int, right int) (int, bool) {
	if right == -1 {
		// catch errors
		return 0, false
	}
	if len(arr) == 0 || right == -1 {
		// no elements in slice
		return 0, false
	}
	if left > right {
		// not found
		return -1, false
	}
	mid := left + (right-left)/2
	if bytes.Equal(arr[mid], target) {
		// found, return
		return mid, true
	} else if bytes.Compare(arr[mid], target) < 0 {
		// a less than b
		return BinarySearchBytes(arr, target, mid+1, right)
	} else {
		return BinarySearchBytes(arr, target, left, mid-1)
	}
}

// SortedInsertByte inserts t (byte slice) into ts (slice of byte slices).
func SortedInsertByte(ts [][]byte, t []byte) [][]byte {
	// find slot
	i, ok := BinarySearchBytes(ts, t, 0, len(ts)-1)
	if !ok {
		// value not found in slice
		i = len(ts)
	}
	// if value is not found, assume index=0
	// fmt.Printf("inserting at position %d, length of ts %d\n", i, len(ts))
	return slices.Insert(ts, i, t)
}
