package h2sc

import (
	"fmt"
	"net"
)

var (
	OUTPUT_MUTE_UNMUTE = []byte("00000")
	OUTPUT_MUTE_ALL    = []byte("+0001")
	OUTPUT_MUTE_VIDEO  = []byte("+0002")
	OUTPUT_MUTE_AUDIO  = []byte("+0003")

	OUTPUT_FREEZE_ENABLE  = []byte("+0001")
	OUTPUT_FREEZE_DISABLE = []byte("00000")

	OUTPUT_ASPECT_ASPECTKEEP    = []byte("00000")
	OUTPUT_ASPECT_FULL          = []byte("+0001")
	OUTPUT_ASPECT_HORIZONAL_FIT = []byte("+0002")
	OUTPUT_ASPECT_VERTICAL_FIT  = []byte("+0003")

	OUTPUT_FORMAT_1080i5994 = []byte("00000")
	OUTPUT_FORMAT_480i      = []byte("+0001")
	OUTPUT_FORMAT_720p5994  = []byte("+0004")
	OUTPUT_FORMAT_1080p5994 = []byte("+0005")
	OUTPUT_FORMAT_1080p60   = []byte("+0006")
	OUTPUT_FORMAT_1080p30   = []byte("+0008")
	// OUTPUT_FORMAT_1080i5994 = []byte("+0016") // Same as 000000
	// OUTPUT_FORMAT_480i = []byte("+0017") // same as +0001

	OUTPUT_SEAMLESS_BEHAVIOR_FREEZE_SEAMLESS    = []byte("00000")
	OUTPUT_SEAMLESS_BEHAVIOR_BLACK_CONNECT      = []byte("+0001")
	OUTPUT_SEAMLESS_BEHAVIOR_BLACK_FADE_CONNECT = []byte("+0002")
	OUTPUT_SEAMLESS_BEHAVIOR_FRAMELOCK          = []byte("0001")

	OUTPUT_ROTATE_NORMAL                                     = []byte("00000")
	OUTPUT_ROTATE_90DEGREES_COUNTER_CLOCKWISE                = []byte("+0001")
	OUTPUT_ROTATE_90DEGREES_CLOCKWISE                        = []byte("+0002")
	OUTPUT_ROTATE_180DEGREES                                 = []byte("+0003")
	OUTPUT_MIRROR_HORIZON                                    = []byte("+0004")
	OUTPUT_MIRROR_VERTICAL                                   = []byte("+0005")
	OUTPUT_ROTATE_90DEGREES_COUNTER_CLOCKWISE_MIRROR_VERICAL = []byte("+0006") // fuck.
	OUTPUT_ROTATE_90DEGREES_CLOCKWISE_MIRROR_VERICAL         = []byte("+0007")

	OUTPUT_POWERSAVE_1MIN  = []byte("00000") // 00000 or +0004
	OUTPUT_POWERSAVE_5MIN  = []byte("+0001") // +0001 or +0005
	OUTPUT_POWERSAVE_10MIN = []byte("+0002") // +0002 or +0006
	OUTPUT_POWERSAVE_NEVER = []byte("+0003")

	OUTPUT_TESTPATTERN_DISABLE = []byte("00000")
	OUTPUT_TESTPATTERN_ENABLE  = []byte("+0001")
)

// ForceMute Forcemute specified
func (h *H2SC) ForceMute(state []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_FORCEMUTE
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// FreezeFrame Freezes frame
func (h *H2SC) FreezeFrame(freeze bool) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	param := OUTPUT_FREEZE_DISABLE
	if freeze {
		param = OUTPUT_FREEZE_ENABLE
	}

	packet := newPacket()
	packet.command = COMMAND_FREEZE
	packet.idnum = h.id
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// ForceAspectRatio Override aspect-ratio
func (h *H2SC) ForceAspectRatio(state []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_ASPECT
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// OverrideOutputFormat Overrides output SDI signal format.
func (h *H2SC) OverrideOutputFormat(format []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_SDIFORMAT
	packet.idnum = h.id
	packet.param = format

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// SeamlessBehavior Change seamless switching behavior.
func (h *H2SC) SeamlessBehavior(state []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_SEAMLESS_AND_LOCK
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// RotateAndMirror Rotating and mirroring.
func (h *H2SC) RotateAndMirror(state []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_ROTATE_AND_MIRROR
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// PowerSave Set power-save behavior.
func (h *H2SC) PowerSave(state []byte) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_POWERSAVE
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// TestPattern Enable test pattern output.
func (h *H2SC) TestPattern(enabled bool) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	param := OUTPUT_TESTPATTERN_DISABLE
	if enabled {
		param = OUTPUT_TESTPATTERN_ENABLE
	}

	packet := newPacket()
	packet.command = COMMAND_TESTPATTERN
	packet.idnum = h.id
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// TODO...
// func (h *H2SC) OnScreenInformation(state []byte]) error {  }
// func (h *H2SC) GenLockHorizonalOffset(state []byte]) error {  }
// func (h *H2SC) GenLockVerticalOffset(state []byte]) error {  }
