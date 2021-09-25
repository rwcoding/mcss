package internal

import (
	"os"
	"testing"
)

func TestConsole(t *testing.T) {
	args := []string{os.Args[0], "do", "-name", "lucy"}
	InitConsole(args)

	if !HasCommand("do") {
		t.Error("command ?")
	}

	if HasCommand("do1") {
		t.Error("why do1 ?")
	}

	if !HasFlag("name") {
		t.Error("flag ?")
	}

	if HasFlag("name1") {
		t.Error("why name1 ??")
	}

	if GetFlag("name") != "lucy" {
		t.Error("name != lucy ??")
	}

	if GetFlagAuto("age", "20") != "20" {
		t.Error("age != 20 ??")
	}

	//复原
	InitConsole(os.Args)
}
