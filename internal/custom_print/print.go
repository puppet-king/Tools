package custom_print

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func Success(format string, a ...any) {
	prefix := "[SUCCESS]" // 可以根据需要自定义前缀
	fullFormat := color.Ize(color.Green, fmt.Sprintf("%s   %s \n", prefix, format))
	fmt.Printf(fullFormat, a...)
}

func Info(format string, a ...any) {
	prefix := "[INFO]" // 可以根据需要自定义前缀
	fullFormat := color.Ize(color.Gray, fmt.Sprintf("%s  %s \n", prefix, format))
	fmt.Printf(fullFormat, a...)
}

func Error(format string, a ...any) {
	prefix := "[ERROR]" // 可以根据需要自定义前缀
	fullFormat := color.Ize(color.Red, fmt.Sprintf("%s  %s \n", prefix, format))
	fmt.Printf(fullFormat, a...)
}

func Trace(format string, a ...any) {
	prefix := "[TRACE]" // 可以根据需要自定义前缀
	fullFormat := color.Ize(color.Cyan, fmt.Sprintf("%s  %s \n", prefix, format))
	fmt.Printf(fullFormat, a...)
}
