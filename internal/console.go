package internal

import (
	"strings"
)

var consoleFlag map[string]string = map[string]string{}
var consoleCommand []string

// InitConsole 初始化控制台参数
func InitConsole(args []string) {
	jump := 0
	length := len(args)
	for i := 0; i < length; i++ {
		if jump == i {
			continue
		}
		if strings.HasPrefix(args[i], "-") {
			val := ""
			if length > i+1 {
				val = args[i+1]
				jump = i + 1
			}
			if strings.HasPrefix(args[i], "--") {
				consoleFlag[args[i][2:]] = val
			} else {
				consoleFlag[args[i][1:]] = val
			}
		} else {
			if args[i] == "start" {
				consoleFlag["worker"] = args[i+1]
				jump = i + 1
			}
			if args[i] == "open" {
				consoleFlag["server"] = args[i+1]
				jump = i + 1
			}
			if args[i] == "close" {
				consoleFlag["server"] = args[i+1]
				jump = i + 1
			}
			consoleCommand = append(consoleCommand, args[i])
		}
	}
}

// HasCommand 是否存在指令
func HasCommand(name string) bool {
	for _, v := range consoleCommand {
		if v == name {
			return true
		}
	}
	return false
}

// GetCommands 是否存在指令
func GetCommands() []string {
	return consoleCommand
}

// HasFlag 是否存在参数
func HasFlag(name string) bool {
	_, ok := consoleFlag[name]
	return ok
}

// GetFlag 参数值
func GetFlag(name string) string {
	return consoleFlag[name]
}

// GetFlagAuto 参数值，如果没有取默认值
func GetFlagAuto(name string, defaultValue string) string {
	val, ok := consoleFlag[name]
	if ok {
		return val
	}
	return defaultValue
}
