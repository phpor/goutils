package debugger

import (
	"runtime"
	"time"
	"fmt"
)

type SimpleMessage struct {
	ctime time.Time
	file  string
	line  int
	fun   string
	tag   string
	text  string
}

func NewSimpleMessage(tag, text string) *SimpleMessage {
	pc, file, line, ok := runtime.Caller(1)
	fun := ""
	if ok {
		fun = runtime.FuncForPC(pc).Name()
	}
	return &SimpleMessage{ctime: time.Now(), file: file, line: line, fun: fun, tag: tag, text: text}

}

// 消息格式化
func (this *SimpleMessage) String() string {
	return fmt.Sprintf("%s %s %s:%d:%s: %s\n", this.ctime.Format("2006-01-02 15:04:05"), this.tag, this.file, this.line, this.fun, this.text)
}
