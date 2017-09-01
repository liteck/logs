/*
** ===============================================
** USER NAME: garlic(QQ:3173413)
** FILE NAME: log.go
** DATE TIME: 2017-07-17 13:48:07
** ===============================================
 */
package logs

import (
	"fmt"
	"io"
	"log"
	"os"
)

type severity int32 // sync/atomic int32

//日志等级分类

const (
	LevelAlert = iota
	LevelCritical
	LevelError
	LevelWarning
	LevelNotice
	LevelInformational
	LevelDebug
	numSeverity = 7
)

//默认颜色配置
var (
	green      = "\033[22;32m"
	brown      = "\033[22;33m"
	yellow     = "\033[01;33m"
	red        = "\033[22;31m"
	blue       = "\033[22;34m"
	magenta    = "\033[22;35m"
	cyan       = "\033[22;36m"
	light_blue = "\033[01;34m"
	light_red  = "\033[01;31m"

	reset = "\033[0m"
)

//定义默认控制栏输出.以及前缀为空.格式为时间日期+文件名
var std = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)

func Info(i ...interface{}) {
	formatPrint(LevelInformational, fmt.Sprint(i...))
}

func Debug(i ...interface{}) {
	formatPrint(LevelDebug, fmt.Sprint(i...))
}

func Notice(i ...interface{}) {
	formatPrint(LevelNotice, fmt.Sprint(i...))
}

func Warning(i ...interface{}) {
	formatPrint(LevelWarning, fmt.Sprint(i...))
}

func Error(i ...interface{}) {
	formatPrint(LevelError, fmt.Sprint(i...))
}

func Cirtical(i ...interface{}) {
	formatPrint(LevelCritical, fmt.Sprint(i...))
}

func Alert(i ...interface{}) {
	formatPrint(LevelAlert, fmt.Sprint(i...))
}

//格式化输出
func formatPrint(level int, i string) {
	l := ""
	color := ""
	switch level {
	case LevelAlert:
		color = yellow
		l = "[A]"
		break
	case LevelCritical:
		color = magenta
		l = "[C]"
		break
	case LevelError:
		color = light_red
		l = "[E]"
		break
	case LevelWarning:
		color = brown
		l = "[W]"
		break
	case LevelNotice:
		color = cyan
		l = "[N]"
		break
	case LevelInformational:
		color = light_blue
		l = "[I]"
		break
	case LevelDebug:
		color = green
		l = "[D]"
		break
	}

	msg := fmt.Sprintf(" %s%s%s %s%s%s", color, l, reset, color, i, reset)
	std.Println(msg)
}

// 定义日志输出位置
func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

// 返回当前的打印标志位
func Flags() int {
	return std.Flags()
}

// 设置标志
func SetFlags(flag int) {
	std.SetFlags(flag)
}

// 返回前缀
func Prefix() string {
	return std.Prefix()
}

// 设置前缀
func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}
