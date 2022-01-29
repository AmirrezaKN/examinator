package log

type Logger interface {
	Errorf(string, ...interface{})
	Debugf(string, ...interface{})
}
