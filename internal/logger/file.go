package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

type fileLogger struct {
	file io.WriteCloser
	path string
}

func newFileLogger(file string) (*fileLogger, error) {
	if file != "" {
		abs, err := filepath.Abs(file)
		if err != nil {
			return nil, err
		}
		file = abs
	}

	err := os.MkdirAll(filepath.Dir(file), 0755)
	if err != nil {
		return nil, err
	}
	l := &fileLogger{
		path: file,
	}
	f, ok := outfile[file]
	if ok {
		l.file = f
	} else {
		log.Printf("[INFO] open log file: %s", l.path)
		f, err := openAppendFile(l.path)
		if err != nil {
			return nil, err
		}
		l.file = f
	}
	return l, nil
}

func (l *fileLogger) Close() error {
	old := l.file

	log.Printf("[INFO] close log file: %s", l.path)
	return old.Close()
}

func (l *fileLogger) Write(p []byte) (n int, err error) {
	if l.file != nil {
		return l.file.Write(p)
	}
	return len(p), nil
}

func openAppendFile(filename string) (*os.File, error) {
	return os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
}
