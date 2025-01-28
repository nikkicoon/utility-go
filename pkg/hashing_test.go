package pkg_test

import (
	"fmt"
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
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

func generateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	result := make([]byte, 0, length)
	for i := range result {
		result[i] = charset[random.Intn(len(charset))]
	}
	return string(result)
}

func gen5000000strings() []string {
	var res []string
	for i := 0; i < 5_000_000; i++ {
		res = append(res, generateRandomString(20))
	}
	return res
}

func BenchmarkCalculateHashSHA1_5000000(b *testing.B) {
	i := gen5000000strings()
	//var c = 0
	b.StartTimer()
	for _, v := range i {
		//c++
		_ = pkg.CalculateHashSHA1(v)
		//fmt.Println(b.Elapsed())
	}
	b.StopTimer()
	//fmt.Println(c)
	fmt.Println(b.Elapsed())
}
