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
	offset   uint32
	dst      WritableFile
	type_crc []uint32
}

func (w *Writer) AddRecord(data []byte) error {

}
