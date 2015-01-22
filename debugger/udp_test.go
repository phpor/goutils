package debugger

import "testing"

func TestInfo(t *testing.T) {
	d := NewUdpDebugger("10.79.80.245:2345")
	d.Enable()
	d.SetLevel(INFO)
	d.Info(&Msg{"hello"})

}
