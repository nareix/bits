
package bio

import (
	"io"
)

type Writer struct {
	io.Writer
	intbuf []byte
}

func (self *Writer) WriteI8(v int8) (err error) {
	self.intbuf[0] = byte(v)
	_, err = self.Write(self.intbuf[:1])
	return
}

func (self *Writer) WriteU8(v uint8) (err error) {
	self.intbuf[0] = byte(v)
	_, err = self.Write(self.intbuf[:1])
	return
}

func (self *Writer) WriteI16BE(v int16) (err error) {
	self.intbuf[0] = byte(v>>8)
	self.intbuf[1] = byte(v)
	_, err = self.Write(self.intbuf[:2])
	return
}

func (self *Writer) WriteU16BE(v uint16) (err error) {
	self.intbuf[0] = byte(v>>8)
	self.intbuf[1] = byte(v)
	_, err = self.Write(self.intbuf[:2])
	return
}

func (self *Writer) WriteI24BE(v int32) (err error) {
	self.intbuf[0] = byte(v>>16)
	self.intbuf[1] = byte(v>>8)
	self.intbuf[2] = byte(v)
	_, err = self.Write(self.intbuf[:3])
	return
}

func (self *Writer) WriteU24BE(v uint32) (err error) {
	self.intbuf[0] = byte(v>>16)
	self.intbuf[1] = byte(v>>8)
	self.intbuf[2] = byte(v)
	_, err = self.Write(self.intbuf[:3])
	return
}

func (self *Writer) WriteI32BE(v int32) (err error) {
	self.intbuf[0] = byte(v>>24)
	self.intbuf[1] = byte(v>>16)
	self.intbuf[2] = byte(v>>8)
	self.intbuf[3] = byte(v)
	_, err = self.Write(self.intbuf[:4])
	return
}

func (self *Writer) WriteU32BE(v uint32) (err error) {
	self.intbuf[0] = byte(v>>24)
	self.intbuf[1] = byte(v>>16)
	self.intbuf[2] = byte(v>>8)
	self.intbuf[3] = byte(v)
	_, err = self.Write(self.intbuf[:4])
	return
}

func (self *Writer) WriteU64BE(v uint64) (err error) {
	self.intbuf[0] = byte(v>>56)
	self.intbuf[1] = byte(v>>48)
	self.intbuf[2] = byte(v>>40)
	self.intbuf[3] = byte(v>>32)
	self.intbuf[4] = byte(v>>24)
	self.intbuf[5] = byte(v>>16)
	self.intbuf[6] = byte(v>>8)
	self.intbuf[7] = byte(v)
	_, err = self.Write(self.intbuf[:8])
	return
}

func (self *Writer) WriteI64BE(v int64) (err error) {
	self.intbuf[0] = byte(v>>56)
	self.intbuf[1] = byte(v>>48)
	self.intbuf[2] = byte(v>>40)
	self.intbuf[3] = byte(v>>32)
	self.intbuf[4] = byte(v>>24)
	self.intbuf[5] = byte(v>>16)
	self.intbuf[6] = byte(v>>8)
	self.intbuf[7] = byte(v)
	_, err = self.Write(self.intbuf[:8])
	return
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{Writer: w, intbuf: make([]byte, 8)}
}

