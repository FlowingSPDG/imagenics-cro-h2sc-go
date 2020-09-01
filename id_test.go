package h2sc

import (
	"testing"
)

func TestReset(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.ResetID(); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
}

func TestIdentify(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("01"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
}
