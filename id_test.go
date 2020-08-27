package h2sc

import (
	"testing"
)

func TestReset(t *testing.T) {
	sc := New()
	if err := sc.Identify("0000"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
}

func TestIdentify(t *testing.T) {
	sc := New()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.Identify("0000")
}