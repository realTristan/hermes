package main

import (
	"fmt"
	"strings"
	"time"

	gzip "github.com/realTristan/Hermes/compression/gzip"
	"github.com/realTristan/Hermes/compression/zlib"
	utils "github.com/realTristan/Hermes/utils"
)

func main() {
	var v string = strings.Repeat("a", 4000)
	TestGzip(v)
	TestZlib(v)
}

// Test the zlib compression and decompression functions.
func TestZlib(v string) {
	fmt.Println("zlib")
	var (
		b   []byte
		err error
		st  time.Time = time.Now()
	)
	if b, err = zlib.Compress(v); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(st))
	st = time.Now()
	if v, err = zlib.Decompress(b); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(st))
	fmt.Println(utils.Size(v))
	fmt.Println(utils.Size(b))
}

// Test the gzip compression and decompression functions.
func TestGzip(v string) {
	fmt.Println("gzip")
	var (
		b   []byte
		err error
		st  time.Time = time.Now()
	)
	if b, err = gzip.Compress(v); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(st))
	st = time.Now()
	if v, err = gzip.Decompress(b); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(st))
	fmt.Println(utils.Size(v))
	fmt.Println(utils.Size(b))
}
