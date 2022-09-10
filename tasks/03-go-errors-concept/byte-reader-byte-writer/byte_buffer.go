package bytebuffer

import (
	"fmt"
	"io"
)

const bufferMaxSize = 1024

type MaxSizeExceededError struct {
	desiredLen int
}

func (e *MaxSizeExceededError) Error() string {
	return fmt.Sprintf("buffer max size exceeded: %d > %d", e.desiredLen, bufferMaxSize)
}

type EndOfBufferError struct{}

func (b *EndOfBufferError) Error() string {
	return "end of buffer"
}

type ByteBuffer struct {
	buffer []byte
	offset int
}

var _ io.ByteWriter = new(ByteBuffer)
var _ io.ByteReader = new(ByteBuffer)

// Необходимо сделать так, чтобы тип *ByteBuffer реализовывал интерфейсы io.ByteWriter и io.ByteReader.
//
// Метод WriteByte должен возвращать ошибку *MaxSizeExceededError при попытке записи в буфер,
// если в нём уже больше bufferMaxSize байт.
//
// Метод ReadByte должен возвращать ошибку *EndOfBufferError при попытке чтения из буфера,
// если ранее буфер уже был вычитан полностью.

func (b *ByteBuffer) WriteByte(c byte) error {
	if len(b.buffer) == bufferMaxSize {
		return new(MaxSizeExceededError)
	}

	b.buffer = append(b.buffer, c)
	return nil
}

func (b *ByteBuffer) ReadByte() (byte, error) {
	if b.offset >= len(b.buffer) {
		return 0, new(EndOfBufferError)
	}

	c := b.buffer[b.offset]
	b.offset++

	return c, nil
}
