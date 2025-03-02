package main

import (
	"fmt"

	"github.com/rockcutter/go_json_scannner/src/scanner"
)

func main() {
	jsonStr := []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`)
	s := scanner.NewJsonScanner(jsonStr)

	s = s.Get("hoge").Get("city")
	fmt.Println(s.IsInvalid())
}
