// logger.go kee > 2021/09/20

package log

import (
	//"fmt"
	"runtime"
	"sync"
	"time"
)

type LoggerWriter interface {
	Level() Level
	Write([]byte) error
}

type loggerWriter struct {
	sync.Mutex
	writer LoggerWriter
	level  Level
}

func NewWriter(writer LoggerWriter) *loggerWriter {
	return &loggerWriter{
		writer: writer,
	}
}

func (l *loggerWriter) write(msg []byte) error {
	l.Lock()
	defer l.Unlock()
	return l.writer.Write(msg)
}

type Logger struct {
	name      string
	writers   []LoggerWriter
	formatter LoggerFormatter
}

// 注册日志通道
func NewLogger(name string, writers ...LoggerWriter) *Logger {
	lg := &Logger{}
	lg.PushWriter(writers...)

	return lg
}

// 创建日志堆栈
func (lg *Logger) PushWriter(w ...LoggerWriter) {
	lg.writers = append(lg.writers, w...)
}

func (lg *Logger) PopWriter() {
	if len(lg.writers) > 0 {
		lg.writers = lg.writers[1:]
	}
}

func (lg *Logger) GetWriters() []LoggerWriter {
	return lg.writers
}

func (lg *Logger) SetFormatter(formatter LoggerFormatter) {
	lg.formatter = formatter
}

func (lg *Logger) Write(calldepth int, level Level, message interface{}, context ...interface{}) {
	if len(lg.writers) == 0 {
		lg.PushWriter(NewConsoleWriter(LevelDebug))
	}

	if lg.formatter == nil {
		lg.formatter = &LineFormatter{}
	}

	// formatter
	_, file, ln, _ := runtime.Caller(calldepth)
	msg := lg.formatter.Format(map[string]interface{}{
		"datetime": time.Now(),
		"file":     file,
		"line":     ln,
		"level":    level,
		"channel":  lg.name,
		"message":  message,
		"context":  context,
	})

	for _, w := range lg.writers {
		if level <= w.Level() {
			go w.Write([]byte(msg))
		}
	}
}

func (lg *Logger) Debug(message interface{}, context ...interface{}) {
	lg.Write(2, LevelDebug, message, context...)
}
func (lg *Logger) Info(message interface{}, context ...interface{}) {
	lg.Write(2, LevelInfo, message, context...)
}
func (lg *Logger) Notice(message interface{}, context ...interface{}) {
	lg.Write(2, LevelNotice, message, context...)
}
func (lg *Logger) Warn(message interface{}, context ...interface{}) {
	lg.Write(2, LevelWarn, message, context...)
}
func (lg *Logger) Error(message interface{}, context ...interface{}) {
	lg.Write(2, LevelError, message, context...)
}
func (lg *Logger) Critical(message interface{}, context ...interface{}) {
	lg.Write(2, LevelCritical, message, context...)
}
func (lg *Logger) Alert(message interface{}, context ...interface{}) {
	lg.Write(2, LevelAlert, message, context...)
}
func (lg *Logger) Emerg(message interface{}, context ...interface{}) {
	lg.Write(2, LevelEmerg, message, context...)
}
