package conv

import (
	"testing"
)

func TestConv(t *testing.T) {
	num, err := Conv("123")
	if err != nil {
		t.Fatal(err)
	}
	if num != 123 {
		t.Fatalf("Expected 123, got %d", num)
	}
}

func TestFailConv(t *testing.T) {
	_, err := Conv("abc")
	if err == nil {
		t.Fatal(err)
	}
}
