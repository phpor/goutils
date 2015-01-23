package debugger

// Dummy is a empty debugger
type Dummy struct {
}

func NewDummyDebugger() *Dummy{
	return &Dummy{}
}
func (this *Dummy) Debug(msg Messager) {}
func (this *Dummy) Info(msg Messager) {}
func (this *Dummy) Notice(msg Messager) {}
func (this *Dummy) Warn(msg Messager) {}
func (this *Dummy) Error(msg Messager) {}

func (this *Dummy) Enable() {}
func (this *Dummy) Disable() {}
func (this *Dummy) SetLevel(level int32) {}
