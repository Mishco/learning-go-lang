package main

import (
	"bytes"
	"io"
	"os"
	"testing"
	
)

func TestMain(t *testing.T) {
	input := 5
	expected := "Decimal: 5\nBinary: 101\nHexadecimal: 0x5\n"

	result := captureOutput(prints_values, input)

	if result != expected {
		t.Errorf("Test failed, expected: '%s', got: '%s'", expected, result)
	}
}

func captureOutput(f func(int), input int) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f(input)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
