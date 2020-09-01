package h2sc

import "time"

type PARAMETER []byte

var (
	PARAMETER_OUTPUT_MUTE_UNMUTE = PARAMETER("00000")
	PARAMETER_OUTPUT_MUTE_ALL    = PARAMETER("+0001")
	PARAMETER_OUTPUT_MUTE_VIDEO  = PARAMETER("+0002")
	PARAMETER_OUTPUT_MUTE_AUDIO  = PARAMETER("+0003")

	PARAMETER_OUTPUT_FREEZE_ENABLE  = PARAMETER("+0001")
	PARAMETER_OUTPUT_FREEZE_DISABLE = PARAMETER("00000")

	PARAMETER_OUTPUT_ASPECT_ASPECTKEEP    = PARAMETER("00000")
	PARAMETER_OUTPUT_ASPECT_FULL          = PARAMETER("+0001")
	PARAMETER_OUTPUT_ASPECT_HORIZONAL_FIT = PARAMETER("+0002")
	PARAMETER_OUTPUT_ASPECT_VERTICAL_FIT  = PARAMETER("+0003")

	PARAMETER_OUTPUT_FORMAT_1080i5994 = PARAMETER("00000")
	PARAMETER_OUTPUT_FORMAT_480i      = PARAMETER("+0001")
	PARAMETER_OUTPUT_FORMAT_720p5994  = PARAMETER("+0004")
	PARAMETER_OUTPUT_FORMAT_1080p5994 = PARAMETER("+0005")
	PARAMETER_OUTPUT_FORMAT_1080p60   = PARAMETER("+0006")
	PARAMETER_OUTPUT_FORMAT_1080p30   = PARAMETER("+0008")
	// PARAMETER_OUTPUT_FORMAT_1080i5994 = PARAMETER("+0016") // Same as 000000
	// PARAMETER_OUTPUT_FORMAT_480i = PARAMETER("+0017") // same as +0001

	PARAMETER_OUTPUT_SEAMLESS_BEHAVIOR_FREEZE_SEAMLESS    = PARAMETER("00000")
	PARAMETER_OUTPUT_SEAMLESS_BEHAVIOR_BLACK_CONNECT      = PARAMETER("+0001")
	PARAMETER_OUTPUT_SEAMLESS_BEHAVIOR_BLACK_FADE_CONNECT = PARAMETER("+0002")
	PARAMETER_OUTPUT_SEAMLESS_BEHAVIOR_FRAMELOCK          = PARAMETER("0001")

	PARAMETER_OUTPUT_ROTATE_NORMAL                                     = PARAMETER("00000")
	PARAMETER_OUTPUT_ROTATE_90DEGREES_COUNTER_CLOCKWISE                = PARAMETER("+0001")
	PARAMETER_OUTPUT_ROTATE_90DEGREES_CLOCKWISE                        = PARAMETER("+0002")
	PARAMETER_OUTPUT_ROTATE_180DEGREES                                 = PARAMETER("+0003")
	PARAMETER_OUTPUT_MIRROR_HORIZON                                    = PARAMETER("+0004")
	PARAMETER_OUTPUT_MIRROR_VERTICAL                                   = PARAMETER("+0005")
	PARAMETER_OUTPUT_ROTATE_90DEGREES_COUNTER_CLOCKWISE_MIRROR_VERICAL = PARAMETER("+0006") // fuck.
	PARAMETER_OUTPUT_ROTATE_90DEGREES_CLOCKWISE_MIRROR_VERICAL         = PARAMETER("+0007")

	PARAMETER_OUTPUT_POWERSAVE_1MIN  = PARAMETER("00000") // 00000 or +0004
	PARAMETER_OUTPUT_POWERSAVE_5MIN  = PARAMETER("+0001") // +0001 or +0005
	PARAMETER_OUTPUT_POWERSAVE_10MIN = PARAMETER("+0002") // +0002 or +0006
	PARAMETER_OUTPUT_POWERSAVE_NEVER = PARAMETER("+0003")

	PARAMETER_OUTPUT_TESTPATTERN_DISABLE = PARAMETER("00000")
	PARAMETER_OUTPUT_TESTPATTERN_ENABLE  = PARAMETER("+0001")
)

// ForceMute Forcemute specified
func (h *H2SC) ForceMute(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_FORCEMUTE
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// FreezeFrame Freezes frame
func (h *H2SC) FreezeFrame(freeze bool) error {
	param := PARAMETER_OUTPUT_FREEZE_DISABLE
	if freeze {
		param = PARAMETER_OUTPUT_FREEZE_ENABLE
	}

	packet := newPacket()
	packet.command = COMMAND_FREEZE
	packet.idnum = h.id
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// ForceAspectRatio Override aspect-ratio
func (h *H2SC) ForceAspectRatio(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_ASPECT
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// OverrideOutputFormat Overrides output SDI signal format.
func (h *H2SC) OverrideOutputFormat(format PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_SDIFORMAT
	packet.idnum = h.id
	packet.param = format

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// SeamlessBehavior Change seamless switching behavior.
func (h *H2SC) SeamlessBehavior(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_SEAMLESS_AND_LOCK
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// RotateAndMirror Rotating and mirroring.
func (h *H2SC) RotateAndMirror(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_ROTATE_AND_MIRROR
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// PowerSave Set power-save behavior.
func (h *H2SC) PowerSave(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_POWERSAVE
	packet.idnum = h.id
	packet.param = state

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// TestPattern Enable test pattern output.
func (h *H2SC) TestPattern(enabled bool) error {
	param := PARAMETER_OUTPUT_TESTPATTERN_DISABLE
	if enabled {
		param = PARAMETER_OUTPUT_TESTPATTERN_ENABLE
	}

	packet := newPacket()
	packet.command = COMMAND_TESTPATTERN
	packet.idnum = h.id
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err := h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// TODO...
// func (h *H2SC) OnScreenInformation(state PARAMETER) error {  }
// func (h *H2SC) GenLockHorizonalOffset(state PARAMETER) error {  }
// func (h *H2SC) GenLockVerticalOffset(state PARAMETER) error {  }
