package benchmarks

import (
	"bufio"
	"bytes"
	"compress/flate"
	"io"
)

func Deflate(data []byte) []byte {
	var buffer bytes.Buffer
	writer, err := flate.NewWriter(&buffer, flate.DefaultCompression)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(writer, bytes.NewReader(data)); err != nil {
		panic(err)
	}
	if err := writer.Close(); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

func Inflate(data []byte) []byte {
	var buffer bytes.Buffer
	reader := flate.NewReader(bytes.NewReader(data))
	writer := bufio.NewWriter(&buffer)
	if _, err := io.Copy(writer, reader); err != nil {
		panic(err)
	}
	return buffer.Bytes()
}
