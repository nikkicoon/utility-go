package pkg_test

import (
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGenerateEmail(t *testing.T) {
	s := pkg.GenerateEmail(true)
	assert.NotZero(t, s)
	a := strings.Split("abcdefghijklmnopqrstuvwxyz0123456789", "")
	b := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	ok := pkg.ContainsMultiple(s, a)
	assert.True(t, ok)
	s = pkg.GenerateEmail(false)
	ok = pkg.ContainsMultiple(s, b)
	assert.True(t, ok)
}

func TestGenerateDSVLine(t *testing.T) {
	s := pkg.GenerateDSVLine(4)
	assert.NotZero(t, s)
	assert.Len(t, strings.Split(s, " "), 4)
}

func TestGenerateDSVFile(t *testing.T) {
	s := pkg.GenerateDSVFile(10)
	assert.NotZero(t, s)
	assert.Len(t, strings.Split(s, "\n"), 10)
	assert.True(t, pkg.ContainsMultiple(strings.Split(strings.Split(s, "\n")[0], " ")[0], []string{"@", "."}))
}
