package bits

import (
	"io"
)

func ReadBytes(r io.Reader, n int) (res []byte, err error) {
	res = make([]byte, n)
	if n, err = r.Read(res); err != nil {
		return
	}
	return
}

func GetUInt64BE(b []byte, n int) (res uint64) {
	n /= 8
	for i := 0; i < n; i++ {
		res <<= 8
		res += uint64(b[i])
	}
	return
}

func GetUIntBE(b []byte, n int) (res uint) {
	return uint(GetUInt64BE(b, n))
}

func GetInt64BE(b []byte, n int) (res int64) {
	uval := GetUInt64BE(b, n)
	if uval&(1<<uint64(n-1)) != 0 {
		res = -int64((1 << uint64(n)) - uval)
	} else {
		res = int64(uval)
	}
	return
}

func GetIntBE(b []byte, n int) (res uint) {
	return uint(GetInt64BE(b, n))
}

func ReadUInt64BE(r io.Reader, n int) (res uint64, err error) {
	var b []byte
	if b, err = ReadBytes(r, n/8); err != nil {
		return
	}
	res = GetUInt64BE(b, n)
	return
}

func ReadUIntBE(r io.Reader, n int) (res uint, err error) {
	var res64 uint64
	res64, err = ReadUInt64BE(r, n)
	res = uint(res64)
	return
}

func ReadInt64BE(r io.Reader, n int) (res int64, err error) {
	var b []byte
	if b, err = ReadBytes(r, n/8); err != nil {
		return
	}
	res = GetInt64BE(b, n)
	return
}

func ReadIntBE(r io.Reader, n int) (res int, err error) {
	var res64 int64
	res64, err = ReadInt64BE(r, n)
	res = int(res64)
	return
}

func ReadString(r io.Reader, n int) (res string, err error) {
	var b []byte
	if b, err = ReadBytes(r, n); err != nil {
		return
	}
	res = string(b)
	return
}

type IntReader struct {
	R   io.Reader
	buf []byte
}

func (self IntReader) ReadInt8() (i int8, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:1]); err != nil {
		return
	}
	i = int8(self.buf[0])
	return
}

func (self IntReader) ReadUInt8() (i uint8, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:1]); err != nil {
		return
	}
	i = uint8(self.buf[0])
	return
}

func (self IntReader) ReadInt16BE() (i int16, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:2]); err != nil {
		return
	}
	i = int16(int8(self.buf[0]))
	i <<= 8; i |= int16(self.buf[1])
	return
}

func (self IntReader) ReadUInt16BE() (i uint16, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:2]); err != nil {
		return
	}
	i = uint16(self.buf[0])
	i <<= 8; i |= uint16(self.buf[1])
	return
}

func (self IntReader) ReadInt24BE() (i int32, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:3]); err != nil {
		return
	}
	i = int32(int8(self.buf[0]))
	i <<= 8; i |= int32(self.buf[1])
	i <<= 8; i |= int32(self.buf[2])
	return
}

func (self IntReader) ReadUInt24BE() (i uint32, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:3]); err != nil {
		return
	}
	i = uint32(self.buf[0])
	i <<= 8; i |= uint32(self.buf[1])
	i <<= 8; i |= uint32(self.buf[2])
	return
}

func (self IntReader) ReadInt32BE() (i int32, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:4]); err != nil {
		return
	}
	i = int32(int8(self.buf[0]))
	i <<= 8; i |= int32(self.buf[1])
	i <<= 8; i |= int32(self.buf[2])
	i <<= 8; i |= int32(self.buf[3])
	return
}

func (self IntReader) ReadUInt32BE() (i uint32, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:4]); err != nil {
		return
	}
	i = uint32(self.buf[0])
	i <<= 8; i |= uint32(self.buf[1])
	i <<= 8; i |= uint32(self.buf[2])
	i <<= 8; i |= uint32(self.buf[3])
	return
}

func (self IntReader) ReadUInt64BE() (i uint64, err error) {
	if _, err = io.ReadFull(self.R, self.buf[0:8]); err != nil {
		return
	}
	i = uint64(self.buf[0])
	i <<= 8; i |= uint64(self.buf[1])
	i <<= 8; i |= uint64(self.buf[2])
	i <<= 8; i |= uint64(self.buf[3])
	i <<= 8; i |= uint64(self.buf[4])
	i <<= 8; i |= uint64(self.buf[5])
	i <<= 8; i |= uint64(self.buf[6])
	i <<= 8; i |= uint64(self.buf[7])
	return
}


func (self IntReader) Read(p []byte) (i int, err error) {
	return self.R.Read(p)
}

var discardBuf = make([]byte, 4096)

func (self IntReader) Skip(n int) (err error) {
	for n > 0 {
		var r int
		s := n
		if s > len(discardBuf) {
			s = len(discardBuf)
		}
		if r, err = self.R.Read(discardBuf[:s]); err != nil {
			return
		}
		n -= r
	}
	return
}

func NewIntReader(r io.Reader) IntReader {
	return IntReader{R: r, buf: make([]byte, 8)}
}

