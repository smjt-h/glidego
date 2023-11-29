package main

import (
	"testing"
)

func TestGetFileSizeInfo(t *testing.T) {
	expected := "That file is 83 MB.\n"
	if getFileSizeInfo() != expected {
		t.Errorf("Expected: %s, but got: %s", expected, getFileSizeInfo())
	}
}

func TestGetOrdinalInfo(t *testing.T) {
	expected := "You're my 193rd best friend.\n"
	if getOrdinalInfo() != expected {
		t.Errorf("Expected: %s, but got: %s", expected, getOrdinalInfo())
	}
}

func TestGetCommaSeparatedValue(t *testing.T) {
	expected := "You owe $6,582,491.\n"
	if getCommaSeparatedValue() != expected {
		t.Errorf("Expected: %s, but got: %s", expected, getCommaSeparatedValue())
	}
}
