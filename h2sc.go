package h2sc

import (
	"fmt"
	"net"
	"time"
)

const (
	// IPaddr Default H2SC IP address(192.168.002.222)
	IPaddr = "192.168.002.222"
	// Port Default H2SC Port(1300)
	Port              = 1300
	fixedPacketLength = 12
)

var (
	header = []byte("#$")
	footer = byte('\r')

	timeout = time.Millisecond * 50 // request timeout for command TCP API. default:50ms
)

// checkPacketLength Checks packet lenth is valid(12).
func checkPacketLength(packet []byte) error {
	if len(packet) != fixedPacketLength {
		return fmt.Errorf("Invalid packet length. [valid]=[%d] , [request]=%d", fixedPacketLength, len(packet))
	}
	return nil
}

// newPacket Generate new packet
func newPacket() *controlPacket {
	return &controlPacket{
		header: header,
		footer: footer,
	}
}

// controlPacket Control command packets
type controlPacket struct {
	header  []byte    // 2bytes fixed header
	idnum   string    // 2bytes ID Number. e.g. 00~99
	command COMMAND   // 2bytes command.
	param   PARAMETER // 5bytes parameter. e.g. 00000, -9999, +9999.
	footer  byte      // 1byte carriage return.(0x0d)
}

func (p *controlPacket) Validate() error {
	if len(p.header) != 2 {
		return fmt.Errorf("Invalid header size")
	}
	if len(p.idnum) != 2 {
		return fmt.Errorf("Invalid ID size")
	}
	if len(p.command) != 2 {
		return fmt.Errorf("Invalid command size")
	}
	if len(p.param) != 5 {
		return fmt.Errorf("Invalid ID size")
	}
	if p.footer != footer {
		return fmt.Errorf("Unknown footer byte")
	}
	return nil
}

// ToFixedArray Generates Fixed length(12) array.
func (p *controlPacket) ToFixedArray() ([12]byte, error) {
	if err := p.Validate(); err != nil {
		return [12]byte{}, err
	}
	b := make([]byte, 0, 12)
	b = append(b, p.header...)
	b = append(b, []byte(p.idnum)...)
	b = append(b, p.command...)
	b = append(b, p.param...)
	b = append(b, footer)
	var packet [12]byte
	copy(packet[:], b)
	return packet, nil
}

// ToSlice Generates variable length slice.
func (p *controlPacket) ToSlice() ([]byte, error) {
	if err := p.Validate(); err != nil {
		return nil, err
	}
	b := make([]byte, 0, 12)
	b = append(b, p.header...)
	b = append(b, []byte(p.idnum)...)
	b = append(b, p.command...)
	b = append(b, p.param...)
	b = append(b, footer)
	return b, nil
}

// H2SC H2SC main struct
type H2SC struct {
	ipAddr net.IP // IP address
	id     string // id number?
	conn   net.Conn
	closed bool
}

// New H2SC Struct Pointer
func New() (*H2SC, error) {
	addr := fmt.Sprintf("%s:%d", IPaddr, Port)
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	return &H2SC{
		ipAddr: net.ParseIP(IPaddr).To4(),
		conn:   conn,
		closed: false,
	}, nil
}

// Disconnect Close connection.
func (h *H2SC) Disconnect() error {
	if err := h.conn.Close(); err != nil {
		return err
	}
	h.closed = true
	return nil
}

// Closed Check if connection is closed or not.
func (h *H2SC) Closed() bool {
	return h.closed
}
