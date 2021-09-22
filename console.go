// console.go kee > 2021/09/21

package log

import (
	"os"
)

type consoleWriter struct {
	level Level
}

func (f *consoleWriter) Write(msg []byte) error {
	_, err := os.Stdout.Write(msg)

	return err
}

func (c *consoleWriter) Level() Level {
	return c.level
}

func NewConsoleWriter(level Level) *consoleWriter {
	return &consoleWriter{level: level}
}
