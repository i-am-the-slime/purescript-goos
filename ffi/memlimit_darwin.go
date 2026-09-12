package goos

import "golang.org/x/sys/unix"

func physicalMemoryBytes() int64 {
	mem, err := unix.SysctlUint64("hw.memsize")
	if err != nil {
		return 0
	}
	return int64(mem)
}
