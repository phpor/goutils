package debugger



type Debugger interface {
	Debug(Messager)
	Info(Messager)
	Notice(Messager)
	Warn(Messager)
	Error(Messager)

	Enable()
	Disable()
	SetLevel(loglevel)
}

type Messager interface {
	String() string
}
