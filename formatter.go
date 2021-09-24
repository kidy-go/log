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
	format     string
	timeFormat string
}

func (lf *LineFormatter) getFormat() string {
	if "" == lf.format {
		lf.format = "[$datetime] ($file, $line) [$level]: $message $context"
	}
	return lf.format
}

func (lf *LineFormatter) getTimeFormat() string {
	if "" == lf.timeFormat {
		lf.timeFormat = dateFormat
	}
	return lf.timeFormat
}

func (lf *LineFormatter) Format(data map[string]interface{}) string {

	if ctx, _ := data["context"]; len(ctx.([]interface{})) == 0 {
		data["context"] = ""
	}
	data["datetime"] = data["datetime"].(time.Time).Format(lf.getTimeFormat())

	msg := lf.getFormat()
	for k, val := range data {
		msg = strings.Replace(msg, "$"+k, fmt.Sprintf("%v", val), -1)
	}

	return fmt.Sprintf("%v\n", msg)
}
