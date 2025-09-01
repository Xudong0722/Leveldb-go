package db

import "encoding/binary"

// key，value的内部编码
type LookupKey struct {
	klength uint32
	key     []byte
	// tag = seq no(7 bytes) + type(1 bytes)
	tag     uint64
	vlength uint32
	value   []byte
}

func NewLookupKeyWithK(key []byte, sequence_num uint64) *LookupKey {
	return &LookupKey{
		klength: uint32(len(key)),
		key:     key,
		tag:     encodeTag(sequence_num, kValue), //TODO
	}
}

func NewLookupKeyWithKV(key []byte, value []byte) *LookupKey {
	return &LookupKey{
		klength: uint32(len(key)),
		key:     key,
		vlength: uint32(len(value)),
		value:   value,
	}
}

func (lookup_key *LookupKey) ToMemKey() []byte {
	buf := make([]byte, binary.MaxVarintLen64*4+len(lookup_key.key)+len(lookup_key.value))
	index := 0
	index += binary.PutUvarint(buf[index:], uint64(lookup_key.klength))
	copy(buf[index:], lookup_key.key)
	index += len(lookup_key.key)
	index += binary.PutUvarint(buf[index:], lookup_key.tag)
	index += binary.PutUvarint(buf[index:], uint64(lookup_key.vlength))
	copy(buf[index:], lookup_key.value)
	index += len(lookup_key.value)
	return buf[:index]
}

func MemKeyToLookupKey(buf []byte) *LookupKey {
	klen, index := binary.Uvarint(buf)
	key := buf[index : index+int(klen)]
	index += int(klen)
	tag, len := binary.Uvarint(buf[index:])
	index += len
	vlen, len := binary.Uvarint(buf[index:])
	index += len

	return &LookupKey{
		klength: uint32(klen),
		key:     key,
		tag:     tag,
		vlength: uint32(vlen),
		value:   buf[index : index+int(vlen)],
	}

}

func encodeTag(seq uint64, typ ValueType) uint64 {
	return (seq << 8) | uint64(typ)
}

func decodeTag(tag uint64) (uint64, ValueType) {
	seq := tag >> 8
	typ := ValueType(tag & 0xff)
	return seq, typ
}
