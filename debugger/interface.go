package debugger

// DEBUG LEVEL
const (
	DEBUG  = 1
	INFO   = 1 << 1
	NOTICE = 1 << 2
	WARN   = 1 << 3
	ERROR  = 1 << 4
)

type Debugger interface {
	Debug(*Msg)
	Info(*Msg)
	Notice(*Msg)
	Warn(*Msg)
	Error(*Msg)

	Enable()
	Disable()
	SetLevel(uint32)
}
