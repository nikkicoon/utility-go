package pkg_test

import (
	"bufio"
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCheckFile(t *testing.T) {
	assert.True(t, pkg.CheckFile("../counter_test"))
	assert.False(t, pkg.CheckFile("../END"))
}

func TestCheckExecutable(t *testing.T) {
	ok, err := pkg.CheckExecutable("/bin/sh", nil)
	assert.True(t, ok)
	assert.Nil(t, err)
}

func TestLineCounter(t *testing.T) {
	file, err := os.Open("../counter_test")
	assert.Nil(t, err)
	reader := bufio.NewReader(file)
	count, err := pkg.LineCounter(nil, reader, true)
	assert.Nil(t, err)
	assert.Equal(t, count, 5)
}

func TestSymlinkFiles(t *testing.T) {
	assert.Nil(t, pkg.SymlinkFiles(nil, "/dev/null", "tmp"))
	assert.NotNil(t, pkg.SymlinkFiles(nil, "marxisminmycomputerohno", "tmp"))
}
