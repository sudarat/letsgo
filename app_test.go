package main

import (
	"fmt"
	"testing"

	"./tour/http/controller/hello"
)

// Test Hello Api
func TestHello(t *testing.T) {
	fmt.Println("Test Hello() ")
	result := hello.Hello()
	// fmt.Println("Result: ", result)
	if result != "Hello" {
		t.Errorf("Fail: got %v want %v", result, "Hello")
	}
}
