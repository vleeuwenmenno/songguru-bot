package logging

import (
	"fmt"
	"time"
)

func PrintLog(msg string, args ...interface{}) {
	fmt.Println(fmt.Sprintf(time.Now().Format("2006-01-02 15:04:05")+" : "+msg, args...))
}
