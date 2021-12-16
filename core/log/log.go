package log

import "fmt"

//Error 自己封装一层
func Error(args ...interface{}) {
	fmt.Println(args...)
	panic("error")
}
