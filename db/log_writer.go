package db

type RecordType byte

const (
	kZeroType RecordType = iota
	kFullType

	//Fragments
	kFirstType
	kMiddleType
	kLastType
)

type Writer struct {
	offset uint32
}
