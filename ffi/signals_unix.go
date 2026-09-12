//go:build !windows && !plan9

package goos

import (
	"os"
	"syscall"
)

func namedSignal(name string) (os.Signal, bool) {
	switch name {
	case "0":
		return syscall.Signal(0), true
	case "HUP":
		return syscall.SIGHUP, true
	case "INT":
		return syscall.SIGINT, true
	case "TERM":
		return syscall.SIGTERM, true
	case "PIPE":
		return syscall.SIGPIPE, true
	case "KILL":
		return syscall.SIGKILL, true
	case "USR1":
		return syscall.SIGUSR1, true
	case "USR2":
		return syscall.SIGUSR2, true
	default:
		return nil, false
	}
}
