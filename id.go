package h2sc

import (
	"fmt"
	"net"
)

// Identify Give H2SC specific ID.
func (h *H2SC) Identify(id string) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	packet := newPacket()
	packet.command = COMMAND_IDENTIFY
	packet.idnum = "00"
	packet.param = []byte("+" + id) // e.g. +0001

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	_, err = conn.Write(ps)
	if err != nil {
		return err
	}

	h.id = id
	return nil
}
