package bytezip

import (
	"bytes"
	"compress/gzip"
	"io"
)

type Level = int

const (
	NoCompression      Level = gzip.NoCompression
	BestSpeed          Level = gzip.BestSpeed
	BestCompression    Level = gzip.BestCompression
	DefaultCompression Level = gzip.DefaultCompression
	HuffmanOnly        Level = gzip.HuffmanOnly
)

// GZipBytes 压缩
func GZipBytes(data []byte, level Level) ([]byte, error) {
	var in bytes.Buffer
	writer, err := gzip.NewWriterLevel(&in, level)
	if err != nil {
		return nil, err
	}
	_, err = writer.Write(data)
	if err != nil {
		return nil, err
	}
	_ = writer.Close()
	return in.Bytes(), nil
}

// UGZipBytes 解压缩
func UGZipBytes(data []byte) ([]byte, error) {
	var out bytes.Buffer
	var in bytes.Buffer
	in.Write(data)
	reader, err := gzip.NewReader(&in)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	_, err = io.Copy(&out, reader)
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
