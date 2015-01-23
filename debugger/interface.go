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
	Debug(Messager)
	Info(Messager)
	Notice(Messager)
	Warn(Messager)
	Error(Messager)

	Enable()
	Disable()
	SetLevel(int32)
}

type Messager interface {
	String() string
}
