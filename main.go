package main

import (
	"fmt"

	"github.com/rockcutter/go_json_scannner/src/scanner"
)

func main() {
	nest := []byte(`{"name": "John", "age": 30, "address": {"city": "New York", "zip": 10001}}`)
	s := scanner.NewJsonScanner(nest)
	name := s.Get("address").ToString()
	fmt.Println(*name)
}
