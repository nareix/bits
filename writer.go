package bits

import (
	"io"
)

func WriteBytes(w io.Writer, b []byte, n int) (err error) {
	if len(b) < n {
		b = append(b, make([]byte, n-len(b))...)
	}
	_, err = w.Write(b[:n])
	return
}

func WriteUInt64BE(w io.Writer, val uint64, n int) (err error) {
	n /= 8
	var b [8]byte
	for i := n - 1; i >= 0; i-- {
		b[i] = byte(val)
		val >>= 8
	}
	return WriteBytes(w, b[:], n)
}

func WriteUIntBE(w io.Writer, val uint, n int) (err error) {
	return WriteUInt64BE(w, uint64(val), n)
}

func WriteInt64BE(w io.Writer, val int64, n int) (err error) {
	n /= 8
	var uval uint
	if val < 0 {
		uval = uint((1 << uint(n*8)) + val)
	} else {
		uval = uint(val)
	}
	return WriteUIntBE(w, uval, n)
}

func WriteIntBE(w io.Writer, val int, n int) (err error) {
	return WriteInt64BE(w, int64(val), n)
}

func WriteString(w io.Writer, val string, n int) (err error) {
	return WriteBytes(w, []byte(val), n)
}

func PutUInt64BE(b []byte, res uint64, n int) {
	n /= 8
	for i := 0; i < n; i++ {
		b[n-i-1] = byte(res)
		res >>= 8
	}
	return
}

func PutUIntBE(b []byte, res uint, n int) {
	PutUInt64BE(b, uint64(res), n)
}

type IntWriter struct {
	W   io.Writer
	buf []byte
}

func (self IntWriter) WriteInt8(v int8) (err error) {
	self.buf[0] = byte(v)
	_, err = self.W.Write(self.buf[:1])
	return
}

func (self IntWriter) WriteUInt8(v uint8) (err error) {
	self.buf[0] = byte(v)
	_, err = self.W.Write(self.buf[:1])
	return
}

func (self IntWriter) WriteInt16BE(v int16) (err error) {
	self.buf[0] = byte(v>>8)
	self.buf[1] = byte(v)
	_, err = self.W.Write(self.buf[:2])
	return
}

func (self IntWriter) WriteUInt16BE(v uint16) (err error) {
	self.buf[0] = byte(v>>8)
	self.buf[1] = byte(v)
	_, err = self.W.Write(self.buf[:2])
	return
}

func (self IntWriter) WriteInt24BE(v int32) (err error) {
	self.buf[0] = byte(v>>16)
	self.buf[1] = byte(v>>8)
	self.buf[2] = byte(v)
	_, err = self.W.Write(self.buf[:3])
	return
}

func (self IntWriter) WriteUInt24BE(v uint32) (err error) {
	self.buf[0] = byte(v>>16)
	self.buf[1] = byte(v>>8)
	self.buf[2] = byte(v)
	_, err = self.W.Write(self.buf[:3])
	return
}

func (self IntWriter) WriteInt32BE(v int32) (err error) {
	self.buf[0] = byte(v>>24)
	self.buf[1] = byte(v>>16)
	self.buf[2] = byte(v>>8)
	self.buf[3] = byte(v)
	_, err = self.W.Write(self.buf[:4])
	return
}

func (self IntWriter) WriteUInt32BE(v uint32) (err error) {
	self.buf[0] = byte(v>>24)
	self.buf[1] = byte(v>>16)
	self.buf[2] = byte(v>>8)
	self.buf[3] = byte(v)
	_, err = self.W.Write(self.buf[:4])
	return
}

func (self IntWriter) WriteUInt64BE(v uint64) (err error) {
	self.buf[0] = byte(v>>56)
	self.buf[1] = byte(v>>48)
	self.buf[2] = byte(v>>40)
	self.buf[3] = byte(v>>32)
	self.buf[4] = byte(v>>24)
	self.buf[5] = byte(v>>16)
	self.buf[6] = byte(v>>8)
	self.buf[7] = byte(v)
	_, err = self.W.Write(self.buf[:8])
	return
}

func (self IntWriter) Write(p []byte) (n int, err error) {
	return self.W.Write(p)
}

func NewIntWriter(w io.Writer) IntWriter {
	return IntWriter{W: w, buf: make([]byte, 8)}
}

