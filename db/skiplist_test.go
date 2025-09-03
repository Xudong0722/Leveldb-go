package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	sl := NewSkipList(StringKeyComprator)
	sl.Insert(string("test"))
	res1 := sl.Contains(string("test"))
	assert.Equal(t, res1, true)
	sl.Insert(string("111"))
	sl.Insert(string("222"))
	sl.Insert(string("333"))

	res2 := sl.Contains(string("111"))
	res3 := sl.Contains(string("3333"))
	assert.Equal(t, res2, true)
	assert.Equal(t, res3, false)
	res4 := sl.Contains(string("222"))
	assert.Equal(t, res4, true)
	res5 := sl.Delete(string("222"))
	assert.Equal(t, res5, true)
	res6 := sl.Contains(string("222"))
	assert.Equal(t, res6, false)
}

func Test_Leetcode_S1(t *testing.T) {
	sl := NewSkipList(IntComprator)
	sl.Insert(1)
	sl.Insert(2)
	sl.Insert(3)
	res1 := sl.Contains(0)
	assert.Equal(t, res1, false)

	sl.Insert(4)
	res2 := sl.Contains(1)
	assert.Equal(t, res2, true)
	res3 := sl.Delete(0)
	assert.Equal(t, res3, false)
	res4 := sl.Delete(1)
	assert.Equal(t, res4, true)
	res5 := sl.Contains(1)
	assert.Equal(t, res5, false)
}

func TestIterator(t *testing.T) {
	sl := NewSkipList(StringKeyComprator)
	sl.Insert(string("test"))
	res1 := sl.Contains(string("test"))
	assert.Equal(t, res1, true)
	sl.Insert(string("111"))
	sl.Insert(string("222"))
	sl.Insert(string("333"))

	res2 := sl.Contains(string("111"))
	res3 := sl.Contains(string("3333"))
	assert.Equal(t, res2, true)
	assert.Equal(t, res3, false)
	res4 := sl.Contains(string("222"))
	assert.Equal(t, res4, true)
	res5 := sl.Delete(string("222"))
	assert.Equal(t, res5, true)
	res6 := sl.Contains(string("222"))
	assert.Equal(t, res6, false)

	iter := NewSkipListIterator(sl)
	iter.SeekToFirst()

	assert.Equal(t, true, iter.Valid())
	//t.Log(iter.Key())
	for ; iter.Valid(); iter.Next() {
		t.Log(iter.Key())
	}

	sl.Insert(string("888"))
	sl.Insert(string("888"))
	sl.Insert(string("666"))
	sl.Insert(string("999"))
	sl.Insert(string("81818"))

	iter.SeekToFirst()

	for ; iter.Valid(); iter.Next() {
		t.Log(iter.Key())
	}


	iter.SeekToLast()

	for ; iter.Valid(); iter.Prev() {
		t.Log(iter.Key())
	}
}
