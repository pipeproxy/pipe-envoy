package logger

import (
	"log"
	"os"
)

var std = newWrapLogger(nopCloser{os.Stderr}, 3)

func Debug() {
	std.Logger.SetFlags(std.Logger.Flags() | log.Llongfile)
	log.SetFlags(log.Flags() | log.Llongfile)
}

func init() {
	Debug()
}

// Info calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Info(v ...interface{}) {
	std.Info(v...)
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Infof(format string, v ...interface{}) {
	std.Infof(format, v...)
}

// Infoln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Infoln(v ...interface{}) {
	std.Infoln(v...)
}

// Error calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Error(v ...interface{}) {
	std.Error(v...)
}

// Errorf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, v ...interface{}) {
	std.Errorf(format, v...)
}

// Errorln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Errorln(v ...interface{}) {
	std.Errorln(v...)
}

// Warn calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Warn(v ...interface{}) {
	std.Warn(v...)
}

// Warnf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	std.Warnf(format, v...)
}

// Warnln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Warnln(v ...interface{}) {
	std.Warnln(v...)
}

// Fatal calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Fatal(v ...interface{}) {
	std.Fatal(v...)
}

// Fatalf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Fatalf(format string, v ...interface{}) {
	std.Fatalf(format, v...)
}

// Fatalln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Fatalln(v ...interface{}) {
	std.Fatalln(v...)
}

// Todo calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Todo(v ...interface{}) {
	std.Todo(v...)
}

// Todof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Todof(format string, v ...interface{}) {
	std.Todof(format, v...)
}

// Todoln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Todoln(v ...interface{}) {
	std.Todoln(v...)
}
