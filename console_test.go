// console_test.go kee > 2021/09/21

package log

import (
	"fmt"
	glog "log"
	"testing"
)

func TestConsole(t *testing.T) {
	cl := NewConsoleWriter(LevelWarn)
	lg := NewLogger("console", cl, cl)

	//if e := lg.Write("> Hello world"); e != nil {
	//panic(e)
	//}
	fmt.Println(lg.GetWriters())

	lg.Emerg("19: Has Error")

	lg.PopWriter()

	lg.Notice("23: Hello NoticeLevel")
	s := 0
	lg.Warn("25: Warnnig")

	glog.Println(fmt.Sprintf("Golang Origin log %d", s))

}
