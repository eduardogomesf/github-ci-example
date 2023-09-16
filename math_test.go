package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Error("Expected 5 + 5 to equal 10")
	}
}
