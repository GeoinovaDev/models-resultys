package syslog

// Logger ...
type Logger struct {
}

// Save ...
func (log Logger) Save(message interface{}) {
	Insert(message)
}
