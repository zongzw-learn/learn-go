package main

import (
	"os"

	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

/*
var a = test()
func test() string {

	fmt.Printf("hello packages \n")
	return "a"
}
*/

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.6s} %{id:013x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {
	// For demo purposes, create two backend for os.Stderr.
	bkdOut := logging.NewLogBackend(os.Stdout, "", 0)
	bkdErr := logging.NewLogBackend(os.Stderr, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.

	bkdOutFmt := logging.NewBackendFormatter(bkdOut, format)
	bkdErrFmt := logging.NewBackendFormatter(bkdErr, format)

	bkdOutFmtLvl := logging.AddModuleLevel(bkdOutFmt)
	bkdOutFmtLvl.SetLevel(logging.DEBUG, "")
	bkdErrFmtLvl := logging.AddModuleLevel(bkdErrFmt)
	bkdErrFmtLvl.SetLevel(logging.DEBUG, "")

	// Set the backends to be used.
	logging.SetBackend(bkdOutFmtLvl, bkdErrFmtLvl)

	log.Debugf("debug %s", Password("secret"))
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("err")
	log.Critical("crit")
}
