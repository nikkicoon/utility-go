package pkg_test

import (
	"fmt"
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSortedInsert(t *testing.T) {
	//logger := pkg.InitializeLogger("debug", "consolewriter", "", true, 1, pkg.SyslogConn{})
	var c []string
	c = pkg.SortedInsert(c, "one", nil)
	assert.NotEmpty(t, c)
	c = pkg.SortedInsert(c, "two", nil)
	r := []string{"one", "two"}
	assert.Equal(t, c, r)
}

func TestDuplicateElements(t *testing.T) {
	var in []int
	in = append(in, 1, 2, 3, 1, 5, 6, 1)
	out := pkg.DuplicateElements(in)
	assert.Len(t, out, 3)
	assert.Equal(t, 1, out[0])
}

func TestPrependInsertReversed(t *testing.T) {
	x := []int{9, 4, 1, 5}
	res := pkg.PrependInsertReversed(x, 8, 7, 6)
	fmt.Println(res)
	assert.Len(t, res, 7)
}

func TestPrependInsertSliced(t *testing.T) {
	x := []int{9, 4, 1, 5}
	res := pkg.PrependInsertSliced(x, 8, 7, 6)
	fmt.Println(res)
	assert.Len(t, res, 7)
}

func TestPrependInsertSliced2(t *testing.T) {
	x := []int{9, 4, 1, 5}
	res := pkg.PrependInsertSliced(x, 1)
	fmt.Println(res)
	assert.Len(t, res, 5)
}

func TestPrependInsertSliced3(t *testing.T) {
	x := []int{9, 4, 1, 5}
	res := pkg.PrependInsertSliced(x)
	fmt.Println(res)
	assert.Len(t, res, 4)
}
