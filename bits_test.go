package bits

import (
	"testing"
	"bytes"
)

func TestBits(t *testing.T) {
	rdata := []byte{0xf3,0xb0}
	rbuf := bytes.NewReader(rdata[:])
	r := &Reader{R: rbuf}
	var u32 uint
	if u32, _ = r.ReadBits(4); u32 != 0xf {
		t.FailNow()
	}
	if u32, _ = r.ReadBits(4); u32 != 0x3 {
		t.FailNow()
	}
	if u32, _ = r.ReadBits(2); u32 != 0x2 {
		t.FailNow()
	}
	if u32, _ = r.ReadBits(2); u32 != 0x3 {
		t.FailNow()
	}

	wbuf := &bytes.Buffer{}
	w := &Writer{W: wbuf}
	w.WriteBits(0xf, 4)
	w.WriteBits(0x3, 4)
	w.WriteBits(0x2, 2)
	w.WriteBits(0x3, 2)
	w.FlushBits()
	wdata := wbuf.Bytes()
	if wdata[0] != 0xf3 || wdata[1] != 0xb0 {
		t.FailNow()
	}
}

