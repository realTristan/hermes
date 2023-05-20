package main

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pierrec/lz4"
)

func main() {
	// Compress and uncompress an input string.
	s := strings.Repeat("hello world", 1000)
	r := strings.NewReader(s)

	// The pipe will uncompress the data from the writer.
	pr, pw := io.Pipe()
	zw := lz4.NewWriter(pw)
	zr := lz4.NewReader(pr)

	var st = time.Now()
	go func() {
		// Compress the input string.
		_, _ = io.Copy(zw, r)
		_ = zw.Close() // Make sure the writer is closed
		_ = pw.Close() // Terminate the pipe
	}()
	fmt.Println(time.Since(st))

	// covnert zr to string
	st = time.Now()
	var b []byte
	var err error
	if b, err = io.ReadAll(zr); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(st))
	fmt.Println(len(b))

	// Output:
	// hello world
}
