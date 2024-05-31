package requtil

type Logger interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
}

type DullLogger struct{}

func (l *DullLogger) Infof(format string, args ...interface{})    {}
func (l *DullLogger) Errorf(format string, args ...interface{})   {}
func (l *DullLogger) Warningf(format string, args ...interface{}) {}
