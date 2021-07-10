package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	var expected Phrase

	expected.Text = "Hello Cloud Run!"
	result := HelloWorld()

	if expected.Text != result.Text {
		t.Errorf("Phrase was incorrect, got: %s, want: %s.", result.Text, expected.Text)
	}
}
