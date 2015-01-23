package debugger

// 设置调试级别
type debuggerBase struct {
	enable bool
	level  uint32
}

func (this *debuggerBase) write(level uint32, fn func()) {
	if !this.enable || this.level&level == 0 {
		return
	}
	fn()
}

func (this *debuggerBase) Enable() {
	this.enable = true
}

func (this *debuggerBase) Disable() {
	this.enable = false
}

func (this *debuggerBase) SetLevel(level uint32) {
	this.level |= level
}
