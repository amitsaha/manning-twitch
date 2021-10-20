package main

import (
	"strings"
	"testing"
)

func TestGetInput(t *testing.T) {
	testInput := strings.NewReader("hello world")

	result, err := capitalize(testInput)

	if err != nil {
		t.Fatal(err)
	}
	expectedResult := "HELLO WORLD"

	if expectedResult != result {
		t.Fatalf("Expected: %s, Got: %s", expectedResult, result)
	}
}
