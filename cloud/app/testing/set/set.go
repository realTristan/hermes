package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	set()
	get()
}

func set() {
	// create a map[string]interface{} to encode
	var data map[string]interface{} = map[string]interface{}{
		"test": "test",
	}
	// marshal the map[string]interface{} to json
	var value string
	if data, err := json.Marshal(data); err != nil {
		panic(err)
	} else {
		// base64 encode the json
		value = base64.StdEncoding.EncodeToString(data)
	}

	// send the request
	resp, err := http.Post("http://localhost:52412/set?key=test&ft=true&value="+value, "application/json", nil)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func get() {
	// Get the value
	resp, err := http.Get("http://localhost:52412/get?key=test")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
