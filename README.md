# Golang bitstream reader/writer

```go
  r := &bits.Reader{R: stream}
  u32, err = r.ReadBits(4)
  u64, err = r.ReadBits64(4)
  p := make([]byte, 4)
  n, err = r.Read(p)
  
  w := &bits.Writer{W: stream}
  err = w.WriteBits(0xf, 4)
  err = w.WriteBits64(0xf, 4)
  p := []byte{1,2,3}
  n, err = w.Write(p)
  err = w.FlushBits()
  
```
