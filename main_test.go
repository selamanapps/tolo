package main

import (
	"testing"
)

func TestVersion(t *testing.T) {
	v := "1.0.0"
	expected := "tolo version 1.0.0"
	result := cmd.Version(v)
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
