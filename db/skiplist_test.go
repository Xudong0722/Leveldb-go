package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/Xudong0722/Leveldb-go/utils"
)

func TestInsert(t *testing.T) {
	sl := NewSkipList(utils.StringComprator)
	sl.Insert(string("test"))
	res1 := sl.Contains(string("test"))
	assert.Equal(t, res1, true)
}