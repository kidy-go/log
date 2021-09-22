// log.go kee > 2021/09/19

package log

import (
	"fmt"
)

type Level int

const (
	LevelEmerg    Level = iota // 0: Emergency: system is unusable
	LevelAlert                 // 1: Alert: action must be taken immediately
	LevelCritical              // 2: Critical: critical conditions
	LevelError                 // 3: Error: error conditions
	LevelWarn                  // 4: Warning: warning conditions
	LevelNotice                // 5: Notice: normal but significant condition
	LevelInfo                  // 6: Informational: informational messages
	LevelDebug                 // 7: Debug: debug-level messages
)

func (lv Level) String() string {
	if str, ok := (map[Level]string{
		LevelEmerg:    "emerg",
		LevelAlert:    "alert",
		LevelCritical: "critical",
		LevelError:    "error",
		LevelWarn:     "warn",
		LevelNotice:   "notice",
		LevelInfo:     "info",
		LevelDebug:    "debug",
	})[lv]; ok {
		return str
	}
	return fmt.Sprintf("unreachable(%d)", lv)
}
