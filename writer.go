
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
	for i := n-1; i >= 0; i-- {
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
		uval = uint((1<<uint(n*8))+val)
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

