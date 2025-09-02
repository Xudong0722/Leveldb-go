package db

import (
	"bytes"

	"github.com/Xudong0722/Leveldb-go/utils"
)

type ValueType byte

const (
	Unknown ValueType = iota
	kValue
	kDelete
)

type MemTable struct {
	//当前内存表使用的数据结构
	mem *SkipList

	//当前内存表的序列号
	id uint32

	//当前内存表占用的字节大小
	approximate_size uint32
}

type MemTableIterator struct {
	//迭代器所对应的内存表
	table *SkipList
}

func NewMemTable() *MemTable {
	return &MemTable{
		mem:              NewSkipList(MemTableKeyComprator),
		id:               0,
		approximate_size: 0,
	}
}

func (mt *MemTable) NewMemTableIterator() *MemTableIterator {
	return &MemTableIterator{
		table: mt.mem,
	}
}

func (mt *MemTable) Get(key []byte) ([]byte, error) {
	lookup_key := NewLookupKeyWithK(key, GetTempSeqNum())
	node, _ := mt.mem.GetGreaterOrEqual(lookup_key.ToMemKey())
	if node == nil {
		return nil, nil
	}

	lookup_key = MemKeyToLookupKey(node.key.([]byte))

	if !bytes.Equal(key, lookup_key.key) {
		return nil, utils.ErrKeyNotFound
	}

	_, typ := decodeTag(lookup_key.tag)
	if typ == kDelete {
		return nil, utils.ErrKeyNotFound
	}
	return lookup_key.value, nil
}

func (mt *MemTable) Put(key []byte, value []byte) {
	lookup_key := NewLookupKeyWithKV(key, value)
	lookup_key.tag = encodeTag(GetAndIncreaseSeqNum(), kValue)
	//fmt.Println(lookup_key)
	mt.mem.Insert(lookup_key.ToMemKey())
}

var initial_num uint64 = 0

// Just for test
func GetAndIncreaseSeqNum() uint64 {
	initial_num += 1
	return initial_num
}
func GetTempSeqNum() uint64 {
	return initial_num
}
