package main

import (
	"fmt"
	"testing"

	"./tour/http/controller/hello"
)

// Test is ...
func Test(t *testing.T) {
	result := hello.Hello()
	fmt.Println("Result : ", result)
	if result != "Hello" {
		t.Errorf("Fail: got %v want %v", result, "Hello")
	}
}
