# golang bit stream reader/writer

```go
  r := &bits.Reader{R: stream}
  u32, err = r.ReadBits(4)
  u64, err = r.ReadBits64(4)
  
  w := &bits.Writer{W: stream}
  err = w.WriteBits(0xf, 4)
  err = w.WriteBits64(0xf, 4)
  err = w.FlushBits()
```
