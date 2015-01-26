package debugger

type loglevel int
// DEBUG LEVEL
const (
	DEBUG  loglevel = 1 << iota
	INFO
	NOTICE
	WARN
	ERROR
)

// String: 对于自定义的struct类型才会用到指针，对于loglevel这种东西是不要写成 *loglevel的
func (this loglevel) String() string {
	switch this{
		case DEBUG: return "DEBUG"
		case INFO: return "INFO"
		case NOTICE: return "NOTICE"
		case WARN: return "WARN"
		case ERROR: return "ERROR"
	}
	return "UNDEFINED"
}
