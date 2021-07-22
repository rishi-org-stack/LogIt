package log

import (
	"time"

	"github.com/golang/glog"
)

type (
	message struct {
		source string
		level  string
		mess   interface{}
	}
)

const (
	infoLog  = "INFO: "
	warnLog  = "WARN: "
	errorLog = "Error: "
)

func Init(source, level string) *message {
	return &message{
		source: source,
		level:  level,
	}
}
func (m *message) Info() {
	glog.Infoln(
		infoLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
func (m *message) Error() {
	glog.Errorln(
		errorLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
func (m *message) Warn() {
	glog.Warningln(
		warnLog+"\n\tDATE: ",
		time.Now().Day(),
		time.Now().Month(),
		time.Now().Year(),
		" time : ",
		time.Now().Hour(),
		time.Now().Minute(),
		"\n\tMessage :\n", "\t\t",
		m.source,
		"\n\t\t",
		m.level,
		"\n\t\t",
		m.mess)
}
