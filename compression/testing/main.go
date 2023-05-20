package main

import (
	"encoding/json"
	"fmt"
	"time"

	gzip "github.com/realTristan/Hermes/compression/gzip"
	"github.com/realTristan/Hermes/compression/zlib"
	utils "github.com/realTristan/Hermes/utils"
)

func main() {
	/*
		var v string = strings.Repeat("computer", 100)
		TestGzip(v)
		TestZlib(v)
	*/
	/*
		// read the file from ../../data/data_hash.json
		var (
			d   map[string]map[string]interface{}
			err error
		)
		if d, err = utils.ReadMapJson("../../data/data_hash.json"); err != nil {
			panic(err)
		}

		// marshal d
		var p, _ = json.Marshal(d)
		fmt.Println(utils.Size(p))
		TestMap(d)
	*/
	// test int array
	var v []int = []int{}
	for i := 0; i < 1000; i++ {
		v = append(v, i)
	}
	//fmt.Println(utils.Size(v))
	TestInt(v)
}

// Test the zlib compression and decompression functions.
func TestZlib(v string) {
	fmt.Println("zlib")
	var (
		b   []byte
		err error
		st  time.Time = time.Now()
	)
	if b, err = zlib.Compress([]byte(v)); err != nil {
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
	if b, err = gzip.Compress([]byte(v)); err != nil {
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

// Compress and decompress a map[string]map[string]interface{}
func TestMap(v map[string]map[string]interface{}) {
	var p, _ = json.Marshal(v)
	var d, _ = gzip.Compress(p)
	fmt.Println(utils.Size(d))
	var st = time.Now()
	var _, _ = gzip.Decompress(d)
	fmt.Println(time.Since(st))
}

// Compress and decompress a []int
func TestInt(v []int) {
	var p, _ = json.Marshal(v)
	var d, _ = gzip.Compress(p)
	fmt.Println(utils.Size(d))
	var st = time.Now()
	var _, _ = gzip.Decompress(d)
	fmt.Println(time.Since(st))
}
