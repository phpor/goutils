// debugger 的UDP实现
package debugger

import "net"

type UdpDebugger struct {
	conn *net.UDPConn
	debuggerBase
}

func NewUdpDebugger(addr string) *UdpDebugger {
	udpAddress, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		panic(err)
	}
	conn, err := net.DialUDP("udp", nil, udpAddress)
	return &UdpDebugger{conn: conn, debuggerBase: debuggerBase{enable: false, level: uint32(0)}}
}


func (this *UdpDebugger) Debug(msg Messager) {
	this.write(DEBUG, msg)
}
func (this *UdpDebugger) Info(msg Messager) {
	this.write(INFO, msg)
}
func (this *UdpDebugger) Notice(msg Messager) {
	this.write(NOTICE, msg)
}
func (this *UdpDebugger) Warn(msg Messager) {
	this.write(WARN, msg)
}
func (this *UdpDebugger) Error(msg Messager) {
	this.write(ERROR, msg)
}

func (this *UdpDebugger) write(level uint32, msg Messager) {
	this.debuggerBase.write(level, func() bool{
			_, err := this.conn.Write([]byte(msg.String()))
			if err != nil {
				return false
			}
			return true
		})
}
