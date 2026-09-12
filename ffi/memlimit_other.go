//go:build !darwin && !linux

package goos

func physicalMemoryBytes() int64 {
	return 0
}
