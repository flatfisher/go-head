package main

import (
	"fmt"
	"os"
	"testing"
)

func TestLineCounter(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Failed to open file.")
		os.Exit(2)
	}
	defer file.Close()
	if 20 != lineCounter(file) {
		t.Error()
	}
}

func TestPrintRows(t *testing.T) {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Failed to open file.")
		os.Exit(2)
	}
	defer file.Close()
	printRows(file)
}
