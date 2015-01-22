package debugger

import "testing"

func TestInfo(t *testing.T) {
	d := NewUdpDebugger("10.79.80.245:2345")
	d.enable()
	d.Info(&Msg{"hello"})

}
