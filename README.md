# Golang bitstream reader/writer

```go
var reader io.Reader
r := &bits.Reader{R: reader}
u32, err = r.ReadBits(4)
u64, err = r.ReadBits64(4)
p := make([]byte, 4)
n, err = r.Read(p)
  
var writer io.Writer
w := &bits.Writer{W: writer}
err = w.WriteBits(0xf, 4)
err = w.WriteBits64(0xf, 4)
n, err = w.Write([]byte{0x34,0x56})
err = w.FlushBits()
```
