package h2sc

import (
	"fmt"
	"net"
	"time"
)

const (
	IPaddr            = "192.168.002.222"
	Port              = 1300
	fixedPacketLength = 12

	timeout = time.Millisecond * 50
)

var (
	header = []byte("#$")
	footer = byte('\r')
)

func checkPacketLength(packet []byte) error {
	if len(packet) != fixedPacketLength {
		return fmt.Errorf("Invalid packet length. [valid]=[%d] , [request]=%d", fixedPacketLength, len(packet))
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
