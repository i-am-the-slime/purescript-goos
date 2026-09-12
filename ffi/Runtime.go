package goos

import (
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"time"

	. "github.com/purescript-native/go-runtime"
)

// Version may be supplied with go build -ldflags -X.
var Version = "dev"
var monotonicOrigin = time.Now()

func init() {
	exports := Foreign("Goos.Runtime")
	exports["version"] = Version
	exports["cpuCount"] = func() Any { return runtime.NumCPU() }
	exports["monotonicMilliseconds"] = func() Any { return float64(time.Since(monotonicOrigin)) / float64(time.Millisecond) }
	exports["sleepMs"] = func(value Any) Any {
		return func() Any { time.Sleep(time.Duration(value.(int)) * time.Millisecond); return nil }
	}
	exports["physicalMemoryBytes"] = func() Any { return float64(physicalMemoryBytes()) }
	exports["setGCPercent"] = func(value Any) Any { return func() Any { debug.SetGCPercent(value.(int)); return nil } }
	exports["setMemoryLimit"] = func(value Any) Any { return func() Any { debug.SetMemoryLimit(int64(value.(float64))); return nil } }
	exports["collectGarbage"] = func() Any { runtime.GC(); return nil }
	exports["writeHeapProfile"] = func(path Any) Any {
		return func() Any {
			file, err := os.Create(path.(string))
			if err != nil {
				panic(err)
			}
			defer file.Close()
			if err := pprof.WriteHeapProfile(file); err != nil {
				panic(err)
			}
			return nil
		}
	}
}
