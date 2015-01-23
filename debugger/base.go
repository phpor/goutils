package debugger

// 设置调试级别
type debuggerBase struct {
	enable bool
	level  int32
}

func (this *debuggerBase) write(level int32, fn func() bool) bool{
	if !this.enable || (this.level&level == 0) {
		return false
	}
	return fn()
}

func (this *debuggerBase) Enable() {
	this.enable = true
}

func (this *debuggerBase) Disable() {
	this.enable = false
}

func (this *debuggerBase) SetLevel(level int32) {
	this.level = level
}
