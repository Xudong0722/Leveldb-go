package db

import "os"

type WritableFile interface {
	//追加写,优先写入本地缓存
	Append([]byte) error
	//关闭文件
	Close() error
	//写到缓存里，不一定会同步写入磁盘中
	Flush() error
	//同步刷新到磁盘
	Sync() error
}

const MaxBlockSize = 32768                      //32KB
const WritableFileBufferSize = MaxBlockSize * 2 //64KB buffer

type FileIO struct {
	buf []byte
	pos uint32 //[0, pos-1] contains data to be written to fd
	fd  *os.File
}

func NewFileIO(path string) (*FileIO, error) {
	fd, err := os.OpenFile(
		path,
		os.O_CREATE|os.O_APPEND|os.O_RDWR,
		1644,
	)

	if err != nil {
		return nil, err
	}
	return &FileIO{
		buf: make([]byte, WritableFileBufferSize),
		fd:  fd,
	}, nil
}

func (fio *FileIO) Append(data []byte) error {
	write_size := uint32(len(data))
	copy_size := min(WritableFileBufferSize-fio.pos, write_size)
	copy(fio.buf[fio.pos:], data[:copy_size])
	write_size -= copy_size

	if write_size == 0 {
		//all data write to buf
		return nil
	}

	//we have left data, so need to do at least one write
	if err := fio.flushBuffer(); err != nil {
		return err
	}

	//small size data can fit in buf, large data are written directly.
	if write_size < WritableFileBufferSize {
		copy(fio.buf, data[copy_size:])
		fio.pos += write_size
		return nil
	}
	return fio.writeUnbuffed(data[copy_size:], write_size)
}

func (fio *FileIO) Close() error {
	return fio.fd.Close()
}

func (fio *FileIO) Flush() error {
	return fio.flushBuffer()
}

func (fio *FileIO) Sync() error {
	if err := fio.flushBuffer(); err != nil {
		return err
	}
	//TODO, logic
	return fio.fd.Sync()
}

func (fio *FileIO) flushBuffer() error {
	if err := fio.writeUnbuffed(fio.buf, fio.pos); err != nil {
		return err
	}
	fio.pos = 0
	return nil
}

func (fio *FileIO) writeUnbuffed(data []byte, sz uint32) error {
	_, err := fio.fd.Write(data[:sz])
	return err
}
