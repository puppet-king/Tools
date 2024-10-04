// Copyright 2024 The wangkai. ALL rights reserved.

/*
Package fileutils
*/
package fileutils

import "fmt"

// MyError 自定义的错误类型
type MyError struct {
	Message string
}

// 实现 Error() 方法，使 MyError 实现 error 接口
func (e *MyError) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
