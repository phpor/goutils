package debugger

// 或许debugger本身是个多余的东西，这本身应该是logger可以做的事情，即： logger.Debug(...)
// 问题
// 1. udp.go 中是否每次发送消息都使用一个新的连接？如果不是，何时关闭连接（尽管是UDP）？
// 2. 如果让debugger提供一个Close()的方法，显得有些ugly；
// 3. 或者说，只要对象存在就永远使用相同的连接
