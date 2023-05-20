package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

// Compress is a function that compresses a string using gzip compression.
//
// Parameters:
//   - v: A string representing the value to compress.
//
// Returns:
//   - A byte slice representing the compressed value.
//   - An error if there was an error compressing the value.
//
// Example usage:
//
//	compressed, err := Compress("value") // compressed == []byte{...}, err == nil
func Compress(v string) ([]byte, error) {
	var (
		b  *bytes.Buffer = new(bytes.Buffer)
		gz *gzip.Writer  = gzip.NewWriter(b)
	)
	if _, err := gz.Write([]byte(v)); err != nil {
		return nil, err
	}
	if err := gz.Close(); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Decompress is a function that decompresses a byte slice using gzip decompression.
//
// Parameters:
//   - v: A byte slice representing the compressed value to decompress.
//
// Returns:
//   - A string representing the decompressed value.
//   - An error if there was an error decompressing the value.
//
// Example usage:
//
//	decompressed, err := Decompress([]byte{...}) // decompressed == "value", err == nil
func Decompress(v []byte) (string, error) {
	var b *bytes.Buffer = bytes.NewBuffer(v)
	if gz, err := gzip.NewReader(b); err != nil {
		return "", err
	} else if s, err := io.ReadAll(gz); err != nil {
		return "", err
	} else {
		return string(s), nil
	}
}
