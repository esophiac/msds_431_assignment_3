package main

import (
	"testing"
)

// verify that the createData function will read and return the csv data
func TestCreateData(t *testing.T) {
	in := "testData.csv"
	out := [][]string{{"cat", "dog", "bird"}, {"1", "2", "3"}}

	result := createData(in)

	if result == nil {
		t.Errorf("expected %s, got %s instead", out, result)
	}

}

// verify that writing to a file works correctly
// WARNING: this will create a new file in your working directory
func TestWriteData(t *testing.T) {
	in1 := [][]string{{"cat", "dog", "bird"}, {"1", "2", "3"}}
	in2 := "testResult.jsonl"
	out := "Your file is ready"

	result := writeData(in1, in2)

	if result != out {
		t.Errorf("expected %s, got %s instead", out, result)
	}

}

// verify that the buildHead function in main.go will return the expected slice
func TestBuildHead(t *testing.T) {
	in := []string{"cat", " dog ", "bird "}
	out := []string{"cat", "dog", "bird"}

	result := buildHead(in)

	if result == nil {
		t.Errorf("expected %s, got %s instead", out, result)
	}
}

// verify that the function to create a map from two slices works correctly
func TestCreateMap(t *testing.T) {
	in1 := []string{"cat", " dog ", "bird "}
	in2 := []string{"1", "2", "3"}
	out := make(map[string]int)

	out["cat"] = 1
	out["dog"] = 2
	out["bird"] = 3

	result := createMap(in1, in2)

	if result == nil {
		t.Errorf("expected %v, got %v instead", out, result)
	}
}
