package pkg_test

import (
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItoB(t *testing.T) {
	assert.True(t, pkg.ItoB(20))
	assert.False(t, pkg.ItoB(0))
	assert.False(t, pkg.ItoB(-1))
}
