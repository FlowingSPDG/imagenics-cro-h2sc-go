package h2sc

import (
	"fmt"
	"net"
)

const (
	IPaddr            = "192.168.002.222"
	Port              = 1300
	fixedPacketLength = 12
)

var (
	header = []byte("#$")
	footer = byte('\r')
)

func checkPacketLength(packet []byte) error {
	if len(packet) != fixedPacketLength {
		return fmt.Errorf("Invalid packet length. [valid]=[12] , [request]=%d", len(packet))
	}
	return nil
}

func newPacket() *controlPacket {
	return &controlPacket{
		header: header,
		footer: footer,
	}
}

type controlPacket struct {
	header  []byte // 2bytes fixed header
	idnum   string // 2bytes ID Number. e.g. 00~99
	command []byte // 2bytes command.
	param   []byte // 5bytes parameter. e.g. 00000, -9999, +9999.
	footer  byte   // 1byte carriage return.(0x0d)
}

func (p *controlPacket) ToFixedSlice() [12]byte {
	b := make([]byte, 12)
	b = append(b, p.header...)
	b = append(b, []byte(p.idnum)...)
	b = append(b, p.command...)
	b = append(b, p.param...)
	b = append(b, footer)
	var packet [12]byte
	copy(packet[:], b)
	return packet
}

func (p *controlPacket) ToSlice() []byte {
	b := make([]byte, 0, 12)
	b = append(b, p.header...)
	b = append(b, []byte(p.idnum)...)
	b = append(b, p.command...)
	b = append(b, p.param...)
	b = append(b, footer)
	return b
}

// H2SC H2SC main struct
type H2SC struct {
	ipAddr net.IP // IP address
	id     string // id number?
}

// New H2SC Struct Pointer
func New() *H2SC {
	return &H2SC{
		ipAddr: net.ParseIP(IPaddr).To4(),
	}
}

// Identify Give H2SC specific ID.
func (h *H2SC) Identify(id string) error {
	addr := fmt.Sprintf("%s:%d", h.ipAddr.String(), Port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()

	// TODO
	packet := newPacket()
	packet.command = COMMAND_IDENTIFY
	packet.idnum = "00"
	packet.param = []byte("+" + id) // e.g. +0001

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	conn.Write(ps)

	h.id = id
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

	param := []byte("00000")
	if enabled {
		param = []byte("+0001")
	}

	packet := newPacket()
	packet.command = COMMAND_TESTPATTERN
	packet.idnum = "01"
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	conn.Write(ps)

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

	param := []byte("00000")
	if freeze {
		param = []byte("+0001")
	}

	packet := newPacket()
	packet.command = COMMAND_FREEZE
	packet.idnum = "01"
	packet.param = param

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	conn.Write(ps)

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
	packet.idnum = "01"
	packet.param = format

	ps := packet.ToSlice()
	if err := checkPacketLength(ps); err != nil {
		return err
	}
	conn.Write(ps)

	return nil
}
