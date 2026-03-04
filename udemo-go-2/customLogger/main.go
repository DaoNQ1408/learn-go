package main

import "fmt"

// khai báo kiểu LogLevel, có nền tảng là int
type LogLevel int

// khai báo enum có kiểu LogLevel
const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR"}

// kdl đc dùng hàm này
// tên hàm
// return type của hàm
func (l LogLevel) String() string  {

	if l < LevelTrace || l > LevelError {
		return "UNKNOWN"
	}

	return levelNames[l]

}


func printLogLevel(level LogLevel) {
	fmt.Printf("Log level: %d %s\n", level, level.String())
}


func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)
	printLogLevel(10) // log level không hợp lệ
	printLogLevel(1) // tương đương với index 1 trong LogLevel
}
