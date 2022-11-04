package log

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func Logger() *logrus.Entry {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		panic("could not get context info for logger!")
	}
	filename := file + ":" + strconv.Itoa(line)
	funcname := runtime.FuncForPC(pc).Name()
	fn := funcname[strings.LastIndex(funcname, ".")+1:]
	return logrus.WithField("file", filename).WithField("function", fn)
}
