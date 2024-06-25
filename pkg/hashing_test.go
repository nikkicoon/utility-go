package pkg_test

import (
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	res := pkg.CalculateHashBin("test")
	assert.Equal(t, []byte{0xa9, 0x4a, 0x8f, 0xe5, 0xcc, 0xb1, 0x9b, 0xa6, 0x1c, 0x4c, 0x8, 0x73, 0xd3, 0x91, 0xe9, 0x87, 0x98, 0x2f, 0xbb, 0xd3}, res)
}

func BenchmarkCalculateHash(b *testing.B) {
	for n := 0; n < 6000000; n++ {
		pkg.CalculateHashBin("test")
	}
	b.Elapsed()
}

func BenchmarkCalculateHashGoRoutines(b *testing.B) {
	for n := 0; n < 6000000; n++ {
		go pkg.CalculateHashBin("test")
	}
	b.Elapsed()
}
