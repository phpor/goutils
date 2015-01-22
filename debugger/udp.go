// debugger 的UDP实现
package debugger

import "net"

type UdpDebugger struct {
	conn *net.UDPConn
	base
}

func NewUdpDebugger(addr string) *UdpDebugger {
	udpAddress, err := net.ResolveUDPAddr("udp4", addr)
	conn, err := net.DialUDP("udp", nil, udpAddress)
	return &UdpDebugger{conn: conn} //, base: {false, uint32(0)}
}

func (this *UdpDebugger) Info(msg *Msg) {
	this.base.write(uint32(INFO), msg, this, this.info)
//	this.conn.Write(msg.Bytes())
}

func (this *UdpDebugger) info(msg *Msg) {
	this.conn.Write(msg.Bytes())
}
