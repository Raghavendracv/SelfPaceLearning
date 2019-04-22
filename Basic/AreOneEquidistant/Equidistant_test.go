package main

import "testing"

// TestAreEquistantForValidInputEndsWithOne validates IsEquistant for valid string and expects true result.
func TestAreEquistantForValidInputEndsWithOne(t *testing.T) {
	actual := AreEquistant("001001001")
	expected := true

	if actual != expected {
		t.Errorf("AreEquistant result was incorrect, got: %t, want: %t.", actual, expected)
	}
}

// TestAreEquistantForInvalidInput validates IsEquistant for invalid string and expects false result.
func TestAreEquistantForInvalidInput(t *testing.T) {
	actual := AreEquistant("01010101010000010001010101010101010010")
	expected := false

	if actual != expected {
		t.Errorf("AreEquistant result was incorrect, got: %t, want: %t.", actual, expected)
	}
}

// TestAreEquistantForValidInputEndsWithZero validates IsEquistant for valid string and expects true result.
func TestAreEquistantForValidInputEndsWithZero(t *testing.T) {
	actual := AreEquistant("01010101010")
	expected := true

	if actual != expected {
		t.Errorf("AreEquistant result was incorrect, got: %t, want: %t.", actual, expected)
	}
}

// TestAreEquistantForValidInputStartsWithOne validates IsEquistant for valid string and expects true result.
func TestAreEquistantForValidInputStartsWithOne(t *testing.T) {
	actual := AreEquistant("101010101010")
	expected := true

	if actual != expected {
		t.Errorf("AreEquistant result was incorrect, got: %t, want: %t.", actual, expected)
	}
}
