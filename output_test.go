package h2sc

import (
	"testing"
	"time"
)

func TestForceMuteAll(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	defer sc.ForceMute(PARAMETER_OUTPUT_MUTE_UNMUTE)
	sc.ForceMute(PARAMETER_OUTPUT_MUTE_ALL)
	t.Log("Muted All...")
	time.Sleep(time.Second * 3)
	t.Log("Unmute...")
}

func TestForceMuteVideo(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	defer sc.ForceMute(PARAMETER_OUTPUT_MUTE_VIDEO)
	sc.ForceMute(PARAMETER_OUTPUT_MUTE_ALL)
	t.Log("Muted Video...")
	time.Sleep(time.Second * 3)
	t.Log("Unmute...")
}

func TestForceMuteAudio(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	defer sc.ForceMute(PARAMETER_OUTPUT_MUTE_UNMUTE)
	sc.ForceMute(PARAMETER_OUTPUT_MUTE_AUDIO)
	t.Log("Muted Audio...")
	time.Sleep(time.Second * 3)
	t.Log("Unmute...")
}

func TestFreezeFrame(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	defer sc.FreezeFrame(false)
	sc.FreezeFrame(true)
	t.Log("Freezed current frame...")
	time.Sleep(time.Second * 3)
	t.Log("Unfreeze...")
}

func TestForceAspect(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	sc.ForceAspectRatio(PARAMETER_OUTPUT_ASPECT_FULL)
	t.Log("Set FULL Aspect...")
}

func TestOverrideOutputFormat1080p5994(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	if err := sc.OverrideOutputFormat(PARAMETER_OUTPUT_FORMAT_1080p5994); err != nil {
		t.Fatalf("Failed to override output format : %v", err)
	}
}

func TestOverrideOutputFormat720p5994(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	if err := sc.OverrideOutputFormat(PARAMETER_OUTPUT_FORMAT_720p5994); err != nil {
		t.Fatalf("Failed to override output format : %v", err)
	}
}

func TestRotate(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	sc.RotateAndMirror(PARAMETER_OUTPUT_ROTATE_180DEGREES)
	defer sc.RotateAndMirror(PARAMETER_OUTPUT_ROTATE_NORMAL)
	t.Log("Rotating 180degrees")
	time.Sleep(time.Second * 3)
}

func TestMirrorHorizonal(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	sc.RotateAndMirror(PARAMETER_OUTPUT_ROTATE_NORMAL)
	defer sc.RotateAndMirror(PARAMETER_OUTPUT_MIRROR_HORIZON)
	t.Log("Mirroring horizon")
	time.Sleep(time.Second * 3)
}

func TestMirrorVertical(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	sc.RotateAndMirror(PARAMETER_OUTPUT_ROTATE_NORMAL)
	defer sc.RotateAndMirror(PARAMETER_OUTPUT_MIRROR_HORIZON)
	t.Log("Mirroring horizon")
	time.Sleep(time.Second * 3)
}

func TestOutputTestPattern(t *testing.T) {
	sc, err := New()
	if err != nil {
		t.Fatalf("Failed to initialize : %v", err)
	}
	defer sc.Disconnect()
	if err := sc.Identify("0001"); err != nil {
		t.Fatalf("Failed to identify : %v", err)
	}
	defer sc.ResetID()
	loops := 3
	interval := time.Second * 3
	for i := 0; i < loops; i++ {
		sc.TestPattern(true)
		time.Sleep(interval)
		sc.TestPattern(false)
		time.Sleep(interval)
	}
}
