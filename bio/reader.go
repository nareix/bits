
package bio

import (
	"io"
)

type readPeeker interface {
	Peek(int) ([]byte,error)
	io.Reader
}

type Reader struct {
	readPeeker
	intbuf []byte
}

func (self *Reader) ReadI8() (i int8, err error) {
	var b []byte
	if b, err = self.Peek(1); err != nil {
		return
	}
	i = int8(b[0])
	return
}

func (self *Reader) ReadU8() (i uint8, err error) {
	var b []byte
	if b, err = self.Peek(1); err != nil {
		return
	}
	i = uint8(b[0])
	return
}

func (self *Reader) ReadI16BE() (i int16, err error) {
	var b []byte
	if b, err = self.Peek(2); err != nil {
		return
	}
	i = int16(int8(b[0]))
	i <<= 8; i |= int16(b[1])
	return
}

func (self *Reader) ReadU16BE() (i uint16, err error) {
	var b []byte
	if b, err = self.Peek(2); err != nil {
		return
	}
	i = uint16(b[0])
	i <<= 8; i |= uint16(b[1])
	return
}

func (self *Reader) ReadI24BE() (i int32, err error) {
	var b []byte
	if b, err = self.Peek(3); err != nil {
		return
	}
	i = int32(int8(b[0]))
	i <<= 8; i |= int32(b[1])
	i <<= 8; i |= int32(b[2])
	return
}

func (self *Reader) ReadU24BE() (i uint32, err error) {
	var b []byte
	if b, err = self.Peek(3); err != nil {
		return
	}
	i = uint32(b[0])
	i <<= 8; i |= uint32(b[1])
	i <<= 8; i |= uint32(b[2])
	return
}

func (self *Reader) ReadI32BE() (i int32, err error) {
	var b []byte
	if b, err = self.Peek(4); err != nil {
		return
	}
	i = int32(int8(b[0]))
	i <<= 8; i |= int32(b[1])
	i <<= 8; i |= int32(b[2])
	i <<= 8; i |= int32(b[3])
	return
}

func (self *Reader) ReadU32BE() (i uint32, err error) {
	var b []byte
	if b, err = self.Peek(4); err != nil {
		return
	}
	i = uint32(b[0])
	i <<= 8; i |= uint32(b[1])
	i <<= 8; i |= uint32(b[2])
	i <<= 8; i |= uint32(b[3])
	return
}

func (self *Reader) ReadU64BE() (i uint64, err error) {
	var b []byte
	if b, err = self.Peek(8); err != nil {
		return
	}
	i = uint64(b[0])
	i <<= 8; i |= uint64(b[1])
	i <<= 8; i |= uint64(b[2])
	i <<= 8; i |= uint64(b[3])
	i <<= 8; i |= uint64(b[4])
	i <<= 8; i |= uint64(b[5])
	i <<= 8; i |= uint64(b[6])
	i <<= 8; i |= uint64(b[7])
	return
}

func (self *Reader) ReadI64BE() (i int64, err error) {
	var b []byte
	if b, err = self.Peek(8); err != nil {
		return
	}
	i = int64(int8(b[0]))
	i <<= 8; i |= int64(b[1])
	i <<= 8; i |= int64(b[2])
	i <<= 8; i |= int64(b[3])
	i <<= 8; i |= int64(b[4])
	i <<= 8; i |= int64(b[5])
	i <<= 8; i |= int64(b[6])
	i <<= 8; i |= int64(b[7])
	return
}

func NewReader(r readPeeker) *Reader {
	return &Reader{readPeeker: r}
}

