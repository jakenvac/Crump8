package crumplib

import "io"

// LogWriter is the writer to log messages to
var LogWriter io.Writer

// LogWrite writes a string to LogWriter
func LogWrite(message string) {
	message += "\n"
	if LogWriter != nil {
		LogWriter.Write([]byte(message))
	}
}
