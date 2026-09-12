package goos

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"sync"

	nodeprocess "github.com/i-am-the-slime/go-ffi/purescript-node-process"
	. "github.com/purescript-native/go-runtime"
)

type signalSubscription struct {
	channel chan os.Signal
	done    chan struct{}
	once    sync.Once
}

func init() {
	exports := Foreign("Goos.Process")
	exports["platform"] = runtime.GOOS
	exports["environment"] = func(key Any) Any { return func() Any { return os.Getenv(key.(string)) } }
	exports["pid"] = func() Any { return os.Getpid() }
	exports["uid"] = func() Any { return os.Getuid() }
	exports["tempDirectory"] = func() Any { return os.TempDir() }
	exports["absolutePath"] = func(path Any) Any {
		return func() Any {
			absolute, err := filepath.Abs(path.(string))
			if err != nil {
				panic(err)
			}
			return absolute
		}
	}
	exports["dirname"] = func(path Any) Any { return filepath.Dir(path.(string)) }
	exports["basename"] = func(path Any) Any { return filepath.Base(path.(string)) }
	exports["joinPath"] = func(a Any) Any { return func(b Any) Any { return filepath.Join(a.(string), b.(string)) } }
	exports["removeFile"] = func(path Any) Any {
		return func() Any {
			if err := os.Remove(path.(string)); err != nil {
				panic(err)
			}
			return nil
		}
	}
	exports["finish"] = func() Any { os.Exit(nodeprocess.ExitCode()); return nil }
	exports["sendSignal"] = func(pid Any) Any {
		return func(name Any) Any {
			return func() Any {
				sig, ok := namedSignal(name.(string))
				if !ok {
					return false
				}
				process, err := os.FindProcess(pid.(int))
				if err != nil {
					return false
				}
				defer process.Release()
				return process.Signal(sig) == nil
			}
		}
	}
	exports["subscribeSignals"] = func(names Any) Any {
		return func() Any {
			values := names.([]Any)
			signals := make([]os.Signal, len(values))
			for i, value := range values {
				sig, ok := namedSignal(value.(string))
				if !ok {
					panic(fmt.Errorf("unsupported signal: %s", value))
				}
				signals[i] = sig
			}
			sub := &signalSubscription{channel: make(chan os.Signal, len(signals)), done: make(chan struct{})}
			if len(signals) > 0 {
				signal.Notify(sub.channel, signals...)
			}
			return sub
		}
	}
	exports["receiveSignal"] = func(value Any) Any {
		return func() Any {
			sub := value.(*signalSubscription)
			select {
			case sig := <-sub.channel:
				return sig.String()
			case <-sub.done:
				return ""
			}
		}
	}
	exports["closeSignals"] = func(value Any) Any {
		return func() Any {
			sub := value.(*signalSubscription)
			sub.once.Do(func() { signal.Stop(sub.channel); close(sub.done) })
			return nil
		}
	}
}
