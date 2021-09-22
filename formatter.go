// formatter.go kee > 2021/09/21

package log

import (
	"fmt"
	"strings"
	"time"
)

type LoggerFormatter interface {
	Format(map[string]interface{}) string
}

const (
	dateFormat = "2006/01/02T15:04:05"
)

func LogFormatter(msg interface{}) []byte {
	date := (time.Now()).Format(dateFormat)
	str := fmt.Sprintf("[%s] %v\n", date, msg)
	return []byte(str)
}

type LineFormatter struct {
	formatter  string
	timeFormat string
}

func (lf *LineFormatter) Format(data map[string]interface{}) string {
	if "" == lf.formatter {
		lf.formatter = "[$datetime] ($file, $line) [$level]: $message"
	}
	if "" == lf.timeFormat {
		lf.timeFormat = dateFormat
	}

	data["datetime"] = data["datetime"].(time.Time).Format(lf.timeFormat)

	msg := lf.formatter
	for k, val := range data {
		msg = strings.Replace(msg, "$"+k, fmt.Sprintf("%v", val), -1)
	}

	return fmt.Sprintf("%v\n", msg)
}
