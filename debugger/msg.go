package debugger

type Msg struct {
	string string
}

// 消息格式化
func (this *Msg) String() string {

	return this.string
}
// 消息格式化
func (this *Msg) Bytes() []byte {

	return []byte(this.String())
}
