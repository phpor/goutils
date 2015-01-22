package debugger

// 设置调试级别
type base struct {
	enable bool
	level  uint32
}

func (this *base) write(level uint32, msg *Msg, debugger Debugger, fn func(Debugger, *Msg)) {
	if !this.enable || this.level&level == 0 {
		return
	}
	fn(debugger, msg)
}

func (this *base) Enable() {
	this.enable = true
}

func (this *base) Disable() {
	this.enable = false
}

func (this *base) SetLevel(level uint32) {
	this.level |= level
}
