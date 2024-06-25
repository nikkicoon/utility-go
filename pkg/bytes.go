package pkg

import (
	"bytes"
	"slices"
)

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
