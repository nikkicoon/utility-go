package pkg

import "unsafe"

// BytesToStringZeroAlloc converts bytes to a string without memory allocation.
// NOTE: The given bytes MUST NOT be modified since they share the same backing array
// with the returned string.
func BytesToStringZeroAlloc(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// StringToBytesZeroAlloc converts a string to a byte slice without memory allocation.
// NOTE: The returned byte slice MUST NOT be modified since it shares the same backing array
// with the given string.
func StringToBytesZeroAlloc(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
