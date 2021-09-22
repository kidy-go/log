// file.go kee > 2021/09/21

package log

import (
	"errors"
	"fmt"
	"os"
)

type fileWriter struct {
	file *os.File
}

func (f *fileWriter) NewFile(file string) error {
	var err error
	if nil != f.file {
		f.file.Close()
	}
	f.file, err = os.Create(file)
	return err
}

func (f *fileWriter) Write(message interface{}) error {
	if nil != f.file {
		return errors.New("file not created")
	}

	msg := fmt.Sprintf("%v\n", message)

	_, err := f.file.Write([]byte(msg))

	return err
}

// Init file logger with map[string]interface{}
// config like:
// {
//		"path" : "logs/error.log",
//		"level": "debug",
//		"days" : 30
//		"daily": true
// }
func (f *fileWriter) Init(conf map[string]string) error {
	// path := conf["path"]
	return nil
}

func NewFileWriter() *fileWriter {
	return &fileWriter{}
}
