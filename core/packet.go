package core

const (
	NormalPacket int32 = iota
	OpenBracket
	CloseBracket
	Signal
)

type Packet struct {
	Contents interface{}
	PktType  int32
	owner    *Process
}
