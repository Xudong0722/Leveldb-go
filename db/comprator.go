package db

import (
	"bytes"
	"encoding/binary"
	"strings"

	"github.com/Xudong0722/Leveldb-go/utils"
)

type Comprator func(a, b interface{}) (int, error)

func MemTableKeyComprator(a, b interface{}) (int, error) {
	sa, ok := a.([]byte)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}
	sb, ok := b.([]byte)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}

	//取出user key来比较, 先比价key，如果key一致，比较版本号
	//key本身升序排列
	//相同的key，按seq降序排列，seq越大越新
	userkey_a, _, tag_a := ExtractCmpKey(sa)
	userkey_b, _, tag_b := ExtractCmpKey(sb)
	res := bytes.Compare(userkey_a, userkey_b)
	if res == 0 {
		if (tag_a >> 8) > (tag_b >> 8) {
			res = -1
		} else if (tag_a >> 8) < (tag_b >> 8) {
			res = 1
		}
	}
	return res, nil
}

func ByteArrayKeyComprator(a, b interface{}) (int, error) {
	sa, ok := a.([]byte)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}
	sb, ok := b.([]byte)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}

	return bytes.Compare(sa, sb), nil
}

func StringKeyComprator(a, b interface{}) (int, error) {
	sa, ok := a.(string)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}
	sb, ok := b.(string)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}

	return strings.Compare(sa, sb), nil
}

// return [user key, user key length, tag]
func ExtractCmpKey(buf []byte) ([]byte, uint32, uint64) {
	klen, index := binary.Uvarint(buf)
	key := buf[index : index+int(klen)]
	index += len(key)
	tag, _ := binary.Uvarint(buf[index:])
	return key, uint32(klen), tag
}

func IntComprator(a, b interface{}) (int, error) {
	sa, ok := a.(int)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}
	sb, ok := b.(int)
	if !ok {
		return 0, utils.ErrTypeMismatch
	}

	if sa == sb {
		return 0, nil
	}
	if sa < sb {
		return -1, nil
	}
	return 1, nil
}
