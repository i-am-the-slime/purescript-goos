//go:build windows || plan9

package goos

import "os"

func namedSignal(name string) (os.Signal, bool) {
	switch name {
	case "INT":
		return os.Interrupt, true
	case "KILL":
		return os.Kill, true
	default:
		return nil, false
	}
}
