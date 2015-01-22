// debugger 的UDP实现
package debugger

import "net"

type UdpDebugger struct {
	conn *net.UDPConn
	base
}

func NewUdpDebugger(addr string) *UdpDebugger {
	udpAddress, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddress)
	return &UdpDebugger{conn: conn, base: base{false, uint32(0)}}
}

func (this *UdpDebugger) Debug(msg *Msg) {
	this.write(DEBUG, msg)
}

func (this *UdpDebugger) Info(msg *Msg) {
	this.write(INFO, msg)
}
func (this *UdpDebugger) Notice(msg *Msg) {
	this.write(NOTICE, msg)
}
func (this *UdpDebugger) Warn(msg *Msg) {
	this.write(WARN, msg)
}
func (this *UdpDebugger) Error(msg *Msg) {
	this.write(ERROR, msg)
}

func (this *UdpDebugger) write(level uint32, msg *Msg) {
	this.base.write(level, func() {
			this.conn.Write(msg.Bytes())
		})
}
