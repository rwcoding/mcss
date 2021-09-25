package internal

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func InitCmd() {
	has := false
	for _, v := range GetCommands() {
		if v == "build" {
			has = true
			if err := CmdBuild(); err != nil {
				log.Fatal("error:", err)
			}
		}
		script, ok := Options.Script[v]
		if !ok || script == "" {
			continue
		}
		has = true
		if err := RunCmd(script.(string)); err != nil {
			log.Fatal("error:", err)
		}
	}
	if has {
		os.Exit(0)
	}
}

func RunCmd(script string) error {
	reg, _ := regexp.Compile("\\s{2,}")
	command := strings.Split(reg.ReplaceAllString(strings.TrimSpace(script), " "), " ")
	var cmd *exec.Cmd
	if len(command) > 1 {
		cmd = exec.Command(command[0], command[1:]...)
	} else {
		cmd = exec.Command(command[0])
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
