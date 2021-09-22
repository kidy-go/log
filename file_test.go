// file_test.go kee > 2021/09/21

package log

import (
	"fmt"
	"os"
	"testing"
)

func TestFile(t *testing.T) {

	fileConfig := map[string]interface{}{
		"path":  "./test.log",
		"level": "DEBUG",
		"daily": false,
		"days":  30,
	}

	fmt.Println("path >", configGet(fileConfig, "path", "./file_test.log"))
	fmt.Println("level >", configGet(fileConfig, "level", "debug"))
	fmt.Println("daily >", configGet(fileConfig, "daily", false))

	f, err := os.Create(configGet(fileConfig, "path", "./file_test.log").(string))
	defer f.Close()
	if err != nil {
		fmt.Printf("File Error %v\n", err)
	}

	f.Write([]byte("helo world"))
	if true == configGet(fileConfig, "daily", 0).(bool) {
		fmt.Println("Daily is configed")
	}

	// fl := NewFileWriter()

	// fl.NewFile("./test.log")

	// if e := fl.Write("Hello File Logs"); e != nil {
	//  	panic(e)
	// }
}
