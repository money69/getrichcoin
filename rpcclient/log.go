// Copyright (c) 2014-2017 The grhsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package rpcclient

import (
	"github.com/grhsuite/grhlog"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log grhlog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until UseLogger is called.
func DisableLog() {
	log = grhlog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger grhlog.Logger) {
	log = logger
}

// LogClosure is a closure that can be printed with %v to be used to
// generate expensive-to-create data for a detailed log level and avoid doing
// the work if the data isn't printed.
type logClosure func() string

// String invokes the log closure and returns the results string.
func (c logClosure) String() string {
	return c()
}

// newLogClosure returns a new closure over the passed function which allows
// it to be used as a parameter in a logging function that is only invoked when
// the logging level is such that the message will actually be logged.
func newLogClosure(c func() string) logClosure {
	return logClosure(c)
}
