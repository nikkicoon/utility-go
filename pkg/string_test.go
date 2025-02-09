package pkg_test

import (
	"github.com/nikkicoon/utility-go/pkg"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimSubstr(t *testing.T) {
	assert.Equal(t, "mushroom mushroom", pkg.TrimSubstr("badger badger badger mushroom mushroom", "badger "))
}

func TestFilepathParts(t *testing.T) {
	assert.Equal(t, pkg.FilepathParts("/home/badger/mushroom/sna.ke"), []string{"home", "badger", "mushroom", "sna.ke"})
}
