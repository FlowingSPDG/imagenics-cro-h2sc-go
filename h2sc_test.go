package h2sc

import (
	"testing"
	"time"
)

func TestOutputTestPattern(t *testing.T) {
	sc := New()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.Identify("0000")
	loops := 3
	interval := time.Second * 3
	for i := 0; i < loops; i++ {
		sc.TestPattern(true)
		time.Sleep(interval)
		sc.TestPattern(false)
		time.Sleep(interval)
	}
	if err := sc.Identify("0000"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
}

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

func TestOverrideOutputFormat1080p5994(t *testing.T) {
	sc := New()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.Identify("0000")
	if err := sc.OverrideOutputFormat(OUTPUT_FORMAT_1080p5994); err != nil {
		t.Fatalf("Failed to override output format : %v", err)
	}
}

func TestOverrideOutputFormat720p5994(t *testing.T) {
	sc := New()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.Identify("0000")
	if err := sc.OverrideOutputFormat(OUTPUT_FORMAT_720p5994); err != nil {
		t.Fatalf("Failed to override output format : %v", err)
	}
}

func TestFreezeCurrentFrame(t *testing.T) {
	sc := New()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.Identify("0000")
	loops := 3
	interval := time.Second * 3
	for i := 0; i < loops; i++ {
		sc.FreezeFrame(true)
		time.Sleep(interval)
		sc.FreezeFrame(false)
		time.Sleep(interval)
	}
}
