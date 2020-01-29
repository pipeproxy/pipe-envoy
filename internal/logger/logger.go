package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

var outfile = map[string]io.WriteCloser{
	"/dev/stdout": nopCloser{os.Stdout},
	"/dev/stderr": nopCloser{os.Stderr},
	"/dev/null":   emptyWriter{},
	"":            emptyWriter{},
}

var loggerFiles = map[string]*fileLogger{}

func NewLogger(file string) (Logger, error) {
	l, ok := loggerFiles[file]
	if ok {
		return newWrapLogger(l, 2), nil
	}

	l, err := newFileLogger(file)
	if err != nil {
		return nil, err
	}
	loggerFiles[file] = l
	return newWrapLogger(l, 2), nil

}

type wrapLogger struct {
	io.WriteCloser
	*log.Logger
	depth int
}

func newWrapLogger(writeCloser io.WriteCloser, depth int) *wrapLogger {
	return &wrapLogger{
		WriteCloser: writeCloser,
		Logger:      log.New(writeCloser, "", log.LstdFlags),
		depth:       depth,
	}
}

// Info calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (w *wrapLogger) Info(v ...interface{}) {
	w.Output(w.depth, "[INFO]  "+fmt.Sprint(v...))
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (w *wrapLogger) Infof(format string, v ...interface{}) {
	w.Output(w.depth, "[INFO]  "+fmt.Sprintf(format, v...))
}

// Infoln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (w *wrapLogger) Infoln(v ...interface{}) {
	w.Output(w.depth, "[INFO]  "+fmt.Sprintln(v...))
}

// Error calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (w *wrapLogger) Error(v ...interface{}) {
	w.Output(w.depth, "[ERROR] "+fmt.Sprint(v...))
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (w *wrapLogger) Errorf(format string, v ...interface{}) {
	w.Output(w.depth, "[ERROR] "+fmt.Sprintf(format, v...))
}

// Errorln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (w *wrapLogger) Errorln(v ...interface{}) {
	w.Output(w.depth, "[ERROR] "+fmt.Sprintln(v...))
}

// Warn calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (w *wrapLogger) Warn(v ...interface{}) {
	w.Output(w.depth, "[WARN]  "+fmt.Sprint(v...))
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (w *wrapLogger) Warnf(format string, v ...interface{}) {
	w.Output(w.depth, "[WARN]  "+fmt.Sprintf(format, v...))
}

// Warnln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (w *wrapLogger) Warnln(v ...interface{}) {
	w.Output(w.depth, "[WARN]  "+fmt.Sprintln(v...))
}

// Fatal calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (w *wrapLogger) Fatal(v ...interface{}) {
	w.Output(w.depth, "[FATAL] "+fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (w *wrapLogger) Fatalf(format string, v ...interface{}) {
	w.Output(w.depth, "[FATAL] "+fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (w *wrapLogger) Fatalln(v ...interface{}) {
	w.Output(w.depth, "[FATAL] "+fmt.Sprintln(v...))
	os.Exit(1)
}

// Todo calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func (w *wrapLogger) Todo(v ...interface{}) {
	w.Output(w.depth, "[TODO]  "+fmt.Sprint(v...))
}

// Todof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func (w *wrapLogger) Todof(format string, v ...interface{}) {
	w.Output(w.depth, "[TODO]  "+fmt.Sprintf(format, v...))
}

// Todoln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func (w *wrapLogger) Todoln(v ...interface{}) {
	w.Output(w.depth, "[TODO]  "+fmt.Sprintln(v...))
}

type Logger interface {
	io.WriteCloser

	// Info calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Info(v ...interface{})

	// Infof calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Infof(format string, v ...interface{})

	// Infoln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Infoln(v ...interface{})

	// Error calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Error(v ...interface{})

	// Errorf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, v ...interface{})

	// Errorln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Errorln(v ...interface{})

	// Warn calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Warn(v ...interface{})

	// Warnf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Warnf(format string, v ...interface{})

	// Warnln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Warnln(v ...interface{})

	// Fatal calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Fatal(v ...interface{})

	// Fatalf calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Fatalf(format string, v ...interface{})

	// Fatalln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Fatalln(v ...interface{})

	// Todo calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Print.
	Todo(v ...interface{})

	// Todof calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Printf.
	Todof(format string, v ...interface{})

	// Todoln calls Output to print to the standard logger.
	// Arguments are handled in the manner of fmt.Println.
	Todoln(v ...interface{})
}
