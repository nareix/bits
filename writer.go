
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

func WriteUIntBE(w io.Writer, val uint, n int) (err error) {
	n /= 8
	var b [8]byte
	for i := n-1; i >= 0; i-- {
		b[i] = byte(val)
		val >>= 8
	}
	return WriteBytes(w, b[:], n)
}

func WriteIntBE(w io.Writer, val int, n int) (err error) {
	n /= 8
	var uval uint
	if val < 0 {
		uval = uint((1<<uint(n*8))+val)
	} else {
		uval = uint(val)
	}
	return WriteUInt(w, uval, n)
}

func WriteString(w io.Writer, val string, n int) (err error) {
	return WriteBytes(w, []byte(val), n)
}

