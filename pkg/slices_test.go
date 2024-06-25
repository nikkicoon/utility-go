package pkg_test

import (
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDuplicateElements(t *testing.T) {
	var in []int
	in = append(in, 1, 2, 3, 1, 5, 6, 1)
	out := pkg.DuplicateElements(in)
	assert.Len(t, out, 3)
	assert.Equal(t, 1, out[0])
}
