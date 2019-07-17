package main

import (
	"os"
    "os/exec"
    "runtime"
)


func execute(osType string) {
	var cmd *exec.Cmd
	if osType == "win" {
		cmd = exec.Command("powershell.exe")
	} else {
		cmd = exec.Command("/bin/sh")
	}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
    if runtime.GOOS == "windows" {
        execute("win")
    } else {
        execute("linux")
    }
}