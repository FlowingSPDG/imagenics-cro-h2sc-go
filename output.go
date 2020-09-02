package h2sc

import "time"

// ForceMute Forcemute specified
func (h *H2SC) ForceMute(state PARAMETER) error {
	packet := newPacket()
	packet.command = COMMAND_FORCEMUTE
	packet.idnum = h.id
	packet.param = state

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
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

	ps, err := packet.ToSlice()
	if err != nil {
		return err
	}
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	h.conn.SetDeadline(time.Now().Add(timeout))
	_, err = h.conn.Write(ps)
	if err != nil {
		return err
	}

	return nil
}

// TODO...
// func (h *H2SC) OnScreenInformation(state PARAMETER) error {  }
// func (h *H2SC) GenLockHorizonalOffset(state PARAMETER) error {  }
// func (h *H2SC) GenLockVerticalOffset(state PARAMETER) error {  }
