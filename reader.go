
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

func ReadUIntBE(r io.Reader, n int) (res uint, err error) {
	n /= 8
	var b []byte
	if b, err = ReadBytes(r, n); err != nil {
		return
	}
	for i := 0; i < n; i++ {
		res <<= 8
		res += uint(b[i])
	}
	return
}

func ReadIntBE(r io.Reader, n int) (res int, err error) {
	n /= 8
	var uval uint
	if uval, err = ReadUIntBE(r, n); err != nil {
		return
	}
	if uval&(1<<uint(n*8-1)) != 0 {
		res = -int((1<<uint(n*8))-uval)
	} else {
		res = int(uval)
	}
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

