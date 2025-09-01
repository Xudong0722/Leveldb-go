package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemTable_Get(t *testing.T) {
	tb := NewMemTable()
	tb.Put([]byte("test-2"), []byte("value-2"))
	tb.Put([]byte("test-5"), []byte("value-5"))
	tb.Put([]byte("test-8"), []byte("value-8"))
	tb.Put([]byte("test-7"), []byte("value-7"))
	tb.Put([]byte("test-6"), []byte("value-6"))
	tb.Put([]byte("test-3"), []byte("value-3"))
	tb.Put([]byte("test-4"), []byte("value-4"))
	tb.Put([]byte("test-0"), []byte("value-0"))

	val, err := tb.Get([]byte("test-2"))
	assert.Nil(t, err)
	assert.Equal(t, val, []byte("value-2"))

	val1, err := tb.Get([]byte("test-8"))
	assert.Nil(t, err)
	assert.Equal(t, val1, []byte("value-8"))

	val2, err := tb.Get([]byte("test-6"))
	assert.Nil(t, err)
	assert.Equal(t, val2, []byte("value-6"))

	val3, err := tb.Get([]byte("test-0"))
	assert.Nil(t, err)
	assert.Equal(t, val3, []byte("value-0"))

	t.Log(string(val3))
}
