package main

import (
	"fmt"
	"testing"

	"./tour/http/controller/callapi"
)

// TestLenMakeIPAddressList for test makeIPAddressList(ipList []string)
func TestLenMakeIPAddressList(t *testing.T) {
	fmt.Println("Test Length  makeIPAddressList(ipList []string)")
	input := []string{"A", "B", "C"}
	result := callapi.MakeIPAddressList(input)
	// fmt.Println("Result: ", result)

	if len(result) != len(input) {
		t.Errorf("Fail: got lenght %v want leght %v", len(result), len(input))
	}

}

// TestDataMakeIPAddressList for test data return from makeIPAddressList(ipList []string)
func TestDataMakeIPAddressList(t *testing.T) {
	fmt.Println("Test Data  makeIPAddressList(ipList []string)")
	input := []string{"A", "B", "C"}
	result := callapi.MakeIPAddressList(input)
	// fmt.Println("Result: ", result)

	for i := 0; i < len(result); i++ {
		r := result[i]
		// fmt.Println(r)
		if r.IPAddress != input[i] {
			t.Errorf("Fail: got %v want %v", r.IPAddress, input[i])
		}
	}

}
