
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
		res = -int64((1<<uint64(n))-uval)
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

