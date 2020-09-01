package h2sc

import (
	"fmt"
	"time"
)

// Identify Give H2SC specific ID.
func (h *H2SC) Identify(id string) error {
	if len(id) != 2 {
		return fmt.Errorf("Invalid id length : %s", id)
	}
	packet := newPacket()
	packet.command = COMMAND_IDENTIFY
	packet.idnum = id
	packet.param = []byte("+00" + id) // e.g. +0001

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

	h.id = id
	return nil
}

// ResetID Reset H2SC ID.
func (h *H2SC) ResetID() error {
	packet := newPacket()
	packet.command = COMMAND_IDENTIFY
	packet.idnum = "00"
	packet.param = []byte("00000") // e.g. +0001

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

	h.id = "00"
	return nil
}
